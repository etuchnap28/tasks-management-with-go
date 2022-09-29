package api

import (
	"github.com/etuchnap28/tasks-mangement-with-go/internal/core/user"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) setupUserService() {
	userRepoFactory := persistence.NewRepositoryFactory[persistence.UserRepository](server.db)
	server.service.user = user.NewUserService(user.WithUserRepositoryFactory(userRepoFactory))
}

type registerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userSvc := server.service.user
	arg := user.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	resStr, err := userSvc.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	result := struct {
		message string
	}{
		message: resStr,
	}
	ctx.JSON(http.StatusOK, result)
}
