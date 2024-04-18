package postgres

import (
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/auth/internal/entity"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/auth/internal/repository"
	"github.com/jmoiron/sqlx"
)

type authRepositoryImpl struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) repository.AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) FindByEmail(email string) (entity.Auth, error) {
	query := `
		SELECT id, email, password, akses
		FROM akun
		WHERE email = $1
	`

	var record entity.Auth
	err := r.DB.Get(&record, query, email)

	return record, err
}
