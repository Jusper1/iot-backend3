package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var ErrInvalidDocument = errors.New("data document tidak valid")

type DocumentService struct {
	Repo *repositories.DocumentRepository
}

func NewDocumentService(
	repo *repositories.DocumentRepository,
) *DocumentService {
	return &DocumentService{
		Repo: repo,
	}
}

func (s *DocumentService) Create(data *models.Document) error {
	data.JenisDokumen = strings.TrimSpace(data.JenisDokumen)
	data.NamaFile = strings.TrimSpace(data.NamaFile)

	if data.JenisDokumen == "" || data.NamaFile == "" {
		return ErrInvalidDocument
	}

	if data.OrderID == nil && data.SPJID == nil {
		return ErrInvalidDocument
	}

	return s.Repo.Create(data)
}

func (s *DocumentService) FindByID(id uint) (*models.Document, error) {
	if id == 0 {
		return nil, ErrInvalidDocument
	}

	return s.Repo.FindByID(id)
}

func (s *DocumentService) FindByOrderID(orderID uint) ([]models.Document, error) {
	if orderID == 0 {
		return nil, ErrInvalidDocument
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *DocumentService) FindBySPJID(spjID uint) ([]models.Document, error) {
	if spjID == 0 {
		return nil, ErrInvalidDocument
	}

	return s.Repo.FindBySPJID(spjID)
}

func (s *DocumentService) FindAll() ([]models.Document, error) {
	return s.Repo.FindAll()
}

func (s *DocumentService) Update(data *models.Document) error {
	if data.ID == 0 {
		return ErrInvalidDocument
	}

	data.JenisDokumen = strings.TrimSpace(data.JenisDokumen)
	data.NamaFile = strings.TrimSpace(data.NamaFile)

	if data.JenisDokumen == "" || data.NamaFile == "" {
		return ErrInvalidDocument
	}

	return s.Repo.Update(data)
}

func (s *DocumentService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidDocument
	}

	return s.Repo.Delete(id)
}