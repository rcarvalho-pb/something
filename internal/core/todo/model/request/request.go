package todo_request

import (
	user_request "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/request"
)

type TodoRequest struct {
	ID          uint32                      `json:"id,omitempty"`
	Name        string                      `json:"name,omitempty"`
	Description string                      `json:"description,omitempty"`
	Status      string                      `json:"status,omitempty"`
	Users       *[]user_request.UserRequest `json:"users"`
}
