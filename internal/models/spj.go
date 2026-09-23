package models

import "time"

type SPJ struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	KebutuhanSPJ *string `json:"kebutuhan_spj"`

	JumlahRangkap int `json:"jumlah_rangkap" gorm:"not null"`

	JenisKertas *string `json:"jenis_kertas"`

	TanggalPrint *Date `json:"tanggal_print"`

	TanggalUpdateList *Date `json:"tanggal_update_list"`

	TanggalParaf *Date `json:"tanggal_paraf"`

	TanggalSign *Date `json:"tanggal_sign"`

	TanggalPengiriman *Date `json:"tanggal_pengiriman"`

	JenisFile *string `json:"jenis_file"`

	PICPrint *string `json:"pic_print"`

	Status *string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (SPJ) TableName() string {
	return "spj"
}