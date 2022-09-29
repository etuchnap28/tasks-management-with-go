package user

import "github.com/etuchnap28/tasks-mangement-with-go/internal/persistence"

type userService struct {
	userRf *persistence.RepositoryFactory[persistence.UserRepository]
}

type userServiceOpts func(service *userService)

func NewUserService(opts ...userServiceOpts) UserService {
	svc := &userService{}
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func WithUserRepositoryFactory(rf *persistence.RepositoryFactory[persistence.UserRepository]) userServiceOpts {
	return func(svc *userService) {
		svc.userRf = rf
	}
}
