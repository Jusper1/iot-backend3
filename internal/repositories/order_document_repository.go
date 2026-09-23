package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderDocumentRepository struct {
	DB *gorm.DB
}

func NewOrderDocumentRepository(db *gorm.DB) *OrderDocumentRepository {
	return &OrderDocumentRepository{DB: db}
}

func (r *OrderDocumentRepository) Create(data *models.OrderDocument) error {
	return r.DB.Create(data).Error
}

func (r *OrderDocumentRepository) FindByID(id uint) (*models.OrderDocument, error) {
	var data models.OrderDocument

	err := r.DB.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *OrderDocumentRepository) FindByOrderID(orderID uint) ([]models.OrderDocument, error) {
	var data []models.OrderDocument

	err := r.DB.
		Where("order_id = ?", orderID).
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *OrderDocumentRepository) Update(data *models.OrderDocument) error {
	return r.DB.Save(data).Error
}

func (r *OrderDocumentRepository) Delete(id uint) error {
	return r.DB.Delete(&models.OrderDocument{}, id).Error
}

func (r *OrderDocumentRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderDocument{}).Error
}