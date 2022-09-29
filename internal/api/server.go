package api

import (
	"github.com/etuchnap28/tasks-mangement-with-go/internal/config"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/core/user"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
)

type Service struct {
	user user.UserService
}

type Server struct {
	config  *config.ServerConfig
	router  *gin.Engine
	db      *pop.Connection
	service Service
}

func NewServer(cfg *config.ServerConfig, db *pop.Connection) (*Server, error) {
	server := &Server{
		config: cfg,
		db:     db,
	}
	server.setupServices()
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.register)

	server.router = router
}

func (server *Server) setupServices() {
	server.setupUserService()
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
