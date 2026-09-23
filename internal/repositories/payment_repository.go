package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) Create(data *models.Payment) error {
	return r.DB.Create(data).Error
}

func (r *PaymentRepository) FindByID(id uint) (*models.Payment, error) {
	var data models.Payment

	err := r.DB.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PaymentRepository) FindByOrderID(orderID uint) ([]models.Payment, error) {
	var data []models.Payment

	err := r.DB.
		Where("order_id = ?", orderID).
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *PaymentRepository) FindAll() ([]models.Payment, error) {
	var data []models.Payment

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *PaymentRepository) Update(data *models.Payment) error {
	return r.DB.Save(data).Error
}

func (r *PaymentRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Payment{}, id).Error
}