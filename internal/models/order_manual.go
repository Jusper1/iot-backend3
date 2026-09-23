package models

import "time"

type OrderManual struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	WaktuPengiriman *Date `json:"waktu_pengiriman"`

	BulanPengirimanSPH *string `json:"bulan_pengiriman_sph"`

	NamaSurat *string `json:"nama_surat"`

	NomorSuratPenawaranHarga *string `json:"nomor_surat_penawaran_harga"`

	NomorFormulirBerlangganan *string `json:"nomor_formulir_berlangganan"`

	NomorKontrakBerlangganan *string `json:"nomor_kontrak_berlangganan"`

	DokumenFullSign *string `json:"dokumen_full_sign"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderManual) TableName() string {
	return "order_manual"
}