package services

import (
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

var (
	ErrInvalidProcurement         = errors.New("data procurement tidak valid")
	ErrProcurementNotFound        = errors.New("procurement tidak ditemukan")
	ErrProcurementFieldNotAllowed = errors.New("field ini hanya berlaku untuk kategori timbangan/RCW (no_po_kut, nomor_surat_penyampaian_daftar_harga, nomor_formulir_pembelian)")
)

type OrderProcurementService struct {
	Repo      *repositories.OrderProcurementRepository
	OrderRepo *repositories.OrderRepository 
}

func NewOrderProcurementService(
	repo *repositories.OrderProcurementRepository,
	orderRepo *repositories.OrderRepository, 
) *OrderProcurementService {
	return &OrderProcurementService{
		Repo:      repo,
		OrderRepo: orderRepo,
	}
}

func (s *OrderProcurementService) validateFields(orderID uint, data *models.OrderProcurement) error {
	order, err := s.OrderRepo.FindByID(orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrInvalidProcurement
	}

	if data.NoPOKUT != nil && !rules.IsProcurementFieldAllowed(order.KategoriOrder, "no_po_kut") {
		return ErrProcurementFieldNotAllowed
	}
	if data.NomorSuratPenyampaianDaftarHarga != nil && !rules.IsProcurementFieldAllowed(order.KategoriOrder, "nomor_surat_penyampaian_daftar_harga") {
		return ErrProcurementFieldNotAllowed
	}
	if data.NomorFormulirPembelian != nil && !rules.IsProcurementFieldAllowed(order.KategoriOrder, "nomor_formulir_pembelian") {
		return ErrProcurementFieldNotAllowed
	}

	return nil
}

func (s *OrderProcurementService) Create(data *models.OrderProcurement) error {
	if data.OrderID == 0 {
		return ErrInvalidProcurement
	}

	if err := s.validateFields(data.OrderID, data); err != nil {
		return err
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

	order, err := s.OrderRepo.FindByID(existing.OrderID)
	if err != nil {
		return err
	}
	if order == nil {
		return ErrInvalidProcurement
	}

	restrictedFields := []string{
		"no_po_kut",
		"nomor_surat_penyampaian_daftar_harga",
		"nomor_formulir_pembelian",
	}
	for _, field := range restrictedFields {
		if _, touched := data[field]; touched && !rules.IsProcurementFieldAllowed(order.KategoriOrder, field) {
			return ErrProcurementFieldNotAllowed
		}
	}

	return s.Repo.Update(id, data)
}

func (s *OrderProcurementService) DeleteByOrderID(orderID uint) error {
	if orderID == 0 {
		return ErrInvalidProcurement
	}

	return s.Repo.DeleteByOrderID(orderID)
}