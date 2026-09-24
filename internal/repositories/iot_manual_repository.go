package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type IOTManualRepository struct {
	DB *gorm.DB
}

func NewIOTManualRepository(db *gorm.DB) *IOTManualRepository {
	return &IOTManualRepository{
		DB: db,
	}
}

func (r *IOTManualRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Manual").
		Preload("Procurement").
		Preload("Payments").
		Where("jenis_order = ?", "manual").
		Where("kategori_order = ?", "iot_manual").
		Order("id DESC").
		Find(&orders).Error

	return orders, err
}

func (r *IOTManualRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Manual").
		Preload("Procurement").
		Preload("Payments").
		Where("jenis_order = ?", "manual").
		Where("kategori_order = ?", "iot_manual").
		First(&order, id).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *IOTManualRepository) FindByKode(kode string) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Manual").
		Preload("Procurement").
		Preload("Payments").
		Where("kode_order = ?", kode).
		Where("jenis_order = ?", "manual").
		Where("kategori_order = ?", "iot_manual").
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *IOTManualRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *IOTManualRepository) Update(
	id uint,
	data map[string]interface{},
) error {

	orderUpdates := make(map[string]interface{})
	manualUpdates := make(map[string]interface{})

	allowedOrderFields := map[string]bool{
		"instansi_id":       true,
		"pic_id":            true,
		"status":            true,
		"status_odoo":       true,
		"nsfp":              true,
		"no_bast":           true,
		"keterangan":        true,
		"tanggal_po":        true,
		"no_po":             true,
		"periode_langganan": true,
		"tanggal_bast":      true,
	}

	allowedManualFields := map[string]bool{
		"nomor_formulir_berlangganan": true,
		"nomor_surat_penawaran_harga": true,
		"bulan_pengiriman_sph":        true,
		"waktu_pengiriman":            true,
		"dokumen_full_sign":           true,
		"nama_surat":                  true,
		"nomor_kontrak_berlangganan":  true,
	}

	for field, value := range data {
		if allowedOrderFields[field] {
			orderUpdates[field] = value
		}
		if allowedManualFields[field] {
			manualUpdates[field] = value
		}
	}

	if len(orderUpdates) == 0 && len(manualUpdates) == 0 {
		return gorm.ErrInvalidData
	}

	return r.DB.Transaction(func(tx *gorm.DB) error {

		if len(orderUpdates) > 0 {
			err := tx.
				Model(&models.Order{}).
				Where("id = ?", id).
				Where("jenis_order = ?", "manual").
				Where("kategori_order = ?", "iot_manual").
				Updates(orderUpdates).Error

			if err != nil {
				return err
			}
		}

		if len(manualUpdates) > 0 {
			err := tx.
				Model(&models.OrderManual{}).
				Where("order_id = ?", id).
				Updates(manualUpdates).Error

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *IOTManualRepository) Delete(order *models.Order) error {
	return r.DB.Delete(order).Error
}

func (r *IOTManualRepository) CreateFull(
	req *models.IOTManualCreateRequest,
	order *models.Order,
) error {

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		manual := models.OrderManual{
			OrderID:                   order.ID,
			NomorFormulirBerlangganan: req.NomorFormulirBerlangganan,
			NomorSuratPenawaranHarga:  req.NomorSuratPenawaranHarga,
			BulanPengirimanSPH:        req.BulanPengirimanSPH,
			WaktuPengiriman:           req.WaktuPengiriman,
			DokumenFullSign:           req.DokumenFullSign,
			NamaSurat:                 req.NamaSurat,
			NomorKontrakBerlangganan:  req.NomorKontrakBerlangganan,
		}

		if err := tx.Create(&manual).Error; err != nil {
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
				OrderID:           order.ID,
				NoInvoiceKUT:      req.Procurement.NoInvoiceKUT,
				TanggalInvoiceKUT: req.Procurement.TanggalInvoiceKUT,
			}

			if err := tx.Create(&procurement).Error; err != nil {
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

	if err != nil {
		return err
	}

	return nil
}
