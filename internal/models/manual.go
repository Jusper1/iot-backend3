package models

import "time"

type ManualOrder struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"uniqueIndex"`

	NamaSurat          string `json:"nama_surat"`
	DokumenFullSign    string `json:"dokumen_full_sign"`
	WaktuPengiriman    string `json:"waktu_pengiriman"`

	BulanPengirimanSPH             string `json:"bulan_pengiriman_sph"`
	NomorSuratPenawaranHarga       string `json:"nomor_surat_penawaran_harga"`
	NomorFormulirBerlangganan      string `json:"nomor_formulir_berlangganan"`
	NomorKontrakBerlangganan       string `json:"nomor_kontrak_berlangganan"`

	NoPO               string `json:"no_po"`
	TanggalPO          *Date  `json:"tanggal_po"`
	NoBAST             string `json:"no_bast"`
	TanggalBAST        *Date  `json:"tanggal_bast"`
	PeriodeLangganan   string `json:"periode_langganan"`
	NomorInvoice       string `json:"nomor_invoice"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ManualOrder) TableName() string {
	return "manual"
}