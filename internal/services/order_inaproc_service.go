package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var ErrInvalidInaproc = errors.New("data order inaproc tidak valid (CATATAN: entitas ini deprecated, gunakan orders.no_invoice_inaproc)")

type OrderInaprocService struct {
	Repo *repositories.OrderInaprocRepository
}

func NewOrderInaprocService(
	repo *repositories.OrderInaprocRepository,
) *OrderInaprocService {
	return &OrderInaprocService{
		Repo: repo,
	}
}

func (s *OrderInaprocService) Create(data *models.OrderInaproc) error {
	if data.OrderID == 0 {
		return ErrInvalidInaproc
	}

	return s.Repo.Create(data)
}

func (s *OrderInaprocService) FindByID(id uint) (*models.OrderInaproc, error) {
	if id == 0 {
		return nil, ErrInvalidInaproc
	}

	return s.Repo.FindByID(id)
}

func (s *OrderInaprocService) FindByOrderID(orderID uint) (*models.OrderInaproc, error) {
	if orderID == 0 {
		return nil, ErrInvalidInaproc
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *OrderInaprocService) Update(data *models.OrderInaproc) error {
	if data.ID == 0 || data.OrderID == 0 {
		return ErrInvalidInaproc
	}

	return s.Repo.Update(data)
}

func (s *OrderInaprocService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidInaproc
	}

	return s.Repo.DeleteByOrderID(orderID)
}