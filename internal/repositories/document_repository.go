package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type DocumentRepository struct {
	DB *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

func (r *DocumentRepository) Create(data *models.Document) error {
	return r.DB.Create(data).Error
}

func (r *DocumentRepository) FindByID(id uint) (*models.Document, error) {
	var data models.Document

	err := r.DB.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *DocumentRepository) FindByOrderID(orderID uint) ([]models.Document, error) {
	var data []models.Document

	err := r.DB.
		Where("order_id = ?", orderID).
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *DocumentRepository) FindBySPJID(spjID uint) ([]models.Document, error) {
	var data []models.Document

	err := r.DB.
		Where("spj_id = ?", spjID).
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *DocumentRepository) FindAll() ([]models.Document, error) {
	var data []models.Document

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *DocumentRepository) Update(data *models.Document) error {
	return r.DB.Save(data).Error
}

func (r *DocumentRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Document{}, id).Error
}