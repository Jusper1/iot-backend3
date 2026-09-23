package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	repository "iot-backend/internal/repositories/master"
)

var (
	ErrInvalidProduk = errors.New("data produk tidak valid")
	ErrKodeProdukExists = errors.New("kode produk sudah digunakan")
)

type ProdukService struct {
	Repo *repository.ProdukRepository
}

func NewProdukService(
	repo *repository.ProdukRepository,
) *ProdukService {
	return &ProdukService{
		Repo: repo,
	}
}

func (s *ProdukService) Create(data *models.MasterProduk) error {
	data.KodeProduk = strings.TrimSpace(data.KodeProduk)
	data.NamaProduk = strings.TrimSpace(data.NamaProduk)

	if data.KodeProduk == "" || data.NamaProduk == "" {
		return ErrInvalidProduk
	}

	if data.Harga < 0 {
		return ErrInvalidProduk
	}

	existing, err := s.Repo.FindByKode(data.KodeProduk)
	if err != nil {
		return err
	}

	if existing != nil {
		return ErrKodeProdukExists
	}

	return s.Repo.Create(data)
}

func (s *ProdukService) FindAll() ([]models.MasterProduk, error) {
	return s.Repo.FindAll()
}

func (s *ProdukService) FindActive() ([]models.MasterProduk, error) {
	return s.Repo.FindActive()
}

func (s *ProdukService) FindByID(id uint) (*models.MasterProduk, error) {
	if id == 0 {
		return nil, ErrInvalidProduk
	}

	return s.Repo.FindByID(id)
}

func (s *ProdukService) FindByKode(kode string) (*models.MasterProduk, error) {
	kode = strings.TrimSpace(kode)

	if kode == "" {
		return nil, ErrInvalidProduk
	}

	return s.Repo.FindByKode(kode)
}

func (s *ProdukService) Update(data *models.MasterProduk) error {
	if data.ID == 0 {
		return ErrInvalidProduk
	}

	data.KodeProduk = strings.TrimSpace(data.KodeProduk)
	data.NamaProduk = strings.TrimSpace(data.NamaProduk)

	if data.KodeProduk == "" || data.NamaProduk == "" {
		return ErrInvalidProduk
	}

	if data.Harga < 0 {
		return ErrInvalidProduk
	}

	existing, err := s.Repo.FindByKode(data.KodeProduk)
	if err != nil {
		return err
	}

	if existing != nil && existing.ID != data.ID {
		return ErrKodeProdukExists
	}

	return s.Repo.Update(data)
}

func (s *ProdukService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidProduk
	}

	return s.Repo.Delete(id)
}