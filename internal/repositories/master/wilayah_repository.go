package master

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type WilayahRepository struct {
	DB *gorm.DB
}

func NewWilayahRepository(db *gorm.DB) *WilayahRepository {
	return &WilayahRepository{DB: db}
}

func (r *WilayahRepository) Create(data *models.MasterWilayah) error {
	return r.DB.Create(data).Error
}

func (r *WilayahRepository) FindAll() ([]models.MasterWilayah, error) {
	var data []models.MasterWilayah

	err := r.DB.
		Order("provinsi ASC, kota_kab ASC").
		Find(&data).Error

	return data, err
}

func (r *WilayahRepository) FindByID(id uint) (*models.MasterWilayah, error) {
	var data models.MasterWilayah

	err := r.DB.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *WilayahRepository) FindByProvinsi(provinsi string) ([]models.MasterWilayah, error) {
	var data []models.MasterWilayah

	err := r.DB.
		Where("provinsi = ?", provinsi).
		Order("kota_kab ASC").
		Find(&data).Error

	return data, err
}

func (r *WilayahRepository) Update(data *models.MasterWilayah) error {
	return r.DB.Save(data).Error
}

func (r *WilayahRepository) Delete(id uint) error {
	return r.DB.Delete(&models.MasterWilayah{}, id).Error
}