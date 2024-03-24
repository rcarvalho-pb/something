package user_iservice

import (
	"log"

	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
	user_request "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/request"
	user_response "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/response"
	user_repository "github.com/rcarvalho-pb/todo-app-go/internal/core/user/port/repository"
	user_service "github.com/rcarvalho-pb/todo-app-go/internal/core/user/port/service"
	encoder "github.com/rcarvalho-pb/todo-app-go/internal/infra/password_encoder"
)

type userService struct {
	user_repository.UserRepository
}

func NewUserService(userRepository user_repository.UserRepository) user_service.UserService {
	return &userService{
		UserRepository: userRepository,
	}
}

func (u *userService) CreateUser(userRequest *user_request.UserRequest) *user_response.UserResponse {
	user := userRequest.ToUserDTO()
	encodedPassword, err := encoder.HashPassword(user.Password)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	user.Password = encodedPassword

	id, err := u.UserRepository.Save(user)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Created,
		Status:     response_status.SuccessMsg,
		Content:    id,
	}
}

func (u *userService) UpdateUser(userRequest *user_request.UserRequest) *user_response.UserResponse {
	user := userRequest.ToUserDTO()
	id, err := u.UserRepository.Update(user)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    id,
	}
}

func (u *userService) UpdateUserPassword(userRequest *user_request.UserRequest) *user_response.UserResponse {
	log.Println(userRequest)
	user, err := u.UserRepository.FindById(userRequest.ID)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.NotFound,
			Status:     response_status.NotFoundErrMsg,
			Content:    err,
		}
	}

	if !encoder.CheckPassword(userRequest.Password, user.Password) {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    nil,
		}
	}

	encodedPassword, err := encoder.HashPassword(userRequest.NewPassword)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	userRequest.Password = encodedPassword

	if err := u.UserRepository.UpdatePassword(userRequest.ToUserDTO()); err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    nil,
	}
}

func (u *userService) FindAllActiveUsers() *user_response.UserResponse {
	users, err := u.UserRepository.FindAllActive()
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    users,
	}
}

func (u *userService) FindAllUsers() *user_response.UserResponse {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    users,
	}
}

func (u *userService) FindUserById(userRequest *user_request.UserRequest) *user_response.UserResponse {
	user, err := u.UserRepository.FindById(userRequest.ID)
	if err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    user,
	}
}

func (u *userService) DeleteUserById(userRequest *user_request.UserRequest) *user_response.UserResponse {
	if err := u.UserRepository.DeleteById(userRequest.ID); err != nil {
		return &user_response.UserResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}
	return &user_response.UserResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    nil,
	}
}
