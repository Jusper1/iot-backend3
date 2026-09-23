package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type ShipmentRepository struct {
	DB *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) *ShipmentRepository {
	return &ShipmentRepository{DB: db}
}

func (r *ShipmentRepository) Create(data *models.Shipment) error {
	return r.DB.Create(data).Error
}

func (r *ShipmentRepository) FindByID(id uint) (*models.Shipment, error) {
	var data models.Shipment

	err := r.DB.
		Preload("Wilayah").
		Preload("Ekspedisi").
		First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ShipmentRepository) FindByOrderID(orderID uint) ([]models.Shipment, error) {
	var data []models.Shipment

	err := r.DB.
		Where("order_id = ?", orderID).
		Preload("Wilayah").
		Preload("Ekspedisi").
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *ShipmentRepository) FindAll() ([]models.Shipment, error) {
	var data []models.Shipment

	err := r.DB.
		Preload("Wilayah").
		Preload("Ekspedisi").
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *ShipmentRepository) Update(data *models.Shipment) error {
	return r.DB.Save(data).Error
}

func (r *ShipmentRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Shipment{}, id).Error
}