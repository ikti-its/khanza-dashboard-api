package repository

import "github.com/ikti-its/khanza-dashboard-api/internal/modules/auth/internal/entity"

type AuthRepository interface {
	FindByEmail(email string) (entity.Auth, error)
}
