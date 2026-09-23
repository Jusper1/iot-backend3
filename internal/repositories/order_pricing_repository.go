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

func (r *OrderPricingRepository) Update(pricing *models.OrderPricing) error {
	return r.DB.Save(pricing).Error
}

func (r *OrderPricingRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderPricing{}).Error
}