package todo_request

import (
	todo_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/dto"
	user_request "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/request"
)

type TodoRequest struct {
	ID          uint32                     `json:"id,omitempty"`
	Name        string                     `json:"name,omitempty"`
	Description string                     `json:"description,omitempty"`
	Status      string                     `json:"status,omitempty"`
	Users       []user_request.UserRequest `json:"users,omitempty"`
}

func (t *TodoRequest) ToTodoDTO() *todo_dto.Todo {
	return &todo_dto.Todo{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
	}
}
