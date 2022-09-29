package entity

import (
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID    `db:"id"`
	UserId      uuid.UUID    `db:"user_id"`
	Name        string       `db:"name"`
	Description nulls.String `db:"description"`
	IsDone      bool         `db:"is_done"`
	DueDate     time.Time    `db:"due_date"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
}
