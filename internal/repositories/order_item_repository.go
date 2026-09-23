package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{DB: db}
}

func (r *OrderItemRepository) Create(item *models.OrderItem) error {
	return r.DB.Create(item).Error
}

func (r *OrderItemRepository) CreateMany(items []models.OrderItem) error {
	if len(items) == 0 {
		return nil
	}

	return r.DB.Create(&items).Error
}

func (r *OrderItemRepository) FindByID(id uint) (*models.OrderItem, error) {
	var item models.OrderItem

	err := r.DB.
		Preload("Produk").
		First(&item, id).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *OrderItemRepository) FindByOrderID(orderID uint) ([]models.OrderItem, error) {
	var items []models.OrderItem

	err := r.DB.
		Where("order_id = ?", orderID).
		Preload("Produk").
		Order("id ASC").
		Find(&items).Error

	return items, err
}

func (r *OrderItemRepository) Update(item *models.OrderItem) error {
	return r.DB.Save(item).Error
}

func (r *OrderItemRepository) Delete(id uint) error {
	return r.DB.Delete(&models.OrderItem{}, id).Error
}

func (r *OrderItemRepository) DeleteByOrderID(orderID uint) error {
	return r.DB.
		Where("order_id = ?", orderID).
		Delete(&models.OrderItem{}).Error
}