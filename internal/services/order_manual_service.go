package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var ErrInvalidManualOrder = errors.New("data order manual tidak valid")

type OrderManualService struct {
	Repo *repositories.OrderManualRepository
}

func NewOrderManualService(
	repo *repositories.OrderManualRepository,
) *OrderManualService {
	return &OrderManualService{
		Repo: repo,
	}
}

func (s *OrderManualService) Create(data *models.OrderManual) error {
	if data.OrderID == 0 {
		return ErrInvalidManualOrder
	}

	return s.Repo.Create(data)
}

func (s *OrderManualService) FindByID(id uint) (*models.OrderManual, error) {
	if id == 0 {
		return nil, ErrInvalidManualOrder
	}

	return s.Repo.FindByID(id)
}

func (s *OrderManualService) FindByOrderID(orderID uint) (*models.OrderManual, error) {
	if orderID == 0 {
		return nil, ErrInvalidManualOrder
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *OrderManualService) Update(data *models.OrderManual) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidManualOrder
	}

	return s.Repo.Update(data)
}

func (s *OrderManualService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidManualOrder
	}

	return s.Repo.DeleteByOrderID(orderID)
}