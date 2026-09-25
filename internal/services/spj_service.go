package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

var (
	ErrInvalidSPJ           = errors.New("data SPJ tidak valid")
	ErrSPJNotFound          = errors.New("SPJ tidak ditemukan")
	ErrJenisKertasTidakValid = errors.New("jenis kertas harus salah satu dari: A4, F4")
	ErrJenisFileTidakValid   = errors.New("jenis file tidak ada di daftar yang diperbolehkan")
)

type SPJService struct {
	Repo *repositories.SPJRepository
}

func NewSPJService(
	repo *repositories.SPJRepository,
) *SPJService {
	return &SPJService{
		Repo: repo,
	}
}

func (s *SPJService) validateDropdowns(data *models.SPJ) error {
	if data.JenisKertas != nil && *data.JenisKertas != "" && !rules.IsValidJenisKertas(*data.JenisKertas) {
		return ErrJenisKertasTidakValid
	}
	if data.JenisFile != nil && *data.JenisFile != "" && !rules.IsValidJenisFile(*data.JenisFile) {
		return ErrJenisFileTidakValid
	}
	return nil
}

func (s *SPJService) Create(data *models.SPJ) error {
	if data.OrderID == 0 {
		return ErrInvalidSPJ
	}

	if data.JumlahRangkap == 0 {
		return ErrInvalidSPJ
	}

	if err := s.validateDropdowns(data); err != nil {
		return err
	}

	if data.KebutuhanSPJ != nil {
		value := strings.TrimSpace(*data.KebutuhanSPJ)
		data.KebutuhanSPJ = &value
	}

	if data.JenisKertas != nil {
		value := strings.TrimSpace(*data.JenisKertas)
		data.JenisKertas = &value
	}

	if data.JenisFile != nil {
		value := strings.TrimSpace(*data.JenisFile)
		data.JenisFile = &value
	}

	if data.PICPrint != nil {
		value := strings.TrimSpace(*data.PICPrint)
		data.PICPrint = &value
	}

	if data.Status != nil {
		value := strings.TrimSpace(*data.Status)
		data.Status = &value
	}

	return s.Repo.Create(data)
}

func (s *SPJService) FindByID(id uint) (*models.SPJ, error) {
	if id == 0 {
		return nil, ErrSPJNotFound
	}

	data, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, ErrSPJNotFound
	}

	return data, nil
}

func (s *SPJService) FindByOrderID(orderID uint) (*models.SPJ, error) {
	if orderID == 0 {
		return nil, ErrInvalidSPJ
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *SPJService) FindAll() ([]models.SPJ, error) {
	return s.Repo.FindAll()
}

func (s *SPJService) Update(data *models.SPJ) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidSPJ
	}

	if data.JumlahRangkap == 0 {
		return ErrInvalidSPJ
	}

	if err := s.validateDropdowns(data); err != nil {
		return err
	}

	if data.KebutuhanSPJ != nil {
		value := strings.TrimSpace(*data.KebutuhanSPJ)
		data.KebutuhanSPJ = &value
	}

	if data.JenisKertas != nil {
		value := strings.TrimSpace(*data.JenisKertas)
		data.JenisKertas = &value
	}

	if data.JenisFile != nil {
		value := strings.TrimSpace(*data.JenisFile)
		data.JenisFile = &value
	}

	if data.PICPrint != nil {
		value := strings.TrimSpace(*data.PICPrint)
		data.PICPrint = &value
	}

	if data.Status != nil {
		value := strings.TrimSpace(*data.Status)
		data.Status = &value
	}

	return s.Repo.Update(data)
}

func (s *SPJService) Delete(id uint) error {
	if id == 0 {
		return ErrSPJNotFound
	}

	return s.Repo.Delete(id)
}