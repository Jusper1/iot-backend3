package master

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type ProdukRepository struct {
	DB *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{DB: db}
}

func (r *ProdukRepository) Create(data *models.MasterProduk) error {
	return r.DB.Create(data).Error
}

func (r *ProdukRepository) FindAll() ([]models.MasterProduk, error) {
	var data []models.MasterProduk

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *ProdukRepository) FindActive() ([]models.MasterProduk, error) {
	var data []models.MasterProduk

	err := r.DB.
		Where("status = ?", 1).
		Order("nama_produk ASC").
		Find(&data).Error

	return data, err
}

func (r *ProdukRepository) FindByID(id uint) (*models.MasterProduk, error) {
	var data models.MasterProduk

	err := r.DB.First(&data, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *ProdukRepository) FindByKode(kode string) (*models.MasterProduk, error) {
	var data models.MasterProduk

	err := r.DB.
		Where("kode_produk = ?", strings.TrimSpace(kode)).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &data, err
}

func (r *ProdukRepository) Update(data *models.MasterProduk) error {
	return r.DB.Save(data).Error
}

func (r *ProdukRepository) Delete(id uint) error {
	return r.DB.Delete(&models.MasterProduk{}, id).Error
}