package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderExportRepository struct {
	DB *gorm.DB
}

func NewOrderExportRepository(db *gorm.DB) *OrderExportRepository {
	return &OrderExportRepository{
		DB: db,
	}
}

func (r *OrderExportRepository) FindForExport(
	filter models.OrderExportFilter,
) ([]models.Order, error) {
	var orders []models.Order

	query := r.DB.Model(&models.Order{}).
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Pricing").
		Preload("Manual").
		Preload("Procurement").
		Preload("Shipments").
		Preload("Shipments.Wilayah").
		Preload("Shipments.Ekspedisi").
		Preload("Payments")

	if filter.KodeOrder != "" {
		query = query.Where("orders.kode_order LIKE ?", "%"+filter.KodeOrder+"%")
	}

	if filter.Status != "" {
		query = query.Where("orders.status = ?", filter.Status)
	}

	if filter.TanggalPODari != "" {
		query = query.Where("orders.tanggal_po >= ?", filter.TanggalPODari)
	}
	if filter.TanggalPOSampai != "" {
		query = query.Where("orders.tanggal_po <= ?", filter.TanggalPOSampai)
	}

	if len(filter.KategoriOrder) > 0 {
		query = query.Where("orders.kategori_order IN ?", filter.KategoriOrder)
	}

	if filter.NamaInstansi != "" || filter.Provinsi != "" {
		instansiQuery := r.DB.Model(&models.MasterInstansi{}).Select("id")

		if filter.NamaInstansi != "" {
			instansiQuery = instansiQuery.Where("nama_instansi LIKE ?", "%"+filter.NamaInstansi+"%")
		}
		if filter.Provinsi != "" {
			instansiQuery = instansiQuery.Where("provinsi = ?", filter.Provinsi)
		}

		query = query.Where("orders.instansi_id IN (?)", instansiQuery)
	}

	if filter.TanggalUangMasukDari != "" || filter.TanggalUangMasukSampai != "" {
		paymentQuery := r.DB.Model(&models.Payment{}).Select("order_id")

		if filter.TanggalUangMasukDari != "" {
			paymentQuery = paymentQuery.Where("tanggal_uang_masuk >= ?", filter.TanggalUangMasukDari)
		}
		if filter.TanggalUangMasukSampai != "" {
			paymentQuery = paymentQuery.Where("tanggal_uang_masuk <= ?", filter.TanggalUangMasukSampai)
		}

		query = query.Where("orders.id IN (?)", paymentQuery)
	}

	err := query.Order("orders.id ASC").Find(&orders).Error

	return orders, err
}