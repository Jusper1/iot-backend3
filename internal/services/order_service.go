package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

var (
	ErrKodeOrderRequired = errors.New("kode order wajib diisi")
	ErrKodeOrderExists   = errors.New("kode order sudah digunakan")
	ErrOrderNotFound     = errors.New("order tidak ditemukan")
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderService(
	orderRepo *repositories.OrderRepository,
) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo,
	}
}

func (s *OrderService) Create(order *models.Order) error {
	order.KodeOrder = strings.TrimSpace(order.KodeOrder)

	if order.KodeOrder == "" {
		return ErrKodeOrderRequired
	}

	exists, err := s.OrderRepo.ExistsByKode(order.KodeOrder)
	if err != nil {
		return err
	}

	if exists {
		return ErrKodeOrderExists
	}

	return s.OrderRepo.Create(order)
}

func (s *OrderService) FindAll() ([]models.Order, error) {
	return s.OrderRepo.FindAll()
}

func (s *OrderService) FindByID(id uint) (*models.Order, error) {
	if id == 0 {
		return nil, ErrOrderNotFound
	}

	order, err := s.OrderRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

func (s *OrderService) FindByKode(kode string) (*models.Order, error) {
	kode = strings.TrimSpace(kode)

	if kode == "" {
		return nil, ErrKodeOrderRequired
	}

	order, err := s.OrderRepo.FindByKode(kode)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

func (s *OrderService) Update(order *models.Order) error {
	if order.ID == 0 {
		return ErrOrderNotFound
	}

	order.KodeOrder = strings.TrimSpace(order.KodeOrder)

	if order.KodeOrder == "" {
		return ErrKodeOrderRequired
	}

	existing, err := s.OrderRepo.FindByKode(order.KodeOrder)
	if err != nil {
		return err
	}

	if existing != nil && existing.ID != order.ID {
		return ErrKodeOrderExists
	}

	return s.OrderRepo.Update(order)
}

func (s *OrderService) Delete(id uint) error {
	if id == 0 {
		return ErrOrderNotFound
	}

	order, err := s.OrderRepo.FindByID(id)
	if err != nil {
		return err
	}

	if order == nil {
		return ErrOrderNotFound
	}

	return s.OrderRepo.Delete(id)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}