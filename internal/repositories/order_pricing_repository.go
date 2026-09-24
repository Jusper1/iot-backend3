package repositories

import (
	"errors"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderPricingRepository struct {
	DB *gorm.DB
}

func NewOrderPricingRepository(db *gorm.DB) *OrderPricingRepository {
	return &OrderPricingRepository{DB: db}
}

func (r *OrderPricingRepository) Create(pricing *models.OrderPricing) error {
	return r.DB.Create(pricing).Error
}

func (r *OrderPricingRepository) FindByID(id uint) (*models.OrderPricing, error) {
	var pricing models.OrderPricing

	err := r.DB.First(&pricing, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &pricing, err
}

func (r *OrderPricingRepository) FindByOrderID(orderID uint) (*models.OrderPricing, error) {
	var pricing models.OrderPricing

	err := r.DB.
		Where("order_id = ?", orderID).
		First(&pricing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &pricing, err
}

func (r *OrderPricingRepository) Update(
	id uint,
	data map[string]interface{},
) error {
	allowedFields := map[string]bool{
		"order_id":                       true,
		"tipe_timbangan":                true,
		"harga_produk":                  true,
		"harga_ppn":                     true,
		"harga_ongkir_kut":              true,
		"harga_ppn_ongkir":              true,
		"total_harga_ongkir":            true,
		"total_harga_jual":              true,
		"harga_produk_reseller":         true,
		"harga_ppn_reseller":            true,
		"harga_ongkir_reseller":         true,
		"harga_ppn_ongkir_reseller":     true,
		"total_harga_ongkir_reseller":   true,
		"total_harga_reseller":          true,
	}

	updates := make(map[string]interface{})

	for field, value := range data {
		if allowedFields[field] {
			updates[field] = value
		}
	}

	if len(updates) == 0 {
		return gorm.ErrInvalidData
	}

	return r.DB.
		Model(&models.OrderPricing{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *OrderPricingRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderPricing{}).Error
}