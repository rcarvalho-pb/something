package user_repository

import (
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type UserRepository interface {
	Save(*user_dto.User) (uint32, error)
	Update(*user_dto.User) (uint32, error)
	UpdatePassword(*user_dto.User) error
	FindAllActive() ([]user_dto.User, error)
	FindAll() ([]user_dto.User, error)
	FindById(id uint32) (*user_dto.User, error)
	DeleteById(id uint32) error
}
