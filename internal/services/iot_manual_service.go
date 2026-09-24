package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

type IOTManualService struct {
	Repository *repositories.IOTManualRepository
}

func NewIOTManualService(
	repository *repositories.IOTManualRepository,
) *IOTManualService {
	return &IOTManualService{
		Repository: repository,
	}
}

func (s *IOTManualService) FindAll() ([]models.Order, error) {
	return s.Repository.FindAll()
}

func (s *IOTManualService) FindByID(id uint) (*models.Order, error) {
	return s.Repository.FindByID(id)
}

func (s *IOTManualService) FindByKode(kode string) (*models.Order, error) {
	return s.Repository.FindByKode(kode)
}

func (s *IOTManualService) CreateFull(
	req *models.IOTManualCreateRequest,
) (*models.Order, error) {

	if strings.TrimSpace(req.KodeOrder) == "" {
		return nil, errors.New("kode_order wajib diisi")
	}

	if req.InstansiID == 0 {
		return nil, errors.New("instansi_id wajib diisi")
	}

	if req.PicID == 0 {
		return nil, errors.New("pic_id wajib diisi")
	}

	if len(req.Items) == 0 {
		return nil, errors.New("minimal harus ada 1 produk")
	}

	for _, item := range req.Items {

		if item.ProdukID == 0 {
			return nil, errors.New("produk_id wajib diisi")
		}

		if item.Qty <= 0 {
			return nil, errors.New("qty harus lebih dari 0")
		}
	}

	existing, err := s.Repository.FindByKode(req.KodeOrder)

	if err == nil && existing != nil {
		return nil, errors.New("kode_order sudah digunakan")
	}

	instansiID := req.InstansiID
	picID := req.PicID

	order := &models.Order{
		KodeOrder:        req.KodeOrder,
		JenisOrder:       "manual",
		KategoriOrder:    "iot_manual",
		InstansiID:       &instansiID,
		PicID:            &picID,
		Status:           req.Status,
		StatusOdoo:       req.StatusOdoo,
		NSFP:             req.NSFP,
		NoBAST:           req.NoBAST,
		Keterangan:       req.Keterangan,
		TanggalPO:        req.TanggalPO,
		NoPO:             req.NoPO,
		PeriodeLangganan: req.PeriodeLangganan,
		TanggalBAST:      req.TanggalBAST,
	}

	if err := s.Repository.CreateFull(req, order); err != nil {
		return nil, err
	}

	result, err := s.Repository.FindByID(order.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *IOTManualService) Update(id uint, data map[string]interface{}) error {
	if id == 0 {
		return errors.New("id order tidak valid")
	}

	order, err := s.Repository.FindByID(id)
	if err != nil {
		return errors.New("data IOT MANUAL tidak ditemukan")
	}

	if order.JenisOrder != "manual" ||
		order.KategoriOrder != "iot_manual" {
		return errors.New("data bukan IOT MANUAL")
	}

	value, exists := data["kode_order"]

	if !exists {
		return errors.New("kode_order wajib diisi")
	}

	kodeOrder, ok := value.(string)
	if !ok {
		return errors.New("kode_order harus berupa string")
	}

	kodeOrder = strings.TrimSpace(kodeOrder)

	if kodeOrder == "" {
		return errors.New("kode_order wajib diisi")
	}

	if kodeOrder != order.KodeOrder {
		return errors.New("kode_order tidak dapat diubah")
	}

	return s.Repository.Update(id, data)
}

func (s *IOTManualService) Delete(id uint) error {

	order, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("data IOT MANUAL tidak ditemukan")
	}

	return s.Repository.Delete(order)
}
