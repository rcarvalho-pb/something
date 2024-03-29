package todo_iservice

import (
	todo_request "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/model/request"
	todo_response "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/model/response"
)

type TodoService interface {
	CreateTodo(*todo_request.TodoRequest) *todo_response.TodoResponse
	UpdateTodo(*todo_request.TodoRequest) *todo_response.TodoResponse
	UpdateTodoUsers(*todo_request.TodoRequest) *todo_response.TodoResponse
	FindAllActiveTodos() *todo_response.TodoResponse
	FindAllTodos() *todo_response.TodoResponse
	FindTodoById(uint32) *todo_response.TodoResponse
	DeleteTodoById(uint32) *todo_response.TodoResponse
}
