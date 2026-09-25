package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

var ErrInvalidManualOrder = errors.New("data order manual tidak valid")

type OrderManualService struct {
	Repo      *repositories.OrderManualRepository
	OrderRepo *repositories.OrderRepository 
}

func NewOrderManualService(
	repo *repositories.OrderManualRepository,
	orderRepo *repositories.OrderRepository, 
) *OrderManualService {
	return &OrderManualService{
		Repo:      repo,
		OrderRepo: orderRepo,
	}
}

func (s *OrderManualService) validateKategori(orderID uint) error {
	order, err := s.OrderRepo.FindByID(orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrInvalidManualOrder
	}
	return rules.ValidateChildEntity(order.KategoriOrder, rules.EntityOrderManual)
}

func (s *OrderManualService) Create(data *models.OrderManual) error {
	if data.OrderID == 0 {
		return ErrInvalidManualOrder
	}

	if err := s.validateKategori(data.OrderID); err != nil {
		return err
	}

	return s.Repo.Create(data)
}

func (s *OrderManualService) FindAll() ([]models.OrderManual, error) {
	return s.Repo.FindAll()
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

	if err := s.validateKategori(data.OrderID); err != nil {
		return err
	}

	return s.Repo.Update(data)
}

func (s *OrderManualService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidManualOrder
	}

	return s.Repo.DeleteByOrderID(orderID)
}