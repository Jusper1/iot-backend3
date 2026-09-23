package repositories

import (

		"errors"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)



type OrderProcurementRepository struct {
	DB *gorm.DB
}

func NewOrderProcurementRepository(db *gorm.DB) *OrderProcurementRepository {
	return &OrderProcurementRepository{DB: db}
}

func (r *OrderProcurementRepository) Create(data *models.OrderProcurement) error {
	return r.DB.Create(data).Error
}

func (r *OrderProcurementRepository) FindByID(id uint) (*models.OrderProcurement, error) {
	var data models.OrderProcurement

	err := r.DB.First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderProcurementRepository) FindByOrderID(orderID uint) (*models.OrderProcurement, error) {
	var data models.OrderProcurement

	err := r.DB.
		Where("order_id = ?", orderID).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderProcurementRepository) Update(data *models.OrderProcurement) error {
	return r.DB.Save(data).Error
}

func (r *OrderProcurementRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderProcurement{}).Error
}