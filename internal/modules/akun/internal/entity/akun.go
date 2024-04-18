package entity

import (
	"github.com/google/uuid"
)

type Akun struct {
	Id       uuid.UUID `db:"id"`
	Nama     string    `db:"nama"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Foto     string    `db:"foto"`
	Akses    int       `db:"akses"`
	Verified bool      `db:"verified"`
}
