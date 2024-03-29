package todo_controller

import todo_iservice "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/port/service"

type todoController struct {
	todo_iservice.TodoService
}

func NewTodoController(todoService todo_iservice.TodoService) *todoController {
	return &todoController{
		TodoService: todoService,
	}
}
