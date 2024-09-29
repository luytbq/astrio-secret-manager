package secret

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

func (repo *Repo) InsertSecretGroup(group *SecretGroup) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		log.Printf("SaveSecret query error: %s", err.Error())
		return errors.New("database error")
	}

	defer func() {
		if err != nil {
			log.Printf("InsertSecretGroup rollback")
			tx.Rollback()
		}
	}()

	groupInsertQuery := `
		insert into secret_groups(user_id, description)
		values ($1, $2)	
		returning id
	`
	log.Printf("InsertSecretGroup start")
	err = tx.QueryRow(groupInsertQuery, group.UserID, group.Description).Scan(&group.ID)
	if err != nil {
		log.Printf("SaveSecret query error: %s", err.Error())
		return errors.New("database error")
	}
	log.Printf("new groupd id: %d", group.ID)

	secretsInsertQuery := `
		insert into secrets(group_id, content)
		values ($1, $2)	
		returning id
	`

	for _, secret := range group.Secrets {
		err = tx.QueryRow(secretsInsertQuery, group.ID, secret.Content).Scan(&secret.ID)
		if err != nil {
			log.Printf("SaveSecret query error: %s", err.Error())
			return errors.New("database error")
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("SaveSecret query error: %s", err.Error())
		return errors.New("database error")
	}

	return nil
}

func (repo *Repo) GetSecretsTotal(params GetSecretsParams) (uint64, error) {
	qb := psql.Select("count(1)").From("secret_groups sg").Where(sq.Eq{"user_id": params.UserID})
	if params.Keyword != "" {
		qb = qb.Where(sq.ILike{"sg.description": "%" + params.Keyword + "%"})
	}
	query, args, err := qb.ToSql()
	if err != nil {
		log.Printf("getSecretsTotal query error: %s", err.Error())
		return 0, err
	}
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, err := repo.DB.QueryContext(context, query, args...)
	if err != nil {
		log.Printf("get secrets error: %s", err.Error())
		return 0, err
	}
	defer rows.Close()
	total := uint64(0)
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return uint64(total), nil
}

func (repo *Repo) GetSecrets(params GetSecretsParams) ([]*SecretGroup, error) {
	groupQB := psql.Select("id, description").
		From("secret_groups").
		Where(sq.Eq{"user_id": params.UserID})
	if params.Keyword != "" {
		groupQB = groupQB.Where(sq.ILike{"description": "%" + params.Keyword + "%"})
	}
	offset := params.Page * params.PageSize
	groupQB = groupQB.Limit(uint64(params.PageSize)).Offset(uint64(offset))

	groupQuery, groupArgs, _ := groupQB.ToSql()
	log.Printf("group query: %s", groupQuery)
	log.Printf("group query params: %+v", groupArgs)

	groupRows, err := repo.DB.Query(groupQuery, groupArgs...)
	if err != nil {
		return nil, err
	}
	defer groupRows.Close()

	groupList := make([]*SecretGroup, 0)
	for groupRows.Next() {
		group := &SecretGroup{}
		err = groupRows.Scan(&group.ID, &group.Description)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, group)
	}

	if len(groupList) == 0 {
		return groupList, nil
	}

	groupIDlist := make([]uint64, 0)
	groupMap := make(map[uint64]*SecretGroup, 0)
	for _, group := range groupList {
		groupIDlist = append(groupIDlist, group.ID)
		groupMap[group.ID] = group
	}

	secretQB := psql.Select("id, group_id, content").
		From("secrets").
		Where(sq.Eq{"group_id": groupIDlist})

	secretQuery, secretArgs, _ := secretQB.ToSql()
	log.Printf("group query: %s", secretQuery)
	log.Printf("group query params: %+v", secretArgs)

	secretRows, err := repo.DB.Query(secretQuery, secretArgs...)
	if err != nil {
		return nil, err
	}
	defer secretRows.Close()

	for secretRows.Next() {
		var secret Secret
		err = secretRows.Scan(&secret.ID, &secret.GroupdID, &secret.Content)
		if err != nil {
			return nil, err
		}

		group := groupMap[secret.GroupdID]
		group.Secrets = append(group.Secrets, secret)
	}
	return groupList, nil
}

func (repo *Repo) NewKey() (*Key, error) {
	log.Println("generate new key")
	query := `insert into keys(value, create_at) values($1, $2) returning id, value, create_at`

	key := &Key{}
	err := repo.DB.QueryRow(query, NewKeyString(), time.Now()).Scan(&key.ID, &key.Value, &key.CreateAt)

	if err != nil {
		log.Printf("NewKey error: %s", err.Error())
		return nil, err
	}

	return key, nil
}

func (repo *Repo) GetKeyById(id uint64) (*Key, error) {
	query := `select id, value, create_at from keys where id = $1`

	key := &Key{}
	err := repo.DB.QueryRow(query, id).Scan(&key.ID, &key.Value, &key.CreateAt)

	if err != nil {
		log.Printf("GetKeyById error: %s", err.Error())
		return nil, err
	}

	return key, nil
}

func (repo *Repo) GetKey() (*Key, error) {
	query := `select id, value, create_at from keys where id = (select max(id) from keys)`

	key := &Key{}
	err := repo.DB.QueryRow(query).Scan(&key.ID, &key.Value, &key.CreateAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return repo.NewKey()
		}
		log.Printf("GetKeyById error: %s", err.Error())
		return nil, err
	}

	if key.OutDated() {
		log.Println("key is outdated")
		return repo.NewKey()
	}

	return key, nil
}
