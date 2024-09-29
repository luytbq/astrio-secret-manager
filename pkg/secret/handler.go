package secret

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-secret-manager/config"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	path := config.App.SERVER_API_PREFIX + "/api/v1/secrets"
	r.GET(path, h.handleQuery)
	r.POST(path, h.handleCreate)
}

func (h *Handler) handleCreate(c *gin.Context) {
	var group SecretGroup
	err := c.BindJSON(&group)

	if err != nil {
		log.Println("error ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	cUserId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}
	group.UserID = cUserId.(uint64)

	err = h.repo.InsertSecretGroup(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": group.ID})
}

func (h *Handler) handleQuery(c *gin.Context) {
	var params GetSecretsParams
	_ = c.BindQuery(&params)

	cUserId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	params.UserID = cUserId.(uint64)

	if params.PageSize < 1 {
		params.PageSize = 10
	}

	secretGroups, err := h.repo.GetSecrets(params)

	if err != nil {
		log.Println("error ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	total, err := h.repo.GetSecretsTotal(params)

	if err != nil {
		log.Println("error ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(200, gin.H{"total": total, "list": secretGroups})

}
