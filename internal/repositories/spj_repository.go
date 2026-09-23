package repositories

import (
	"errors"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type SPJRepository struct {
	DB *gorm.DB
}

func NewSPJRepository(db *gorm.DB) *SPJRepository {
	return &SPJRepository{DB: db}
}

func (r *SPJRepository) Create(data *models.SPJ) error {
	return r.DB.Create(data).Error
}

func (r *SPJRepository) FindByID(id uint) (*models.SPJ, error) {
	var data models.SPJ

	err := r.DB.
		Preload("Order").
		Preload("Documents").
		First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *SPJRepository) FindByOrderID(orderID uint) (*models.SPJ, error) {
	var data models.SPJ

	err := r.DB.
		Where("order_id = ?", orderID).
		Preload("Documents").
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *SPJRepository) FindAll() ([]models.SPJ, error) {
	var data []models.SPJ

	err := r.DB.
		Preload("Order").
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *SPJRepository) Update(data *models.SPJ) error {
	return r.DB.Save(data).Error
}

func (r *SPJRepository) Delete(id uint) error {
	return r.DB.Delete(&models.SPJ{}, id).Error
}