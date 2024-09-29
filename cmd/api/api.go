package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-secret-manager/pkg/secret"
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
	handler := secret.NewHandler(repo)

	handler.RegisterRoutes(engine)

	_ = engine.Run(":" + server.Port)

	// return err
}

func authMiddleware(c *gin.Context) {
	authHeaders := c.Request.Header["Authorization"]
	if len(authHeaders) != 1 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid header"))
		return
	}

	token := strings.ReplaceAll(authHeaders[0], "Bearer ", "")

	log.Println("token: " + token)
	if token == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	aasUrl := "http://localhost:8000/auth/api/v1/users/verify"
	req, err := http.NewRequest("GET", aasUrl, nil)

	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("internal server error"))
		return
	}

	req.Header.Set("Astrio-Auth-Token", authHeaders[0])

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Printf("Error making request: %v", err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("internal server error"))
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Println("res: " + string(body))

	if err != nil {
		log.Println("Error reading response body: %v", err)
	}

	log.Printf("AAS verify status code: %d", res.StatusCode)

	if res.StatusCode == http.StatusBadRequest {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if res.StatusCode == http.StatusInternalServerError {
		c.AbortWithError(http.StatusBadGateway, errors.New("unauthorized"))
		return
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error making request: %v", err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("internal server error"))
		return
	}

	type AASVerifyRes struct {
		UserID uint64 `json:"user_id"`
	}

	var aasVerifyRes *AASVerifyRes

	json.Unmarshal(body, &aasVerifyRes)

	c.Set("user_id", aasVerifyRes.UserID)

	c.Next()
}
