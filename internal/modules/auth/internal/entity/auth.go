package entity

import "github.com/google/uuid"

type Auth struct {
	Id       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Akses    int       `db:"akses"`
}
