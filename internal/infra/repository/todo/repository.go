package todo_repository

import (
	"github.com/jmoiron/sqlx"
	todo_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/dto"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type todoRepository struct {
	*sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *todoRepository {
	return &todoRepository{
		DB: db,
	}
}

func (t *todoRepository) Save(todo *todo_dto.Todo, users []user_dto.User) (uint32, error) {
	id := 0
	tx := t.MustBegin()
	res, err := tx.NamedExec("INSERT INTO todos (name, description) VALUES (:name, :description)", todo)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if users
}

func (t *todoRepository) Update(todo *todo_dto.Todo) error {
	return nil
}

func (t *todoRepository) UpdateUsers(id uint32, users []user_dto.User) error {
	return nil
}

func (t *todoRepository) FindAllActive() ([]todo_dto.Todo, error) {
	return nil, nil
}

func (t *todoRepository) FindAll() ([]todo_dto.Todo, error) {
	return nil, nil
}

func (t *todoRepository) FindById(id uint32) ([]todo_dto.Todo, error) {
	return nil, nil
}

func (t *todoRepository) FindByUserId(id uint32) ([]todo_dto.Todo, error) {
	return nil, nil
}

func (t *todoRepository) DeleteById(id uint32) error {
	return nil
}
