package user

import (
	"context"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/core/entity"
	"github.com/etuchnap28/tasks-mangement-with-go/internal/persistence"
)

type UserService interface {
	Register(ctx context.Context, arg RegisterParams) (string, error)
}

type RegisterParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (svc *userService) Register(ctx context.Context, arg RegisterParams) (string, error) {
	var result string

	err := svc.userRf.WithTransaction(
		ctx, func(repo *persistence.UserRepository) error {
			err := repo.InsertUser(
				&entity.User{
					FirstName: arg.FirstName,
					LastName:  arg.LastName,
					Email:     arg.Email,
					Password:  arg.Password,
				},
			)
			if err != nil {
				return err
			}
			result = "Registered"
			return nil
		},
	)

	return result, err
}
