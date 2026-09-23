package models

import "time"

type Shipment struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	WilayahID   *uint `json:"wilayah_id"`
	EkspedisiID *uint `json:"ekspedisi_id"`

	Resi *string `json:"resi"`

	Berat float64 `json:"berat" gorm:"not null"`

	Ongkir float64 `json:"ongkir" gorm:"not null"`

	TanggalKirim *Date `json:"tanggal_kirim"`

	TanggalDiterima *Date `json:"tanggal_diterima"`

	StatusPengiriman *string `json:"status_pengiriman"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Wilayah   *MasterWilayah   `json:"wilayah,omitempty" gorm:"foreignKey:WilayahID"`
	Ekspedisi *MasterEkspedisi `json:"ekspedisi,omitempty" gorm:"foreignKey:EkspedisiID"`
}

func (Shipment) TableName() string {
	return "shipments"
}