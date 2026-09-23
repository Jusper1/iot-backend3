package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	repository "iot-backend/internal/repositories/master"
)

var ErrInvalidPIC = errors.New("data PIC tidak valid")

type PICService struct {
	Repo *repository.PICRepository
}

func NewPICService(
	repo *repository.PICRepository,
) *PICService {
	return &PICService{
		Repo: repo,
	}
}

func (s *PICService) Create(data *models.MasterPIC) error {
	data.NamaPIC = strings.TrimSpace(data.NamaPIC)
	if data.Email != nil {
	value := strings.TrimSpace(*data.Email)
	data.Email = &value
}

if data.NIK != nil {
	value := strings.TrimSpace(*data.NIK)
	data.NIK = &value
}

if data.NoHP != nil {
	value := strings.TrimSpace(*data.NoHP)
	data.NoHP = &value
}
	if data.NamaPIC == "" {
		return ErrInvalidPIC
	}

	return s.Repo.Create(data)
}

func (s *PICService) FindAll() ([]models.MasterPIC, error) {
	return s.Repo.FindAll()
}

func (s *PICService) FindByID(id uint) (*models.MasterPIC, error) {
	if id == 0 {
		return nil, ErrInvalidPIC
	}

	return s.Repo.FindByID(id)
}

func (s *PICService) FindByInstansiID(instansiID uint) ([]models.MasterPIC, error) {
	if instansiID == 0 {
		return nil, ErrInvalidPIC
	}

	return s.Repo.FindByInstansiID(instansiID)
}

func (s *PICService) Update(data *models.MasterPIC) error {
	if data.ID == 0 {
		return ErrInvalidPIC
	}

	data.NamaPIC = strings.TrimSpace(data.NamaPIC)
	if data.Email != nil {
	value := strings.TrimSpace(*data.Email)
	data.Email = &value
}

if data.NIK != nil {
	value := strings.TrimSpace(*data.NIK)
	data.NIK = &value
}

if data.NoHP != nil {
	value := strings.TrimSpace(*data.NoHP)
	data.NoHP = &value
}

	if data.NamaPIC == "" {
		return ErrInvalidPIC
	}

	return s.Repo.Update(data)
}

func (s *PICService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidPIC
	}

	return s.Repo.Delete(id)
}