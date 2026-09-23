package master

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type PICRepository struct {
	DB *gorm.DB
}

func NewPICRepository(db *gorm.DB) *PICRepository {
	return &PICRepository{DB: db}
}

func (r *PICRepository) Create(data *models.MasterPIC) error {
	return r.DB.Create(data).Error
}

func (r *PICRepository) FindAll() ([]models.MasterPIC, error) {
	var data []models.MasterPIC

	err := r.DB.
		Preload("Instansi").
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *PICRepository) FindByID(id uint) (*models.MasterPIC, error) {
	var data models.MasterPIC

	err := r.DB.
		Preload("Instansi").
		First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PICRepository) FindByInstansiID(instansiID uint) ([]models.MasterPIC, error) {
	var data []models.MasterPIC

	err := r.DB.
		Where("instansi_id = ?", instansiID).
		Order("nama_pic ASC").
		Find(&data).Error

	return data, err
}

func (r *PICRepository) Update(data *models.MasterPIC) error {
	return r.DB.Save(data).Error
}

func (r *PICRepository) Delete(id uint) error {
	return r.DB.Delete(&models.MasterPIC{}, id).Error
}