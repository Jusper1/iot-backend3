package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

type IOTInaprocService struct {
	Repository *repositories.IOTInaprocRepository
}

func NewIOTInaprocService(
	repository *repositories.IOTInaprocRepository,
) *IOTInaprocService {
	return &IOTInaprocService{
		Repository: repository,
	}
}

func (s *IOTInaprocService) FindAll() ([]models.Order, error) {
	return s.Repository.FindAll()
}

func (s *IOTInaprocService) FindByID(id uint) (*models.Order, error) {
	return s.Repository.FindByID(id)
}

func (s *IOTInaprocService) FindByKode(kode string) (*models.Order, error) {
	return s.Repository.FindByKode(kode)
}

func (s *IOTInaprocService) Create(order *models.Order) error {

	if strings.TrimSpace(order.KodeOrder) == "" {
		return errors.New("kode_order wajib diisi")
	}

	if order.InstansiID == nil {
		return errors.New("instansi_id wajib diisi")
	}

	if order.PicID == nil {
		return errors.New("pic_id wajib diisi")
	}

	order.JenisOrder = "inaproc"
	order.KategoriOrder = "iot_inaproc"

	existing, err := s.Repository.FindByKode(order.KodeOrder)

	if err == nil && existing != nil {
		return errors.New("kode_order sudah digunakan")
	}

	return s.Repository.Create(order)
}

func (s *IOTInaprocService) CreateFull(
	req *models.IOTInaprocCreateRequest,
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

	if req.Status != nil && *req.Status != "" && !rules.IsValidStatusPesanan(*req.Status) {
		return nil, errors.New("status pesanan tidak ada di daftar yang diperbolehkan")
	}
	if req.StatusOdoo != nil && *req.StatusOdoo != "" && !rules.IsValidStatusOdoo(*req.StatusOdoo) {
		return nil, errors.New("status odoo tidak ada di daftar yang diperbolehkan")
	}
	if req.KodeBayar != nil && *req.KodeBayar != "" && !rules.IsValidKodeBayar(*req.KodeBayar) {
		return nil, errors.New("kode bayar tidak ada di daftar yang diperbolehkan")
	}

	existing, err := s.Repository.FindByKode(req.KodeOrder)

	if err == nil && existing != nil {
		return nil, errors.New("kode_order sudah digunakan")
	}


	instansiID := req.InstansiID
	picID := req.PicID

	order := &models.Order{
		KodeOrder:        req.KodeOrder,
		JenisOrder:       "inaproc",
		KategoriOrder:    "iot_inaproc",
		InstansiID:       &instansiID,
		PicID:            &picID,
		Status:           req.Status,
		StatusOdoo:       req.StatusOdoo,
		NSFP:             req.NSFP,
		NoBAST:           req.NoBAST,
		KodeBayar:        req.KodeBayar,
		NoInvoiceInaproc: req.NoInvoiceInaproc,
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

func (s *IOTInaprocService) Update(id uint, data map[string]interface{}) error {
	if id == 0 {
		return errors.New("id order tidak valid")
	}

	order, err := s.Repository.FindByID(id)
	if err != nil {
		return errors.New("data IOT INAPROC tidak ditemukan")
	}

	if order.JenisOrder != "inaproc" ||
		order.KategoriOrder != "iot_inaproc" {
		return errors.New("data bukan IOT INAPROC")
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

	if raw, touched := data["status"]; touched {
		if v, ok := raw.(string); ok && v != "" && !rules.IsValidStatusPesanan(v) {
			return errors.New("status pesanan tidak ada di daftar yang diperbolehkan")
		}
	}
	if raw, touched := data["status_odoo"]; touched {
		if v, ok := raw.(string); ok && v != "" && !rules.IsValidStatusOdoo(v) {
			return errors.New("status odoo tidak ada di daftar yang diperbolehkan")
		}
	}
	if raw, touched := data["kode_bayar"]; touched {
		if v, ok := raw.(string); ok && v != "" && !rules.IsValidKodeBayar(v) {
			return errors.New("kode bayar tidak ada di daftar yang diperbolehkan")
		}
	}

	return s.Repository.Update(id, data)
}

func (s *IOTInaprocService) Delete(id uint) error {

	order, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("data IOT INAPROC tidak ditemukan")
	}

	return s.Repository.Delete(order)
}