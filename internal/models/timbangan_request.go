package models

type TimbanganCreateRequest struct {
	KodeOrder     string `json:"kode_order"`
	JenisOrder    string `json:"jenis_order"`    // "inaproc" atau "manual"
	KategoriOrder string `json:"kategori_order"` // timbangan_inaproc | timbangan_manual | rcw_360 | rcw_800w

	InstansiID uint `json:"instansi_id"`
	PicID      uint `json:"pic_id"`

	Status     *string `json:"status"`
	StatusOdoo *string `json:"status_odoo"`

	NSFP             *string `json:"nsfp"`
	NoBAST           *string `json:"no_bast"`
	KodeBayar        *string `json:"kode_bayar"`
	NoInvoiceInaproc *string `json:"no_invoice_inaproc"`

	Keterangan  *string `json:"keterangan"`
	TanggalPO   *Date   `json:"tanggal_po"`
	NoPO        *string `json:"no_po"`
	TanggalBAST *Date   `json:"tanggal_bast"`

	Items []TimbanganItemRequest `json:"items"`


	Procurement *TimbanganProcurementRequest `json:"procurement"`
	Pricing *TimbanganPricingRequest `json:"pricing"`
	Shipment *TimbanganShipmentRequest `json:"shipment"`
	Payment *TimbanganPaymentRequest `json:"payment"`
}

type TimbanganItemRequest struct {
	ProdukID uint    `json:"produk_id"`
	Qty      int64   `json:"qty"`
	Harga    float64 `json:"harga"`
	PPN      float64 `json:"ppn"`
	Subtotal float64 `json:"subtotal"`
}

type TimbanganProcurementRequest struct {
	NoPOKUT                           *string `json:"no_po_kut"`
	TanggalInvoiceKUT                 *Date   `json:"tanggal_invoice_kut"`
	NomorSuratPenyampaianDaftarHarga  *string `json:"nomor_surat_penyampaian_daftar_harga"`
	NomorFormulirPembelian            *string `json:"nomor_formulir_pembelian"`
	NoInvoiceKUT                      *string `json:"no_invoice_kut"`
}

type TimbanganPricingRequest struct {
	TipeTimbangan            *string `json:"tipe_timbangan"`
	HargaProduk              float64 `json:"harga_produk"`
	HargaPPN                 float64 `json:"harga_ppn"`
	HargaOngkirKUT           float64 `json:"harga_ongkir_kut"`
	HargaPPNOngkir           float64 `json:"harga_ppn_ongkir"`
	TotalHargaOngkir         float64 `json:"total_harga_ongkir"`
	TotalHargaJual           float64 `json:"total_harga_jual"`
	HargaProdukReseller      float64 `json:"harga_produk_reseller"`
	HargaPPNReseller         float64 `json:"harga_ppn_reseller"`
	HargaOngkirReseller      float64 `json:"harga_ongkir_reseller"`
	HargaPPNOngkirReseller   float64 `json:"harga_ppn_ongkir_reseller"`
	TotalHargaOngkirReseller float64 `json:"total_harga_ongkir_reseller"`
	TotalHargaReseller       float64 `json:"total_harga_reseller"`
}

type TimbanganShipmentRequest struct {
	WilayahID        *uint   `json:"wilayah_id"`
	EkspedisiID      *uint   `json:"ekspedisi_id"`
	Resi             *string `json:"resi"`
	Berat            float64 `json:"berat"`
	Ongkir           float64 `json:"ongkir"`
	TanggalKirim     *Date   `json:"tanggal_kirim"`
	TanggalDiterima  *Date   `json:"tanggal_diterima"`
	StatusPengiriman *string `json:"status_pengiriman"`
}

type TimbanganPaymentRequest struct {
	JumlahUangMasuk  float64 `json:"jumlah_uang_masuk"`
	TanggalUangMasuk *Date   `json:"tanggal_uang_masuk"`
	Status           *string `json:"status"`
	Rekening         *string `json:"rekening"`
	BuktiPembayaran  *string `json:"bukti_pembayaran"`
}