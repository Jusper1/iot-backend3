package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)


var ErrInvalidProcurement = errors.New("data procurement tidak valid")

type OrderProcurementService struct {
	Repo *repositories.OrderProcurementRepository
}

func NewOrderProcurementService(
	repo *repositories.OrderProcurementRepository,
) *OrderProcurementService {
	return &OrderProcurementService{
		Repo: repo,
	}
}

func (s *OrderProcurementService) Create(data *models.OrderProcurement) error {
	if data.OrderID == 0 {
		return ErrInvalidProcurement
	}

	return s.Repo.Create(data)
}

func (s *OrderProcurementService) FindByID(id uint) (*models.OrderProcurement, error) {
	if id == 0 {
		return nil, ErrInvalidProcurement
	}

	return s.Repo.FindByID(id)
}

func (s *OrderProcurementService) FindByOrderID(orderID uint) (*models.OrderProcurement, error) {
	if orderID == 0 {
		return nil, ErrInvalidProcurement
	}

	return s.Repo.FindByOrderID(orderID)
}

var (
	ErrProcurementNotFound = errors.New("procurement tidak ditemukan")
)
func (s *OrderProcurementService) Update(
	id uint,
	data map[string]interface{},
) error {
	if id == 0 {
		return ErrInvalidProcurement
	}

	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if existing == nil {
		return ErrProcurementNotFound
	}

	if len(data) == 0 {
		return ErrInvalidProcurement
	}

	return s.Repo.Update(id, data)
}

func (s *OrderProcurementService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidProcurement
	}

	return s.Repo.DeleteByOrderID(orderID)
}