package persistence

import (
	"github.com/etuchnap28/tasks-mangement-with-go/internal/core/entity"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

type UserRepository struct {
	db *pop.Connection
}

func NewUserRepository(db *pop.Connection) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) ListUsers() ([]*entity.User, error) {
	var users []*entity.User
	err := repo.db.All(users)
	return users, err
}

func (repo *UserRepository) FindByID(id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := repo.db.Where("id = ?", id).First(&user)
	return &user, err
}

func (repo *UserRepository) InsertUser(user *entity.User) error {
	return repo.db.Create(user)
}
