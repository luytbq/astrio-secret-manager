package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-secret-manager/config"
	"github.com/luytbq/astrio-secret-manager/pkg/common"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	path := config.App.SERVER_API_PREFIX + "/api/v1/users"
	r.GET(path+"/infos", h.handleGetInfos)
}

func (h *Handler) handleGetInfos(c *gin.Context) {
	code, resBytes, err := common.RequestAAS("GET", "/users/infos", nil, &c.Request.Header)

	if err != nil {
		c.JSON(code, err)
		return
	}

	c.Status(code)
	if resBytes != nil {
		_, _ = c.Writer.Write(resBytes)
	}
}
