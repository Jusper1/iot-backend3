package repositories

import (
	"errors"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderManualRepository struct {
	DB *gorm.DB
}

func NewOrderManualRepository(db *gorm.DB) *OrderManualRepository {
	return &OrderManualRepository{DB: db}
}

func (r *OrderManualRepository) Create(data *models.OrderManual) error {
	return r.DB.Create(data).Error
}

func (r *OrderManualRepository) FindByID(id uint) (*models.OrderManual, error) {
	var data models.OrderManual

	err := r.DB.First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderManualRepository) FindByOrderID(orderID uint) (*models.OrderManual, error) {
	var data models.OrderManual

	err := r.DB.
		Where("order_id = ?", orderID).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderManualRepository) Update(data *models.OrderManual) error {
	return r.DB.Save(data).Error
}

func (r *OrderManualRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderManual{}).Error
}