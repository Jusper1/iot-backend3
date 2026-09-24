package models

type IOTInaprocCreateRequest struct {
	KodeOrder string `json:"kode_order"`

	InstansiID uint `json:"instansi_id"`
	PicID      uint `json:"pic_id"`

	Status     *string `json:"status"`
	StatusOdoo *string `json:"status_odoo"`

	NSFP             *string `json:"nsfp"`
	NoBAST           *string `json:"no_bast"`
	KodeBayar        *string `json:"kode_bayar"`
	NoInvoiceInaproc *string `json:"no_invoice_inaproc"`

	Keterangan       *string `json:"keterangan"`
	TanggalPO        *Date   `json:"tanggal_po"`
	NoPO             *string `json:"no_po"`
	PeriodeLangganan *string `json:"periode_langganan"`
	TanggalBAST      *Date   `json:"tanggal_bast"`

	Items []IOTInaprocItemRequest `json:"items"`

	Procurement *IOTInaprocProcurementRequest `json:"procurement"`

	Payment *IOTInaprocPaymentRequest `json:"payment"`
}

type IOTInaprocItemRequest struct {
	ProdukID uint    `json:"produk_id"`
	Qty      int64   `json:"qty"`
	Harga    float64 `json:"harga"`
	PPN      float64 `json:"ppn"`
	Subtotal float64 `json:"subtotal"`
}

type IOTInaprocProcurementRequest struct {
	NoInvoiceKUT      *string `json:"no_invoice_kut"`
	TanggalInvoiceKUT *Date   `json:"tanggal_invoice_kut"`
}

type IOTInaprocPaymentRequest struct {
	JumlahUangMasuk  float64 `json:"jumlah_uang_masuk"`
	TanggalUangMasuk *Date   `json:"tanggal_uang_masuk"`
	Status           *string `json:"status"`
	Rekening         *string `json:"rekening"`
	BuktiPembayaran  *string `json:"bukti_pembayaran"`
}