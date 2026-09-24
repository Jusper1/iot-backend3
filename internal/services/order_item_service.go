package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var (
	ErrOrderItemNotFound = errors.New("order item tidak ditemukan")
	ErrInvalidOrderItem  = errors.New("data order item tidak valid")
)

type OrderItemService struct {
	Repo *repositories.OrderItemRepository
}

func NewOrderItemService(
	repo *repositories.OrderItemRepository,
) *OrderItemService {
	return &OrderItemService{
		Repo: repo,
	}
}

func (s *OrderItemService) Create(item *models.OrderItem) error {
	if item.OrderID == 0 {
		return ErrInvalidOrderItem
	}

	if item.ProdukID == 0 {
		return ErrInvalidOrderItem
	}

	if item.Qty == 0 {
		return ErrInvalidOrderItem
	}

	return s.Repo.Create(item)
}

func (s *OrderItemService) CreateMany(items []models.OrderItem) error {
	if len(items) == 0 {
		return ErrInvalidOrderItem
	}

	for _, item := range items {
		if item.OrderID == 0 || item.ProdukID == 0 || item.Qty == 0 {
			return ErrInvalidOrderItem
		}
	}

	return s.Repo.CreateMany(items)
}

func (s *OrderItemService) FindAll() ([]models.OrderItem, error) {
	return s.Repo.FindAll()
}

func (s *OrderItemService) FindByID(id uint) (*models.OrderItem, error) {
	if id == 0 {
		return nil, ErrOrderItemNotFound
	}

	data, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *OrderItemService) FindByOrderID(orderID uint) ([]models.OrderItem, error) {
	if orderID == 0 {
		return nil, ErrInvalidOrderItem
	}

	return s.Repo.FindByOrderID(orderID)
}

func (s *OrderItemService) Update(item *models.OrderItem) error {
	if item.ID == 0 || item.OrderID == 0 || item.ProdukID == 0 {
		return ErrInvalidOrderItem
	}

	if item.Qty == 0 {
		return ErrInvalidOrderItem
	}

	return s.Repo.Update(item)
}

func (s *OrderItemService) Delete(id uint) error {
	if id == 0 {
		return ErrOrderItemNotFound
	}

	return s.Repo.Delete(id)
}

func (s *OrderItemService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidOrderItem
	}

	return s.Repo.DeleteByOrderID(orderID)
}