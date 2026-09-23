package models

import "time"

type OrderDocument struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	JenisDokumen string `json:"jenis_dokumen" gorm:"not null"`

	TanggalDokumen *Date `json:"tanggal_dokumen"`

	NomorDokumen *string `json:"nomor_dokumen"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderDocument) TableName() string {
	return "order_documents"
}