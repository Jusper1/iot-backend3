package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)


var ErrInvalidOrderDocument = errors.New("data dokumen order tidak valid")

type OrderDocumentService struct {
	Repo *repositories.OrderDocumentRepository
}

func NewOrderDocumentService(
	repo *repositories.OrderDocumentRepository,
) *OrderDocumentService {
	return &OrderDocumentService{
		Repo: repo,
	}
}

func (s *OrderDocumentService) Create(data *models.OrderDocument) error {
	if data.OrderID == 0 {
		return ErrInvalidOrderDocument
	}

	data.JenisDokumen = strings.TrimSpace(data.JenisDokumen)

	if data.JenisDokumen == "" {
		return ErrInvalidOrderDocument
	}

	return s.Repo.Create(data)
}

func (s *OrderDocumentService) FindByID(id uint) (*models.OrderDocument, error) {
	if id == 0 {
		return nil, ErrInvalidOrderDocument
	}

	return s.Repo.FindByID(id)
}

func (s *OrderDocumentService) FindByOrderID(orderID uint) ([]models.OrderDocument, error) {
	if orderID == 0 {
		return nil, ErrInvalidOrderDocument
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *OrderDocumentService) Update(data *models.OrderDocument) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidOrderDocument
	}

	data.JenisDokumen = strings.TrimSpace(data.JenisDokumen)

	if data.JenisDokumen == "" {
		return ErrInvalidOrderDocument
	}

	return s.Repo.Update(data)
}

func (s *OrderDocumentService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidOrderDocument
	}

	return s.Repo.Delete(id)
}

func (s *OrderDocumentService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidOrderDocument
	}

	return s.Repo.DeleteByOrderID(orderID)
}