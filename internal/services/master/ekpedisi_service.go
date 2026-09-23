package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	repository "iot-backend/internal/repositories/master"
)

var (
	ErrInvalidEkspedisi = errors.New("data ekspedisi tidak valid")
	ErrEkspedisiExists = errors.New("nama ekspedisi sudah digunakan")
)

type EkspedisiService struct {
	Repo *repository.EkspedisiRepository
}

func NewEkspedisiService(
	repo *repository.EkspedisiRepository,
) *EkspedisiService {
	return &EkspedisiService{
		Repo: repo,
	}
}

func (s *EkspedisiService) Create(data *models.MasterEkspedisi) error {
	data.NamaEkspedisi = strings.TrimSpace(data.NamaEkspedisi)

	if data.NamaEkspedisi == "" {
		return ErrInvalidEkspedisi
	}

	existing, err := s.Repo.FindByNama(data.NamaEkspedisi)
	if err != nil {
		return err
	}

	if existing != nil {
		return ErrEkspedisiExists
	}

	return s.Repo.Create(data)
}

func (s *EkspedisiService) FindAll() ([]models.MasterEkspedisi, error) {
	return s.Repo.FindAll()
}

func (s *EkspedisiService) FindActive() ([]models.MasterEkspedisi, error) {
	return s.Repo.FindActive()
}

func (s *EkspedisiService) FindByID(id uint) (*models.MasterEkspedisi, error) {
	if id == 0 {
		return nil, ErrInvalidEkspedisi
	}

	return s.Repo.FindByID(id)
}

func (s *EkspedisiService) FindByNama(nama string) (*models.MasterEkspedisi, error) {
	nama = strings.TrimSpace(nama)

	if nama == "" {
		return nil, ErrInvalidEkspedisi
	}

	return s.Repo.FindByNama(nama)
}

func (s *EkspedisiService) Update(data *models.MasterEkspedisi) error {
	if data.ID == 0 {
		return ErrInvalidEkspedisi
	}

	data.NamaEkspedisi = strings.TrimSpace(data.NamaEkspedisi)

	if data.NamaEkspedisi == "" {
		return ErrInvalidEkspedisi
	}

	existing, err := s.Repo.FindByNama(data.NamaEkspedisi)
	if err != nil {
		return err
	}

	if existing != nil && existing.ID != data.ID {
		return ErrEkspedisiExists
	}

	return s.Repo.Update(data)
}

func (s *EkspedisiService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidEkspedisi
	}

	return s.Repo.Delete(id)
}