package models

import "time"

type Inaproc struct {
	ID               uint    `json:"id" gorm:"primaryKey"`
	OrderID          uint    `json:"order_id" gorm:"uniqueIndex"`
	NoPO             string  `json:"no_po"`
	TanggalPO        *Date   `json:"tanggal_po"`
	NoBAST           string  `json:"no_bast"`
	TanggalBAST      *Date   `json:"tanggal_bast"`
	NoInvoiceInaproc string  `json:"no_invoice_inaproc"`
	KodeBayar        string  `json:"kode_bayar"`
	NSFP             string  `json:"nsfp"`
	NoInvoiceKUT     string  `json:"no_invoice_kut"`
	TanggalUangMasuk *Date   `json:"tanggal_uang_masuk"`
	JumlahUangMasuk  float64 `json:"jumlah_uang_masuk"`
	Rekening         string  `json:"rekening"`
	// BulanPengiriman     string     `json:"bulan_pengiriman"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Inaproc) TableName() string {
	return "order_inaproc"
}
