package models

import "time"

type Payment struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	JumlahUangMasuk float64 `json:"jumlah_uang_masuk" gorm:"not null"`

	TanggalUangMasuk *Date `json:"tanggal_uang_masuk"`

	Status *string `json:"status"`

	BuktiPembayaran *string `json:"bukti_pembayaran"`

	Rekening *string `json:"rekening"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}