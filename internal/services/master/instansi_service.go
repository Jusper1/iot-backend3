package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	repository "iot-backend/internal/repositories/master"
)

var ErrInvalidInstansi = errors.New("data instansi tidak valid")

type InstansiService struct {
	Repo *repository.InstansiRepository
}

func NewInstansiService(
	repo *repository.InstansiRepository,
) *InstansiService {
	return &InstansiService{
		Repo: repo,
	}
}

func (s *InstansiService) Create(data *models.MasterInstansi) error {
	data.NamaInstansi = strings.TrimSpace(data.NamaInstansi)

	if data.NamaInstansi == "" {
		return ErrInvalidInstansi
	}

	return s.Repo.Create(data)
}

func (s *InstansiService) FindAll() ([]models.MasterInstansi, error) {
	return s.Repo.FindAll()
}

func (s *InstansiService) FindByID(id uint) (*models.MasterInstansi, error) {
	if id == 0 {
		return nil, ErrInvalidInstansi
	}

	return s.Repo.FindByID(id)
}

func (s *InstansiService) Search(keyword string) ([]models.MasterInstansi, error) {
	return s.Repo.Search(keyword)
}

func (s *InstansiService) Update(data *models.MasterInstansi) error {
	if data.ID == 0 {
		return ErrInvalidInstansi
	}

	data.NamaInstansi = strings.TrimSpace(data.NamaInstansi)

	if data.NamaInstansi == "" {
		return ErrInvalidInstansi
	}

	return s.Repo.Update(data)
}

func (s *InstansiService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidInstansi
	}

	return s.Repo.Delete(id)
}