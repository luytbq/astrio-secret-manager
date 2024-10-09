package secret

import "time"

type User struct {
	ID uint64
}

type SecretGroup struct {
	ID          uint64   `json:"id"`
	UserID      uint64   `json:"-"`
	Description string   `json:"description"`
	Secrets     []Secret `json:"secrets"`
}

type Secret struct {
	ID          uint64 `json:"id"`
	GroupdID    uint64 `json:"-"`
	Content     string `json:"content"`
	Description string `json:"description"`
	KeyID       uint64 `json:"-"`
	Nonce       string `json:"-"`
	Encrypt     bool   `json:"encrypt"`
}

type Key struct {
	ID       uint64
	Value    string
	CreateAt time.Time
}

type GetSecretsParams struct {
	UserID   uint64 `form:"-"`
	Keyword  string `form:"keyword"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

type GetSecretsResponse struct {
	List []SecretGroup
}
