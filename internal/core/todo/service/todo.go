package todo_service

import (
	todo_request "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/model/request"
	todo_response "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/model/response"
	todo_irepository "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/port/repository"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
)

type todoService struct {
	todo_irepository.TodoRepository
}

func NewTodoRepository(todoRepository todo_irepository.TodoRepository) *todoService {
	return &todoService{
		TodoRepository: todoRepository,
	}
}

func (t *todoService) CreateTodo(todoRequest *todo_request.TodoRequest) *todo_response.TodoResponse {
	todo := todoRequest.ToTodoDTO()
	var users []user_dto.User
	for _, user := range todoRequest.Users {
		users = append(users, *user.ToUserDTO())
	}
	id, err := t.TodoRepository.Save(todo, users)
	if err != nil {
		return &todo_response.TodoResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}

	return &todo_response.TodoResponse{
		StatusCode: response_status.Created,
		Status:     response_status.SuccessMsg,
		Content:    id,
	}
}

func (t *todoService) UpdateTodo(TodoRequest *todo_request.TodoRequest) *todo_response.TodoResponse {
	return nil
}

func (t *todoService) UpdateTodoUsers(todoRequest *todo_request.TodoRequest) *todo_response.TodoResponse {
	todo := todoRequest.ToTodoDTO()
	var users []user_dto.User
	if len(todoRequest.Users) > 1 {
		for _, user := range todoRequest.Users {
			users = append(users, *user.ToUserDTO())
		}
	}

	if err := t.TodoRepository.UpdateUsers(todo.ID, users); err != nil {
		return &todo_response.TodoResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}
	if err := t.TodoRepository.UpdateUsers(todo.ID, users); err != nil {
		return &todo_response.TodoResponse{
			StatusCode: response_status.InternalError,
			Status:     response_status.InternalErrMsg,
			Content:    err,
		}
	}
	return &todo_response.TodoResponse{
		StatusCode: response_status.Ok,
		Status:     response_status.SuccessMsg,
		Content:    nil,
	}
}

func (t *todoService) FindAllActiveTodos(todoRequest *todo_request.TodoRequest) *todo_response.TodoResponse {
	return nil
}

func (t *todoService) FindAllTodos() *todo_response.TodoResponse {
	return nil
}

func (t *todoService) FindByUserId(id uint32) *todo_response.TodoResponse {
	return nil
}

func (t *todoService) FindTodoById(id uint32) *todo_response.TodoResponse {
	return nil
}

func (t *todoService) DeleteTodoById(id uint32) *todo_response.TodoResponse {
	return nil
}
