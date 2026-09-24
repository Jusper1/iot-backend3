package models

import "time"

type MasterProduk struct {
	ID uint `json:"id" gorm:"primaryKey"`

	KodeProduk 	string `json:"kode_produk" gorm:"not null;uniqueIndex"`
	NamaProduk 	string `json:"nama_produk" gorm:"not null"`
	Harga 		float64 `json:"harga" gorm:"not null"`
	Status 		string `json:"status" gorm:"not null"`
	JenisProduk *string `json:"jenis_produk"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MasterProduk) TableName() string {
	return "master_produk"
}