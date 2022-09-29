package entity

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID    `db:"id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Username  nulls.String `db:"username"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
