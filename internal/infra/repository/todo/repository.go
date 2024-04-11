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
	var id int64
	tx := t.MustBegin()
	res, err := tx.NamedExec("INSERT INTO todos (name, description) VALUES (:name, :description)", todo)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	if users != nil {
		for _, user := range users {
			tx.MustExec("INSERT INTO todos_users VALUES ($1, $2)", id, user.ID)
		}
	}
	tx.Commit()
	return uint32(id), nil
}

func (t *todoRepository) Update(todo *todo_dto.Todo) error {
	_, err := t.NamedExec("UPDATE todos SET name=:name, description=:description WHERE id = :id", todo)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateUsers(id uint32, users []user_dto.User) error {
	tx := t.MustBegin()
	tx.MustExec("DELETE FROM todos_users WHERE todo_id = $1", id)
	for _, user := range users {
		tx.MustExec("INSERT INTO todos_users (todo_id, user_id) VALUES ($1, $2)", id, user.ID)
	}

	err := tx.Commit()
	return err
}

func (t *todoRepository) FindAllActive() ([]todo_dto.Todo, error) {
	var todos []todo_dto.Todo

	if err := t.Select(&todos, "SELECT * FROM todos WHERE active = TRUE"); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) FindAll() ([]todo_dto.Todo, error) {
	var todos []todo_dto.Todo

	if err := t.Select(&todos, "SELECT * FROM todos"); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) FindById(id uint32) (*todo_dto.Todo, error) {
	var todo todo_dto.Todo
	if err := t.Get(&todo, "SELECT * FROM todos WHERE id = $1", id); err != nil {
		return &todo_dto.Todo{}, err
	}
	return &todo, nil
}

func (t *todoRepository) FindByUserId(id uint32) ([]todo_dto.Todo, error) {
	var todos []todo_dto.Todo
	if err := t.Select(&todos, "SELECT * FROM todos t JOIN ON todos_users tu WHERE tu.user_id = $1", id); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) DeleteById(id uint32) error {
	if _, err := t.Exec("UPDATE todos SET active = FALSE WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
