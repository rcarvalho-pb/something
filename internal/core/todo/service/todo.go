package todo_service

import todo_irepository "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/port/repository"

type todoService struct {
	todo_irepository.TodoRepository
}

func NewTodoRepository(todoRepository todo_irepository.TodoRepository) *todoService {
	return &todoService{
		TodoRepository: todoRepository,
	}
}
