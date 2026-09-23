package models

import "time"

type OrderPricing struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	TipeTimbangan *string `json:"tipe_timbangan"`

	HargaProduk float64 `json:"harga_produk" gorm:"not null"`
	HargaPPN    float64 `json:"harga_ppn" gorm:"not null"`

	HargaOngkirKUT float64 `json:"harga_ongkir_kut" gorm:"not null"`
	HargaPPNOngkir float64 `json:"harga_ppn_ongkir" gorm:"not null"`

	TotalHargaOngkir float64 `json:"total_harga_ongkir" gorm:"not null"`
	TotalHargaJual   float64 `json:"total_harga_jual" gorm:"not null"`

	HargaProdukReseller float64 `json:"harga_produk_reseller" gorm:"not null"`
	HargaPPNReseller    float64 `json:"harga_ppn_reseller" gorm:"not null"`

	HargaOngkirReseller       float64 `json:"harga_ongkir_reseller" gorm:"not null"`
	HargaPPNOngkirReseller    float64 `json:"harga_ppn_ongkir_reseller" gorm:"not null"`

	TotalHargaOngkirReseller float64 `json:"total_harga_ongkir_reseller" gorm:"not null"`
	TotalHargaReseller       float64 `json:"total_harga_reseller" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderPricing) TableName() string {
	return "order_pricing"
}