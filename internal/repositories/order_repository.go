package repositories

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Order("id DESC").
		Find(&orders).Error

	return orders, err
}

func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Preload("Instansi").
		Preload("PIC").
		Preload("Items").
		Preload("Items.Produk").
		Preload("Pricing").
		Preload("Inaproc").
		Preload("Manual").
		Preload("Procurement").
		Preload("Documents").
		Preload("Payments").
		Preload("Shipments").
		Preload("Shipments.Wilayah").
		Preload("Shipments.Ekspedisi").
		Preload("SPJ").
		First(&order, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &order, err
}

func (r *OrderRepository) FindByKode(kode string) (*models.Order, error) {
	var order models.Order

	err := r.DB.
		Where("kode_order = ?", strings.TrimSpace(kode)).
		First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &order, err
}

func (r *OrderRepository) ExistsByKode(kode string) (bool, error) {
	var count int64

	err := r.DB.
		Model(&models.Order{}).
		Where("kode_order = ?", strings.TrimSpace(kode)).
		Count(&count).Error

	return count > 0, err
}

func (r *OrderRepository) Update(order *models.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Order{}, id).Error
}