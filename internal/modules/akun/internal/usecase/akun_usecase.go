package usecase

import (
	"fmt"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/config"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/exception"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/helper"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/entity"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/model"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/repository"
)

type AkunUseCase struct {
	Repository repository.AkunRepository
	Config     *config.Config
}

func NewAkunUseCase(repository *repository.AkunRepository, cfg *config.Config) *AkunUseCase {
	return &AkunUseCase{
		Repository: *repository,
		Config:     cfg,
	}
}

func (u *AkunUseCase) Create(request *model.CreateAkunRequest) model.AkunResponse {
	if role := request.Akses; role == 1000 {
		panic(&exception.ForbiddenError{
			Message: "Not allowed to create this role",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err, "Failed to encrypt password")

	if request.Foto == "" {
		request.Foto = fmt.Sprintf("%s/file/img/default.png", u.Config.Get("APP_URL", "http://localhost:8080/v1"))
	}

	akun := entity.Akun{
		Id:       helper.MustNew(),
		Nama:     request.Nama,
		Email:    request.Email,
		Password: string(encrypted),
		Foto:     request.Foto,
		Akses:    request.Akses,
		Verified: request.Verified,
	}

	if err := u.Repository.Insert(&akun); err != nil {
		exception.PanicIfError(err, "Failed to insert akun")
	}

	response := model.AkunResponse{
		Id:       akun.Id.String(),
		Nama:     akun.Nama,
		Email:    akun.Email,
		Foto:     akun.Foto,
		Akses:    akun.Akses,
		Verified: akun.Verified,
	}

	return response
}

func (u *AkunUseCase) Get() []model.AkunResponse {
	akun, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all akun")

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			Id:       akun.Id.String(),
			Nama:     akun.Nama,
			Email:    akun.Email,
			Foto:     akun.Foto,
			Akses:    akun.Akses,
			Verified: akun.Verified,
		}
	}

	return response
}

func (u *AkunUseCase) GetPage(page, size int) model.AkunPageResponse {
	akun, total, err := u.Repository.FindPage(page, size)
	exception.PanicIfError(err, "Failed to get paged akun")

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			Id:       akun.Id.String(),
			Nama:     akun.Nama,
			Email:    akun.Email,
			Foto:     akun.Foto,
			Akses:    akun.Akses,
			Verified: akun.Verified,
		}
	}

	pagedResponse := model.AkunPageResponse{
		Page:  page,
		Size:  size,
		Total: total,
		Akun:  response,
	}

	return pagedResponse
}

func (u *AkunUseCase) GetById(id string) model.AkunResponse {
	akun, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	response := model.AkunResponse{
		Id:       akun.Id.String(),
		Nama:     akun.Nama,
		Email:    akun.Email,
		Foto:     akun.Foto,
		Akses:    akun.Akses,
		Verified: akun.Verified,
	}

	return response
}

func (u *AkunUseCase) Update(request *model.UpdateAkunRequest, id string) model.AkunResponse {
	akun, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	akun.Id = helper.MustParse(id)
	akun.Nama = request.Nama
	akun.Email = request.Email
	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		exception.PanicIfError(err, "Failed to encrypt password")

		akun.Password = string(encrypted)
	}
	akun.Foto = request.Foto
	akun.Verified = request.Verified

	if err := u.Repository.Update(&akun); err != nil {
		exception.PanicIfError(err, "Failed to update akun")
	}

	response := model.AkunResponse{
		Id:       akun.Id.String(),
		Nama:     akun.Nama,
		Email:    akun.Email,
		Foto:     akun.Foto,
		Akses:    akun.Akses,
		Verified: akun.Verified,
	}

	return response
}

func (u *AkunUseCase) Delete(id string) {
	akun, err := u.Repository.FindById(helper.MustParse(id))
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if err := u.Repository.Delete(&akun); err != nil {
		exception.PanicIfError(err, "Failed to delete akun")
	}
}
