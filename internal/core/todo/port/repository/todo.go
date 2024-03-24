package todo_irepository

import (
	todo_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/dto"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type TodoRepository interface {
	Save(*todo_dto.Todo, *[]user_dto.User) (uint32, error)
	Update(*todo_dto.Todo) (uint32, error)
	UpdateUsers(*todo_dto.Todo, *[]user_dto.User) error
	FindAllActive() (*[]todo_dto.Todo, error)
	FindAll() (*[]todo_dto.Todo, error, error)
	FindById(id uint32) (*todo_dto.Todo, error)
	FindByUserId(id uint32) (*[]todo_dto.Todo, error)
	DeleteById(id uint32) error
}
