package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var (
	ErrPricingNotFound = errors.New("pricing order tidak ditemukan")
	ErrInvalidPricing  = errors.New("data pricing tidak valid")
)

type OrderPricingService struct {
	Repo *repositories.OrderPricingRepository
}

func NewOrderPricingService(
	repo *repositories.OrderPricingRepository,
) *OrderPricingService {
	return &OrderPricingService{
		Repo: repo,
	}
}

func (s *OrderPricingService) Create(data *models.OrderPricing) error {
	if data.OrderID == 0 {
		return ErrInvalidPricing
	}

	return s.Repo.Create(data)
}

func (s *OrderPricingService) FindByID(id uint) (*models.OrderPricing, error) {
	if id == 0 {
		return nil, ErrPricingNotFound
	}

	return s.Repo.FindByID(id)
}

func (s *OrderPricingService) FindByOrderID(orderID uint) (*models.OrderPricing, error) {
	if orderID == 0 {
		return nil, ErrInvalidPricing
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *OrderPricingService) Update(id uint, data map[string]interface{}) error {
	if id == 0 {
		return ErrInvalidPricing
	}

	existing, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if existing == nil {
		return ErrPricingNotFound
	}

	if len(data) == 0 {
		return ErrInvalidPricing
	}

	return s.Repo.Update(id, data)
}

func (s *OrderPricingService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidPricing
	}

	return s.Repo.DeleteByOrderID(orderID)
}