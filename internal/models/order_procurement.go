package models

import "time"

type OrderProcurement struct {
	ID 									uint `json:"id" gorm:"primaryKey"`
	OrderID 							uint `json:"order_id" gorm:"not null"`
	NoPOKUT 							*string `json:"no_po_kut" gorm:"column:no_po_kut"`
	TanggalInvoiceKUT 					*Date `json:"tanggal_invoice_kut" gorm:"column:tanggal_invoice_kut"`
	NomorSuratPenyampaianDaftarHarga 	*string `json:"nomor_surat_penyampaian_daftar_harga" gorm:"column:nomor_surat_penyampaian_daftar_harga"`
	NomorFormulirPembelian 				*string `json:"nomor_formulir_pembelian" gorm:"column:nomor_formulir_pembelian"`
	NoInvoiceKUT 						*string `json:"no_invoice_kut" gorm:"column:no_invoice_kut"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderProcurement) TableName() string {
	return "order_procurement"
}

