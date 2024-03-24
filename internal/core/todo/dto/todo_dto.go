package todo_dto

import "time"

var Status = []string{
	"to do",
	"doing",
	"done",
}

type Todo struct {
	ID               uint32    `db:"id"`
	Name             string    `db:"name"`
	Description      string    `db:"description"`
	Status           string    `db:"status"`
	Active           bool      `db:"active"`
	CreatedAt        time.Time `db:"created_at"`
	LastModifiedDate time.Time `db:"last_modified_date"`
}

func (t *Todo) isValidStatus() bool {
	for _, status := range Status {
		if t.Status == status {
			return true
		}
	}

	return false
}
