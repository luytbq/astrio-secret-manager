package api

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-secret-manager/pkg/common"
	"github.com/luytbq/astrio-secret-manager/pkg/secret"
	"github.com/luytbq/astrio-secret-manager/pkg/user"
)

type Server struct {
	Port   string
	Prefix string
	DB     *sql.DB
}

func NewServer(port string, prefix string, db *sql.DB) *Server {
	return &Server{
		Port:   port,
		Prefix: prefix,
		DB:     db,
	}
}

func (server *Server) Run() {
	engine := gin.Default()

	engine.Use(authMiddleware)

	repo := secret.NewRepo(server.DB)
	secretHandler := secret.NewHandler(repo)
	secretHandler.RegisterRoutes(engine)

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(engine)

	_ = engine.Run(":" + server.Port)

	// return err
}

func authMiddleware(c *gin.Context) {
	log.Printf("authMiddleware")
	code, resBytes, err := common.RequestAAS("GET", "/users/infos", nil, &c.Request.Header)

	if err != nil {
		c.AbortWithError(code, err)
		return
	}

	type AASInfoRes struct {
		UserID uint64 `json:"user_id"`
	}

	var aasVerifyRes *AASInfoRes

	json.Unmarshal(resBytes, &aasVerifyRes)

	c.Set("user_id", aasVerifyRes.UserID)

	c.Next()
}
