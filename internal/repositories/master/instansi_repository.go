package master

import (
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type InstansiRepository struct {
	DB *gorm.DB
}

func NewInstansiRepository(db *gorm.DB) *InstansiRepository {
	return &InstansiRepository{DB: db}
}

func (r *InstansiRepository) Create(data *models.MasterInstansi) error {
	return r.DB.Create(data).Error
}

func (r *InstansiRepository) FindAll() ([]models.MasterInstansi, error) {
	var data []models.MasterInstansi

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *InstansiRepository) FindByID(id uint) (*models.MasterInstansi, error) {
	var data models.MasterInstansi

	err := r.DB.First(&data, id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *InstansiRepository) Search(keyword string) ([]models.MasterInstansi, error) {
	var data []models.MasterInstansi

	keyword = strings.TrimSpace(keyword)

	query := r.DB

	if keyword != "" {
		query = query.Where(
			"nama_instansi LIKE ?",
			"%"+keyword+"%",
		)
	}

	err := query.
		Order("nama_instansi ASC").
		Find(&data).Error

	return data, err
}

func (r *InstansiRepository) Update(data *models.MasterInstansi) error {
	return r.DB.Save(data).Error
}

func (r *InstansiRepository) Delete(id uint) error {
	return r.DB.Delete(&models.MasterInstansi{}, id).Error
}