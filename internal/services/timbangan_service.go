package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

type TimbanganService struct {
	Repository *repositories.TimbanganRepository
}

func NewTimbanganService(
	repository *repositories.TimbanganRepository,
) *TimbanganService {
	return &TimbanganService{
		Repository: repository,
	}
}

func (s *TimbanganService) FindAll() ([]models.Order, error) {
	return s.Repository.FindAll()
}

func (s *TimbanganService) FindByID(id uint) (*models.Order, error) {
	return s.Repository.FindByID(id)
}

func (s *TimbanganService) FindByKode(kode string) (*models.Order, error) {
	return s.Repository.FindByKode(kode)
}

func (s *TimbanganService) CreateFull(
	req *models.TimbanganCreateRequest,
) (*models.Order, error) {

	req.KodeOrder = strings.TrimSpace(req.KodeOrder)
	if req.KodeOrder == "" {
		return nil, errors.New("kode_order wajib diisi")
	}

	if !rules.IsKategoriTimbanganRCW(req.KategoriOrder) {
		return nil, errors.New("kategori_order harus salah satu dari: timbangan_inaproc, timbangan_manual, rcw_360, rcw_800w (kategori IOT pakai endpoint /iot-inaproc atau /iot-manual)")
	}
	if req.JenisOrder != "inaproc" && req.JenisOrder != "manual" {
		return nil, errors.New("jenis_order harus 'inaproc' atau 'manual'")
	}

	if req.InstansiID == 0 {
		return nil, errors.New("instansi_id wajib diisi")
	}
	if req.PicID == 0 {
		return nil, errors.New("pic_id wajib diisi")
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

	if len(req.Items) == 0 {
		return nil, errors.New("minimal 1 item produk wajib diisi")
	}
	for _, item := range req.Items {
		if item.ProdukID == 0 {
			return nil, errors.New("produk_id di items wajib diisi")
		}
		if item.Qty <= 0 {
			return nil, errors.New("qty di items harus lebih dari 0")
		}
	}

	if req.Procurement != nil {
		if req.Procurement.NoPOKUT != nil && !rules.IsProcurementFieldAllowed(req.KategoriOrder, "no_po_kut") {
			return nil, errors.New("no_po_kut tidak berlaku untuk kategori ini")
		}
		if req.Procurement.NomorSuratPenyampaianDaftarHarga != nil && !rules.IsProcurementFieldAllowed(req.KategoriOrder, "nomor_surat_penyampaian_daftar_harga") {
			return nil, errors.New("nomor_surat_penyampaian_daftar_harga tidak berlaku untuk kategori ini")
		}
		if req.Procurement.NomorFormulirPembelian != nil && !rules.IsProcurementFieldAllowed(req.KategoriOrder, "nomor_formulir_pembelian") {
			return nil, errors.New("nomor_formulir_pembelian tidak berlaku untuk kategori ini")
		}
	}

	if req.Pricing == nil {
		return nil, errors.New("data pricing (order_pricing) wajib diisi untuk kategori timbangan/RCW")
	}
	if req.Pricing.TipeTimbangan != nil && *req.Pricing.TipeTimbangan != "" && !rules.IsValidTipeTimbangan(*req.Pricing.TipeTimbangan) {
		return nil, errors.New("tipe timbangan tidak ada di daftar yang diperbolehkan")
	}
	if req.Pricing.HargaProduk < 0 || req.Pricing.TotalHargaJual < 0 {
		return nil, errors.New("harga tidak boleh negatif")
	}

	if req.Shipment != nil {
		if req.Shipment.Berat < 0 || req.Shipment.Ongkir < 0 {
			return nil, errors.New("berat/ongkir tidak boleh negatif")
		}
	}

	existing, err := s.Repository.FindByKode(req.KodeOrder)
	if err == nil && existing != nil {
		return nil, errors.New("kode_order sudah digunakan")
	}

	order := &models.Order{
		KodeOrder:        req.KodeOrder,
		JenisOrder:       req.JenisOrder,
		KategoriOrder:    req.KategoriOrder,
		InstansiID:       &req.InstansiID, 
		PicID:            &req.PicID,     
		Status:           req.Status,
		StatusOdoo:       req.StatusOdoo,
		NSFP:             req.NSFP,
		NoBAST:           req.NoBAST,
		KodeBayar:        req.KodeBayar,
		NoInvoiceInaproc: req.NoInvoiceInaproc,
		Keterangan:       req.Keterangan,
		TanggalPO:        req.TanggalPO,
		NoPO:             req.NoPO,
		TanggalBAST:      req.TanggalBAST,
	}

	if err := s.Repository.CreateFull(req, order); err != nil {
		return nil, err
	}

	return s.Repository.FindByID(order.ID)
}

func (s *TimbanganService) Update(id uint, data map[string]interface{}) error {
	if id == 0 {
		return errors.New("ID tidak valid")
	}

	order, err := s.Repository.FindByID(id)
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	if kodeOrder, touched := data["kode_order"]; touched {
		if kodeOrder != order.KodeOrder {
			return errors.New("kode_order tidak dapat diubah")
		}
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

func (s *TimbanganService) Delete(id uint) error {
	order, err := s.Repository.FindByID(id)
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	return s.Repository.Delete(order)
}