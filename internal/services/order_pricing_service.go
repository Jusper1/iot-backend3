package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

var (
	ErrPricingNotFound         = errors.New("pricing order tidak ditemukan")
	ErrInvalidPricing          = errors.New("data pricing tidak valid")
	ErrTipeTimbanganTidakValid = errors.New("tipe timbangan tidak ada di daftar yang diperbolehkan")
)

type OrderPricingService struct {
	Repo      *repositories.OrderPricingRepository
	OrderRepo *repositories.OrderRepository 
}

func NewOrderPricingService(
	repo *repositories.OrderPricingRepository,
	orderRepo *repositories.OrderRepository, 
) *OrderPricingService {
	return &OrderPricingService{
		Repo:      repo,
		OrderRepo: orderRepo,
	}
}


func (s *OrderPricingService) validateKategori(orderID uint) error {
	order, err := s.OrderRepo.FindByID(orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrInvalidPricing
	}
	return rules.ValidateChildEntity(order.KategoriOrder, rules.EntityOrderPricing)
}

func (s *OrderPricingService) Create(data *models.OrderPricing) error {
	if data.OrderID == 0 {
		return ErrInvalidPricing
	}

	if err := s.validateKategori(data.OrderID); err != nil {
		return err
	}

	if data.TipeTimbangan != nil && *data.TipeTimbangan != "" && !rules.IsValidTipeTimbangan(*data.TipeTimbangan) {
		return ErrTipeTimbanganTidakValid
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

	if err := s.validateKategori(existing.OrderID); err != nil {
		return err
	}

	if raw, touched := data["tipe_timbangan"]; touched {
		if value, ok := raw.(string); ok && value != "" && !rules.IsValidTipeTimbangan(value) {
			return ErrTipeTimbanganTidakValid
		}
	}

	return s.Repo.Update(id, data)
}

func (s *OrderPricingService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidPricing
	}

	return s.Repo.DeleteByOrderID(orderID)
}