package models

type IOTManualCreateRequest struct {
	KodeOrder string `json:"kode_order"`

	InstansiID uint `json:"instansi_id"`
	PicID      uint `json:"pic_id"`

	Status     *string `json:"status"`
	StatusOdoo *string `json:"status_odoo"`

	NSFP             *string `json:"nsfp"`
	NoBAST           *string `json:"no_bast"`
	Keterangan       *string `json:"keterangan"`
	TanggalPO        *Date   `json:"tanggal_po"`
	NoPO             *string `json:"no_po"`
	PeriodeLangganan *string `json:"periode_langganan"`
	TanggalBAST      *Date   `json:"tanggal_bast"`

	NomorFormulirBerlangganan *string `json:"nomor_formulir_berlangganan"`
	NomorSuratPenawaranHarga  *string `json:"nomor_surat_penawaran_harga"`
	BulanPengirimanSPH        *string `json:"bulan_pengiriman_sph"`
	WaktuPengiriman           *string `json:"waktu_pengiriman"`
	DokumenFullSign           *string `json:"dokumen_full_sign"`
	NamaSurat                 *string `json:"nama_surat"`
	NomorKontrakBerlangganan  *string `json:"nomor_kontrak_berlangganan"`

	Items []IOTManualItemRequest `json:"items"`

	Procurement *IOTManualProcurementRequest `json:"procurement"`

	Payment *IOTManualPaymentRequest `json:"payment"`
}

type IOTManualItemRequest struct {
	ProdukID uint    `json:"produk_id"`
	Qty      int64   `json:"qty"`
	Harga    float64 `json:"harga"`
	PPN      float64 `json:"ppn"`
	Subtotal float64 `json:"subtotal"`
}

type IOTManualProcurementRequest struct {
	NoInvoiceKUT      *string `json:"no_invoice_kut"`
	TanggalInvoiceKUT *Date   `json:"tanggal_invoice_kut"`
}

type IOTManualPaymentRequest struct {
	JumlahUangMasuk  float64 `json:"jumlah_uang_masuk"`
	TanggalUangMasuk *Date   `json:"tanggal_uang_masuk"`
	Status           *string `json:"status"`
	Rekening         *string `json:"rekening"`
	BuktiPembayaran  *string `json:"bukti_pembayaran"`
}
