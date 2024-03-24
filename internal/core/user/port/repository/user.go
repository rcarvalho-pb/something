package user_repository

import (
	"errors"

	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

var (
	DuplicateUser = errors.New("duplicate user")
	UserNotFound  = errors.New("user not found")
)

type UserRepository interface {
	Save(*user_dto.User) (uint32, error)
	Update(*user_dto.User) (uint32, error)
	UpdatePassword(*user_dto.User) error
	FindAllActive() (*[]user_dto.User, error)
	FindAll() (*[]user_dto.User, error)
	FindById(id uint32) (*user_dto.User, error)
	DeleteById(id uint32) error
}
