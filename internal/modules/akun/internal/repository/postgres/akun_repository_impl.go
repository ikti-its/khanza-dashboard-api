package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/entity"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type akunRepositoryImpl struct {
	DB *sqlx.DB
}

func NewAkunRepository(db *sqlx.DB) repository.AkunRepository {
	return &akunRepositoryImpl{db}
}

func (r *akunRepositoryImpl) Insert(akun *entity.Akun) error {
	query := `
		INSERT INTO akun (id, nama, email, password, foto, akses, verified) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.DB.Exec(query, akun.Id, akun.Nama, akun.Email, akun.Password, akun.Foto, akun.Akses, akun.Verified)

	return err
}

func (r *akunRepositoryImpl) Find() ([]entity.Akun, error) {
	query := `
		SELECT id, nama, email, password, foto, akses, verified
		FROM akun
	`

	var records []entity.Akun
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *akunRepositoryImpl) FindPage(page, size int) ([]entity.Akun, int, error) {
	query := `
		SELECT id, nama, email, password, foto, akses, verified
		FROM akun
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM akun"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Akun
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *akunRepositoryImpl) FindById(id uuid.UUID) (entity.Akun, error) {
	query := `
		SELECT id, nama, email, password, foto, akses, verified
		FROM akun
		WHERE id = $1
	`

	var record entity.Akun
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *akunRepositoryImpl) Update(akun *entity.Akun) error {
	query := `
		UPDATE akun 
		SET nama = $2, email = $3, password = $4, foto = $5, akses = $6, verified = $7, updated_at = $8
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, akun.Id, akun.Nama, akun.Email, akun.Password, akun.Foto, akun.Akses, akun.Verified, time.Now())

	return err
}

func (r *akunRepositoryImpl) Delete(akun *entity.Akun) error {
	query := `
		DELETE FROM akun
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, akun.Id)

	return err
}
