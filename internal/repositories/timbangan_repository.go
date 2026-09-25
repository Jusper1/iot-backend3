package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type TimbanganRepository struct {
	DB *gorm.DB
}

func NewTimbanganRepository(db *gorm.DB) *TimbanganRepository {
	return &TimbanganRepository{
		DB: db,
	}
}

var timbanganKategoriList = []string{
	"timbangan_inaproc",
	"timbangan_manual",
	"rcw_360",
	"rcw_800w",
}

func (r *TimbanganRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Pricing").
		Preload("Shipments").
		Preload("Shipments.Wilayah").
		Preload("Shipments.Ekspedisi").
		Preload("Payments").
		Where("kategori_order IN ?", timbanganKategoriList).
		Order("id DESC").
		Find(&orders).Error

	return orders, err
}

func (r *TimbanganRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Pricing").
		Preload("Shipments").
		Preload("Shipments.Wilayah").
		Preload("Shipments.Ekspedisi").
		Preload("Payments").
		Where("kategori_order IN ?", timbanganKategoriList).
		First(&order, id).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *TimbanganRepository) FindByKode(kode string) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Pricing").
		Preload("Shipments").
		Preload("Shipments.Wilayah").
		Preload("Shipments.Ekspedisi").
		Preload("Payments").
		Where("kode_order = ?", kode).
		Where("kategori_order IN ?", timbanganKategoriList).
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *TimbanganRepository) Update(
	id uint,
	data map[string]interface{},
) error {
	updates := make(map[string]interface{})

	allowedFields := map[string]bool{
		"instansi_id":        true,
		"pic_id":             true,
		"status":             true,
		"status_odoo":        true,
		"nsfp":               true,
		"no_bast":            true,
		"kode_bayar":         true,
		"no_invoice_inaproc": true,
		"keterangan":         true,
		"tanggal_po":         true,
		"no_po":              true,
		"tanggal_bast":       true,

	}

	for field, value := range data {
		if allowedFields[field] {
			updates[field] = value
		}
	}

	if len(updates) == 0 {
		return gorm.ErrInvalidData
	}

	return r.DB.
		Model(&models.Order{}).
		Where("id = ?", id).
		Where("kategori_order IN ?", timbanganKategoriList).
		Updates(updates).Error
}

func (r *TimbanganRepository) Delete(order *models.Order) error {
	return r.DB.Delete(order).Error
}

func (r *TimbanganRepository) CreateFull(
	req *models.TimbanganCreateRequest,
	order *models.Order,
) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for _, itemReq := range req.Items {
			item := models.OrderItem{
				OrderID:  order.ID,
				ProdukID: itemReq.ProdukID,
				Qty:      itemReq.Qty,
				Harga:    itemReq.Harga,
				PPN:      itemReq.PPN,
				Subtotal: itemReq.Subtotal,
			}

			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		if req.Procurement != nil {
			procurement := models.OrderProcurement{
				OrderID:                          order.ID,
				NoPOKUT:                          req.Procurement.NoPOKUT,
				TanggalInvoiceKUT:                req.Procurement.TanggalInvoiceKUT,
				NomorSuratPenyampaianDaftarHarga: req.Procurement.NomorSuratPenyampaianDaftarHarga,
				NomorFormulirPembelian:           req.Procurement.NomorFormulirPembelian,
				NoInvoiceKUT:                     req.Procurement.NoInvoiceKUT,
			}

			if err := tx.Create(&procurement).Error; err != nil {
				return err
			}
		}

		pricing := models.OrderPricing{
			OrderID:                  order.ID,
			TipeTimbangan:            req.Pricing.TipeTimbangan,
			HargaProduk:              req.Pricing.HargaProduk,
			HargaPPN:                 req.Pricing.HargaPPN,
			HargaOngkirKUT:           req.Pricing.HargaOngkirKUT,
			HargaPPNOngkir:           req.Pricing.HargaPPNOngkir,
			TotalHargaOngkir:         req.Pricing.TotalHargaOngkir,
			TotalHargaJual:           req.Pricing.TotalHargaJual,
			HargaProdukReseller:      req.Pricing.HargaProdukReseller,
			HargaPPNReseller:         req.Pricing.HargaPPNReseller,
			HargaOngkirReseller:      req.Pricing.HargaOngkirReseller,
			HargaPPNOngkirReseller:   req.Pricing.HargaPPNOngkirReseller,
			TotalHargaOngkirReseller: req.Pricing.TotalHargaOngkirReseller,
			TotalHargaReseller:       req.Pricing.TotalHargaReseller,
		}

		if err := tx.Create(&pricing).Error; err != nil {
			return err
		}

		if req.Shipment != nil {
			shipment := models.Shipment{
				OrderID:          order.ID,
				WilayahID:        req.Shipment.WilayahID,
				EkspedisiID:      req.Shipment.EkspedisiID,
				Resi:             req.Shipment.Resi,
				Berat:            req.Shipment.Berat,
				Ongkir:           req.Shipment.Ongkir,
				TanggalKirim:     req.Shipment.TanggalKirim,
				TanggalDiterima:  req.Shipment.TanggalDiterima,
				StatusPengiriman: req.Shipment.StatusPengiriman,
			}

			if err := tx.Create(&shipment).Error; err != nil {
				return err
			}
		}

		if req.Payment != nil {
			payment := models.Payment{
				OrderID:          order.ID,
				JumlahUangMasuk:  req.Payment.JumlahUangMasuk,
				TanggalUangMasuk: req.Payment.TanggalUangMasuk,
				Status:           req.Payment.Status,
				Rekening:         req.Payment.Rekening,
				BuktiPembayaran:  req.Payment.BuktiPembayaran,
			}

			if err := tx.Create(&payment).Error; err != nil {
				return err
			}
		}

		return nil
	})
}