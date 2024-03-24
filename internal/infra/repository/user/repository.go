package user_irepository

import (
	"log"

	"github.com/jmoiron/sqlx"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type userRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Save(user *user_dto.User) (uint32, error) {
	res, err := u.NamedExec("INSERT INTO users (first_name, last_name, email, password) VALUES (:first_name, :last_name, :email, :password)", user)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}

func (u *userRepository) Update(user *user_dto.User) (uint32, error) {
	log.Println(user)
	res, err := u.NamedExec("UPDATE users SET first_name = :first_name, last_name = :last_name, email = :email WHERE id = :id", &user)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (u *userRepository) UpdatePassword(user *user_dto.User) error {
	_, err := u.NamedExec("UPDATE users SET password = :password WHERE id = :id", &user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindAll() (*[]user_dto.User, error) {
	var users []user_dto.User
	if err := u.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepository) FindAllActive() (*[]user_dto.User, error) {
	var users []user_dto.User
	if err := u.Select(&users, "SELECT * FROM users WHERE active = TRUE"); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepository) FindById(id uint32) (*user_dto.User, error) {
	var user user_dto.User
	if err := u.Get(&user, "SELECT * FROM users WHERE id = $1 AND active = TRUE", id); err != nil {
		return &user_dto.User{}, err
	}

	return &user, nil
}

func (u *userRepository) DeleteById(id uint32) error {
	tx := u.MustBegin()

	tx.MustExec("UPDATE users SET active = FALSE WHERE id = $1", id)
	tx.MustExec("DELETE FROM todos_users WHERE user_id = $1", id)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
