package todo_controller

import (
	"encoding/json"
	"io"
	"net/http"

	todo_request "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/model/request"
	todo_iservice "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/port/service"
	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
	json_response "github.com/rcarvalho-pb/todo-app-go/internal/infra/response"
)

type TodoController struct {
	todo_iservice.TodoService
}

func NewTodoController(todoService todo_iservice.TodoService) *TodoController {
	return &TodoController{
		TodoService: todoService,
	}
}

func (t *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		json_response.ERROR(w, response_status.BadRequest, err)
		return
	}
	var todoRequest todo_request.TodoRequest
	if err = json.Unmarshal(body, &todoRequest); err != nil {
		json_response.ERROR(w, response_status.UnprocessableEntity, err)
		return
	}
	todoResponse := t.TodoService.CreateTodo(&todoRequest)
	if todoResponse.StatusCode != response_status.Created {
		json_response.ERROR(w, todoResponse.StatusCode, todoResponse.Content.(error))
		return
	}
	json_response.JSON(w, todoResponse.StatusCode, todoResponse)

}

func (t *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) UpdateTodoUsers(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) FindById(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) FindAllActiveTodos(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) FindAllTodos(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) FindByUserId(w http.ResponseWriter, r *http.Request) {

}

func (t *TodoController) DeleteById(w http.ResponseWriter, r *http.Request) {

}
