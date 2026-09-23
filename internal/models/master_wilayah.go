package models

import "time"

type MasterWilayah struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Provinsi string `json:"provinsi" gorm:"not null"`

	KotaKab string `json:"kota_kab" gorm:"not null"`

	Alamat *string `json:"alamat"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MasterWilayah) TableName() string {
	return "master_wilayah"
}