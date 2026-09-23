package master

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type EkspedisiRepository struct {
	DB *gorm.DB
}

func NewEkspedisiRepository(db *gorm.DB) *EkspedisiRepository {
	return &EkspedisiRepository{DB: db}
}

func (r *EkspedisiRepository) Create(data *models.MasterEkspedisi) error {
	return r.DB.Create(data).Error
}

func (r *EkspedisiRepository) FindAll() ([]models.MasterEkspedisi, error) {
	var data []models.MasterEkspedisi

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *EkspedisiRepository) FindActive() ([]models.MasterEkspedisi, error) {
	var data []models.MasterEkspedisi

	err := r.DB.
		Where("status = ?", 1).
		Order("nama_ekspedisi ASC").
		Find(&data).Error

	return data, err
}

func (r *EkspedisiRepository) FindByID(id uint) (*models.MasterEkspedisi, error) {
	var data models.MasterEkspedisi

	err := r.DB.First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *EkspedisiRepository) FindByNama(nama string) (*models.MasterEkspedisi, error) {
	var data models.MasterEkspedisi

	err := r.DB.
		Where("nama_ekspedisi = ?", strings.TrimSpace(nama)).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *EkspedisiRepository) Update(data *models.MasterEkspedisi) error {
	return r.DB.Save(data).Error
}

func (r *EkspedisiRepository) Delete(id uint) error {
	return r.DB.Delete(&models.MasterEkspedisi{}, id).Error
}