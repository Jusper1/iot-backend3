package models

import "time"

type OrderProcurement struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	NoInvoiceKUT *string `json:"no_invoice_kut"`

	NomorFormulirPembelian *string `json:"nomor_formulir_pembelian"`

	NomorSuratPenyampaianDaftarHarga *string `json:"nomor_surat_penyampaian_daftar_harga"`

	TanggalInvoiceKUT *Date `json:"tanggal_invoice_kut"`

	NoPOKUT *string `json:"no_po_kut"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderProcurement) TableName() string {
	return "order_procurement"
}