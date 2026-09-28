package services

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
	"iot-backend/internal/rules"
)

type OrderExportService struct {
	Repository *repositories.OrderExportRepository
}

func NewOrderExportService(
	repository *repositories.OrderExportRepository,
) *OrderExportService {
	return &OrderExportService{
		Repository: repository,
	}
}

type exportColumn struct {
	Header   string
	Kategori []string
	Value    func(o *models.Order) interface{}
}


var (
	katAll = rules.AllKategori

	katExceptIotInaproc = []string{
		rules.KategoriIoTManual,
		rules.KategoriTimbanganInaproc, rules.KategoriTimbanganManual,
		rules.KategoriRCW360, rules.KategoriRCW800W,
	}

	katExceptIotManual = []string{
		rules.KategoriIoTInaproc,
		rules.KategoriTimbanganInaproc, rules.KategoriTimbanganManual,
		rules.KategoriRCW360, rules.KategoriRCW800W,
	}

	katLangganan = []string{rules.KategoriIoTInaproc, rules.KategoriIoTManual}
	katIoTManual = []string{rules.KategoriIoTManual}

	katTimbanganRCW = []string{
		rules.KategoriTimbanganInaproc, rules.KategoriTimbanganManual,
		rules.KategoriRCW360, rules.KategoriRCW800W,
	}
)


func strVal(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func dateVal(v *models.Date) string {
	if v == nil {
		return ""
	}
	return string(*v)
}

func firstPayment(o *models.Order) *models.Payment {
	if len(o.Payments) == 0 {
		return nil
	}
	return &o.Payments[0]
}

func firstShipment(o *models.Order) *models.Shipment {
	if len(o.Shipments) == 0 {
		return nil
	}
	return &o.Shipments[0]
}


var exportColumns = []exportColumn{
	{"Kode Pemesanan", katAll, func(o *models.Order) interface{} { return o.KodeOrder }},
	{"Kategori", katAll, func(o *models.Order) interface{} { return o.KategoriOrder }},
	{"Nama Instansi/Perusahaan", katAll, func(o *models.Order) interface{} {
		if o.Instansi == nil {
			return ""
		}
		return o.Instansi.NamaInstansi
	}},
	{"Nama PIC", katAll, func(o *models.Order) interface{} {
		if o.PIC == nil {
			return ""
		}
		return o.PIC.NamaPIC
	}},
	{"Nomor Telepon PIC", katAll, func(o *models.Order) interface{} {
		if o.PIC == nil {
			return ""
		}
		return strVal(o.PIC.NoHP)
	}},
	{"Email", katExceptIotInaproc, func(o *models.Order) interface{} {
		if o.PIC == nil {
			return ""
		}
		return strVal(o.PIC.Email)
	}},
	{"NIK", katAll, func(o *models.Order) interface{} {
		if o.PIC == nil {
			return ""
		}
		return strVal(o.PIC.NIK)
	}},
	{"No NPWP", katAll, func(o *models.Order) interface{} {
		if o.Instansi == nil {
			return ""
		}
		return strVal(o.Instansi.NPWP)
	}},
	{"Quantity", katAll, func(o *models.Order) interface{} {
		if len(o.Items) == 0 {
			return ""
		}
		var total int64
		for _, it := range o.Items {
			total += it.Qty
		}
		return total
	}},
	{"Harga + PPN", katAll, func(o *models.Order) interface{} {
		if len(o.Items) == 0 {
			return ""
		}
		var total float64
		for _, it := range o.Items {
			total += it.Subtotal
		}
		return total
	}},
	{"Periode Berlangganan", katLangganan, func(o *models.Order) interface{} { return strVal(o.PeriodeLangganan) }},
	{"Status Pesanan", katAll, func(o *models.Order) interface{} { return strVal(o.Status) }},
	{"Status Odoo", katAll, func(o *models.Order) interface{} { return strVal(o.StatusOdoo) }},

	{"Bulan Pengiriman SPH", katIoTManual, func(o *models.Order) interface{} {
		if o.Manual == nil {
			return ""
		}
		return strVal(o.Manual.BulanPengirimanSPH)
	}},
	{"Nomor Surat Penawaran Harga", katIoTManual, func(o *models.Order) interface{} {
		if o.Manual == nil {
			return ""
		}
		return strVal(o.Manual.NomorSuratPenawaranHarga)
	}},
	{"Nomor Surat Penyampaian Daftar Harga", katTimbanganRCW, func(o *models.Order) interface{} {
		if o.Procurement == nil {
			return ""
		}
		return strVal(o.Procurement.NomorSuratPenyampaianDaftarHarga)
	}},
	{"Nomor Formulir Pembelian", katTimbanganRCW, func(o *models.Order) interface{} {
		if o.Procurement == nil {
			return ""
		}
		return strVal(o.Procurement.NomorFormulirPembelian)
	}},
	{"Nomor Formulir Berlangganan", katIoTManual, func(o *models.Order) interface{} {
		if o.Manual == nil {
			return ""
		}
		return strVal(o.Manual.NomorFormulirBerlangganan)
	}},
	{"Nomor Kontrak Berlangganan", katIoTManual, func(o *models.Order) interface{} {
		if o.Manual == nil {
			return ""
		}
		return strVal(o.Manual.NomorKontrakBerlangganan)
	}},
	{"Nomor PO KUT", katTimbanganRCW, func(o *models.Order) interface{} {
		if o.Procurement == nil {
			return ""
		}
		return strVal(o.Procurement.NoPOKUT)
	}},

	{"Tanggal Pesanan (PO)", katAll, func(o *models.Order) interface{} { return dateVal(o.TanggalPO) }},
	{"Tanggal BAST", katAll, func(o *models.Order) interface{} { return dateVal(o.TanggalBAST) }},
	{"Nomor BAST", katAll, func(o *models.Order) interface{} { return strVal(o.NoBAST) }},
	{"Nomor Invoice KUT", katAll, func(o *models.Order) interface{} {
		if o.Procurement == nil {
			return ""
		}
		return strVal(o.Procurement.NoInvoiceKUT)
	}},
	{"Tanggal Invoice KUT", katAll, func(o *models.Order) interface{} {
		if o.Procurement == nil {
			return ""
		}
		return dateVal(o.Procurement.TanggalInvoiceKUT)
	}},
	{"Nomor Invoice Inaproc", katExceptIotManual, func(o *models.Order) interface{} { return strVal(o.NoInvoiceInaproc) }},
	{"NSFP", katAll, func(o *models.Order) interface{} { return strVal(o.NSFP) }},

	{"Tanggal Uang Masuk", katAll, func(o *models.Order) interface{} {
		if p := firstPayment(o); p != nil {
			return dateVal(p.TanggalUangMasuk)
		}
		return ""
	}},
	{"Jumlah Uang Masuk", katAll, func(o *models.Order) interface{} {
		if p := firstPayment(o); p != nil {
			return p.JumlahUangMasuk
		}
		return ""
	}},
	{"Rekening Penerima", katAll, func(o *models.Order) interface{} {
		if p := firstPayment(o); p != nil {
			return strVal(p.Rekening)
		}
		return ""
	}},

	{"Kode Bayar", katExceptIotManual, func(o *models.Order) interface{} { return strVal(o.KodeBayar) }},
	{"Keterangan", katAll, func(o *models.Order) interface{} { return strVal(o.Keterangan) }},

	{"Alamat", katAll, func(o *models.Order) interface{} {
		if o.Instansi == nil {
			return ""
		}
		return strVal(o.Instansi.Alamat)
	}},
	{"Kota/Kab", katAll, func(o *models.Order) interface{} {
		if o.Instansi == nil {
			return ""
		}
		return strVal(o.Instansi.KotaKab)
	}},
	{"Provinsi", katAll, func(o *models.Order) interface{} {
		if o.Instansi == nil {
			return ""
		}
		return strVal(o.Instansi.Provinsi)
	}},


	{"Tipe Timbangan", katTimbanganRCW, func(o *models.Order) interface{} {
		if o.Pricing == nil {
			return ""
		}
		return strVal(o.Pricing.TipeTimbangan)
	}},
	{"Harga Produk", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaProduk })},
	{"Harga PPN", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaPPN })},
	{"Harga Ongkir KUT", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaOngkirKUT })},
	{"Harga PPN 11% Ongkir", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaPPNOngkir })},
	{"Total Harga + Ongkir (Exclude PPN)", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.TotalHargaOngkir })},
	{"Total Harga Jual (Include PPN)", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.TotalHargaJual })},

	{"Wilayah Pengiriman", katTimbanganRCW, func(o *models.Order) interface{} {
		if sh := firstShipment(o); sh != nil && sh.Wilayah != nil {
			return sh.Wilayah.Provinsi + " - " + sh.Wilayah.KotaKab
		}
		return ""
	}},
	{"Berat (Kg)", katTimbanganRCW, func(o *models.Order) interface{} {
		if sh := firstShipment(o); sh != nil {
			return sh.Berat
		}
		return ""
	}},

	{"Harga Produk Reseller", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaProdukReseller })},
	{"Harga PPN 11% Reseller", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaPPNReseller })},
	{"Harga Ongkir Reseller", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaOngkirReseller })},
	{"Harga PPN 11% Ongkir Reseller", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.HargaPPNOngkirReseller })},
	{"Total Harga + Ongkir Reseller (Exclude PPN)", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.TotalHargaOngkirReseller })},
	{"Total Harga Reseller (Include PPN)", katTimbanganRCW, pricingVal(func(p *models.OrderPricing) float64 { return p.TotalHargaReseller })},

	{"Nama Ekspedisi", katTimbanganRCW, func(o *models.Order) interface{} {
		if sh := firstShipment(o); sh != nil && sh.Ekspedisi != nil {
			return sh.Ekspedisi.NamaEkspedisi
		}
		return ""
	}},
	{"Resi", katTimbanganRCW, func(o *models.Order) interface{} {
		if sh := firstShipment(o); sh != nil {
			return strVal(sh.Resi)
		}
		return ""
	}},
	{"Tanggal Barang Diterima", katTimbanganRCW, func(o *models.Order) interface{} {
		if sh := firstShipment(o); sh != nil {
			return dateVal(sh.TanggalDiterima)
		}
		return ""
	}},
}

func pricingVal(get func(p *models.OrderPricing) float64) func(o *models.Order) interface{} {
	return func(o *models.Order) interface{} {
		if o.Pricing == nil {
			return ""
		}
		return get(o.Pricing)
	}
}

func activeColumns(kategoriFilter []string) []exportColumn {
	active := kategoriFilter
	if len(active) == 0 {
		active = rules.AllKategori
	}

	activeSet := make(map[string]bool, len(active))
	for _, k := range active {
		activeSet[k] = true
	}

	var result []exportColumn
	for _, col := range exportColumns {
		for _, k := range col.Kategori {
			if activeSet[k] {
				result = append(result, col)
				break
			}
		}
	}
	return result
}

func (s *OrderExportService) GenerateExcel(
	filter models.OrderExportFilter,
) ([]byte, error) {
	orders, err := s.Repository.FindForExport(filter)
	if err != nil {
		return nil, err
	}

	columns := activeColumns(filter.KategoriOrder)

	f := excelize.NewFile()
	defer f.Close()

	const sheet = "Data Pesanan"
	f.SetSheetName(f.GetSheetName(0), sheet)

	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "FFFFFF"},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"305496"}, Pattern: 1},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	})
	if err != nil {
		return nil, err
	}

	for i, col := range columns {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, col.Header)

		colName, _ := excelize.ColumnNumberToName(i + 1)
		width := 20.0
		if col.Header == "Nama Instansi/Perusahaan" {
			width = 30
		}
		f.SetColWidth(sheet, colName, colName, width)
	}

	lastHeaderCell, _ := excelize.CoordinatesToCellName(len(columns), 1)
	f.SetCellStyle(sheet, "A1", lastHeaderCell, headerStyle)
	f.SetRowHeight(sheet, 1, 32)

	for r, order := range orders {
		o := order
		for c, col := range columns {
			value := col.Value(&o)
			if s, ok := value.(string); ok && s == "" {
				continue
			}
			cell, _ := excelize.CoordinatesToCellName(c+1, r+2)
			f.SetCellValue(sheet, cell, value)
		}
	}

	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("gagal membuat file excel: %w", err)
	}

	return buf.Bytes(), nil
}