package user_service

import (
	user_request "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/request"
	user_response "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/response"
)

type UserService interface {
	CreateUser(*user_request.UserRequest) *user_response.UserResponse
	UpdateUser(*user_request.UserRequest) *user_response.UserResponse
	UpdateUserPassword(*user_request.UserRequest) *user_response.UserResponse
	FindAllActiveUsers() *user_response.UserResponse
	FindAllUsers() *user_response.UserResponse
	FindUserById(*user_request.UserRequest) *user_response.UserResponse
	DeleteUserById(*user_request.UserRequest) *user_response.UserResponse
}
