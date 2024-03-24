package user_dto

import "time"

type User struct {
	ID               uint32    `db:"id"`
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	Role             string    `db:"role"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	Active           bool      `db:"active"`
	CreatedAt        time.Time `db:"created_at"`
	LastModifiedDate time.Time `db:"last_modified_date"`
}
