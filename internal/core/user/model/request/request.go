package user_request

import (
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type UserRequest struct {
	ID          uint32 `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

func (u *UserRequest) ToUserDTO() *user_dto.User {
	return &user_dto.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}
