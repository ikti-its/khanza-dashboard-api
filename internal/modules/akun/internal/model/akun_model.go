package model

type CreateAkunRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Foto     string `json:"foto"`
	Akses    int    `json:"akses" validate:"required,numeric,oneof=1000 2000"`
	Verified bool   `json:"verified" validate:"boolean"`
}

type UpdateAkunRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Foto     string `json:"foto" validate:"required"`
	Verified bool   `json:"verified" validate:"boolean"`
}

type AkunResponse struct {
	Id       string `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Foto     string `json:"foto"`
	Akses    int    `json:"akses"`
	Verified bool   `json:"verified"`
}

type AkunPageResponse struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
	Akun  []AkunResponse `json:"akun"`
}
