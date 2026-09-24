package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type IOTInaprocRepository struct {
	DB *gorm.DB
}

func NewIOTInaprocRepository(db *gorm.DB) *IOTInaprocRepository {
	return &IOTInaprocRepository{
		DB: db,
	}
}

func (r *IOTInaprocRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Payments").
		Where("jenis_order = ?", "inaproc").
		Where("kategori_order = ?", "iot_inaproc").
		Order("id DESC").
		Find(&orders).Error

	return orders, err
}

func (r *IOTInaprocRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Payments").
		Where("jenis_order = ?", "inaproc").
		Where("kategori_order = ?", "iot_inaproc").
		First(&order, id).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *IOTInaprocRepository) FindByKode(kode string) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Procurement").
		Preload("Payments").
		Where("kode_order = ?", kode).
		Where("jenis_order = ?", "inaproc").
		Where("kategori_order = ?", "iot_inaproc").
		First(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *IOTInaprocRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *IOTInaprocRepository) Update(
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
		"periode_langganan":  true,
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
		Where("jenis_order = ?", "inaproc").
		Where("kategori_order = ?", "iot_inaproc").
		Updates(updates).Error
}

func (r *IOTInaprocRepository) Delete(order *models.Order) error {
	return r.DB.Delete(order).Error
}

func (r *IOTInaprocRepository) CreateFull(
	req *models.IOTInaprocCreateRequest,
	order *models.Order,
) error {

	err := r.DB.Transaction(func(tx *gorm.DB) error {

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