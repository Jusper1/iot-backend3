package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

var (
	ErrKodeOrderRequired = errors.New("kode order wajib diisi")
	ErrKodeOrderExists   = errors.New("kode order sudah digunakan")
	ErrOrderNotFound     = errors.New("order tidak ditemukan")
	ErrKategoriTidakValid   = errors.New("kategori_order tidak dikenal, harus salah satu dari: iot_inaproc, iot_manual, timbangan_inaproc, timbangan_manual, rcw_360, rcw_800w")
	ErrStatusTidakValid     = errors.New("status pesanan tidak ada di daftar yang diperbolehkan")
	ErrStatusOdooTidakValid = errors.New("status odoo tidak ada di daftar yang diperbolehkan")
	ErrKodeBayarTidakValid  = errors.New("kode bayar tidak ada di daftar yang diperbolehkan")
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

	if !rules.IsKategoriValid(order.KategoriOrder) {
		return ErrKategoriTidakValid
	}

	if order.Status != nil && *order.Status != "" && !rules.IsValidStatusPesanan(*order.Status) {
		return ErrStatusTidakValid
	}
	if order.StatusOdoo != nil && *order.StatusOdoo != "" && !rules.IsValidStatusOdoo(*order.StatusOdoo) {
		return ErrStatusOdooTidakValid
	}
	if order.KodeBayar != nil && *order.KodeBayar != "" && !rules.IsValidKodeBayar(*order.KodeBayar) {
		return ErrKodeBayarTidakValid
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

	if !rules.IsKategoriValid(order.KategoriOrder) {
		return ErrKategoriTidakValid
	}
	if order.Status != nil && *order.Status != "" && !rules.IsValidStatusPesanan(*order.Status) {
		return ErrStatusTidakValid
	}
	if order.StatusOdoo != nil && *order.StatusOdoo != "" && !rules.IsValidStatusOdoo(*order.StatusOdoo) {
		return ErrStatusOdooTidakValid
	}
	if order.KodeBayar != nil && *order.KodeBayar != "" && !rules.IsValidKodeBayar(*order.KodeBayar) {
		return ErrKodeBayarTidakValid
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