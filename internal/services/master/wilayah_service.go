package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	repository "iot-backend/internal/repositories/master"
)

var ErrInvalidWilayah = errors.New("data wilayah tidak valid")

type WilayahService struct {
	Repo *repository.WilayahRepository
}

func NewWilayahService(
	repo *repository.WilayahRepository,
) *WilayahService {
	return &WilayahService{
		Repo: repo,
	}
}

func (s *WilayahService) Create(data *models.MasterWilayah) error {
	data.Provinsi = strings.TrimSpace(data.Provinsi)
	data.KotaKab = strings.TrimSpace(data.KotaKab)

	if data.Provinsi == "" || data.KotaKab == "" {
		return ErrInvalidWilayah
	}

	return s.Repo.Create(data)
}

func (s *WilayahService) FindAll() ([]models.MasterWilayah, error) {
	return s.Repo.FindAll()
}

func (s *WilayahService) FindByID(id uint) (*models.MasterWilayah, error) {
	if id == 0 {
		return nil, ErrInvalidWilayah
	}

	return s.Repo.FindByID(id)
}

func (s *WilayahService) FindByProvinsi(provinsi string) ([]models.MasterWilayah, error) {
	provinsi = strings.TrimSpace(provinsi)

	if provinsi == "" {
		return nil, ErrInvalidWilayah
	}

	return s.Repo.FindByProvinsi(provinsi)
}

func (s *WilayahService) Update(data *models.MasterWilayah) error {
	if data.ID == 0 {
		return ErrInvalidWilayah
	}

	data.Provinsi = strings.TrimSpace(data.Provinsi)
	data.KotaKab = strings.TrimSpace(data.KotaKab)

	if data.Provinsi == "" || data.KotaKab == "" {
		return ErrInvalidWilayah
	}

	return s.Repo.Update(data)
}

func (s *WilayahService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidWilayah
	}

	return s.Repo.Delete(id)
}