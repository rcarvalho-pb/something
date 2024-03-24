package user_response

import response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"

type UserResponse struct {
	StatusCode response_status.StatusCode `json:"-"`
	Status     response_status.StatusMsg  `json:"status"`
	Content    any                        `json:"content"`
}
