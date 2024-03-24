package user_controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
	user_request "github.com/rcarvalho-pb/todo-app-go/internal/core/user/model/request"
	user_service "github.com/rcarvalho-pb/todo-app-go/internal/core/user/port/service"
	json_response "github.com/rcarvalho-pb/todo-app-go/internal/infra/response"
)

type UserController struct {
	user_service.UserService
}

func NewUserControler(userService user_service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	var userRequest user_request.UserRequest
	if err = json.Unmarshal(body, &userRequest); err != nil {
		json_response.ERROR(w, response_status.UnprocessableEntity, err)
		return
	}
	log.Println(userRequest)

	userResponse := u.UserService.CreateUser(&userRequest)
	log.Println(userResponse)
	if userResponse.StatusCode != response_status.Created {
		json_response.ERROR(w, userResponse.StatusCode, userResponse.Content.(error))
		return
	}

	json_response.JSON(w, userResponse.StatusCode, userResponse)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	var userRequest user_request.UserRequest

	if err = json.Unmarshal(body, &userRequest); err != nil {
		json_response.ERROR(w, response_status.UnprocessableEntity, err)
		return
	}

	userRequest.ID = uint32(id)

	userResponse := u.UserService.UpdateUser(&userRequest)
	if userResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, userResponse.StatusCode, userResponse.Content.(error))
		return
	}

	json_response.JSON(w, userResponse.StatusCode, userResponse)
}

func (u *UserController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	var userRequest user_request.UserRequest
	if err = json.Unmarshal(body, &userRequest); err != nil {
		json_response.ERROR(w, response_status.UnprocessableEntity, err)
		return
	}

	userRequest.ID = uint32(id)

	userResponse := u.UserService.UpdateUserPassword(&userRequest)
	if userResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, userResponse.StatusCode, err)
		return
	}

	json_response.JSON(w, userResponse.StatusCode, userResponse)
}

func (u *UserController) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	usersResponse := u.UserService.FindAllUsers()
	if usersResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, usersResponse.StatusCode, usersResponse.Content.(error))
		return
	}

	json_response.JSON(w, usersResponse.StatusCode, usersResponse)
}

func (u *UserController) FindAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	usersResponse := u.UserService.FindAllActiveUsers()
	if usersResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, usersResponse.StatusCode, usersResponse.Content.(error))
		return
	}

	json_response.JSON(w, usersResponse.StatusCode, usersResponse)
}

func (u *UserController) FindUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	userRequest := &user_request.UserRequest{ID: uint32(id)}

	userResponse := u.UserService.FindUserById(userRequest)
	if userResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, userResponse.StatusCode, err)
		return
	}

	json_response.JSON(w, userResponse.StatusCode, userResponse)
}

func (u *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}

	userRequest := &user_request.UserRequest{
		ID: uint32(id),
	}

	userResponse := u.UserService.DeleteUserById(userRequest)
	if userResponse.StatusCode != response_status.Ok {
		json_response.ERROR(w, userResponse.StatusCode, userResponse.Content.(error))
		return
	}

	json_response.JSON(w, userResponse.StatusCode, userResponse)
}
