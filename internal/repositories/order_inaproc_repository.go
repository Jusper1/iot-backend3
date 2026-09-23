package repositories

import (
	"errors"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderInaprocRepository struct {
	DB *gorm.DB
}

func NewOrderInaprocRepository(db *gorm.DB) *OrderInaprocRepository {
	return &OrderInaprocRepository{DB: db}
}

func (r *OrderInaprocRepository) Create(data *models.OrderInaproc) error {
	return r.DB.Create(data).Error
}

func (r *OrderInaprocRepository) FindByID(id uint) (*models.OrderInaproc, error) {
	var data models.OrderInaproc

	err := r.DB.First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderInaprocRepository) FindByOrderID(orderID uint) (*models.OrderInaproc, error) {
	var data models.OrderInaproc

	err := r.DB.
		Where("order_id = ?", orderID).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *OrderInaprocRepository) Update(data *models.OrderInaproc) error {
	return r.DB.Save(data).Error
}

func (r *OrderInaprocRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderInaproc{}).Error
}