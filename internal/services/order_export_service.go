package services

import (
	// "bytes"
	"fmt"

	"github.com/xuri/excelize/v2"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
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

var exportHeaders = []string{
	"Kode Pemesanan",
	"Kategori",
	"Nama Instansi",
	"Nama PIC",
	"No Telp PIC",
	"Email",
	"NIK",
	"No NPWP",
	"Alamat",
	"Kota/Kab",
	"Provinsi",
	"Quantity",
	"Harga + PPN",
	"Periode Berlangganan",
	"Status Pesanan",
	"Status Odoo",
	"Bulan Pengiriman SPH",
	"No. Surat Penawaran Harga",
	"No. Surat Penyampaian Daftar Harga",
	"No. Formulir Pembelian",
	"No. Formulir Berlangganan",
	"No. Kontrak Berlangganan",
	"No. PO KUT",
	"Tanggal Pesanan (PO)",
	"Tanggal BAST",
	"No. BAST",
	"No. Invoice KUT",
	"Tanggal Invoice KUT",
	"No. Invoice Inaproc",
	"NSFP",
	"Tanggal Uang Masuk",
	"Jumlah Uang Masuk",
	"Rekening Penerima",
	"Kode Bayar",
	"Keterangan",
	"Tipe Timbangan",
	"Harga Produk",
	"Harga PPN",
	"Harga Ongkir KUT",
	"Harga PPN 11% Ongkir",
	"Total Harga + Ongkir",
	"Total Harga Jual",
	"Wilayah Pengiriman",
	"Berat (Kg)",
	"Harga Produk Reseller",
	"Harga PPN Reseller",
	"Harga Ongkir Reseller",
	"Harga PPN Ongkir Reseller",
	"Total Harga + Ongkir Reseller",
	"Total Harga Reseller",
	"Nama Ekspedisi",
	"Resi",
	"Tanggal Barang Diterima",
}

func (s *OrderExportService) GenerateExcel(
	filter models.OrderExportFilter,
) ([]byte, error) {
	orders, err := s.Repository.FindForExport(filter)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	defer f.Close()

	const sheet = "IOT"
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

	for col, header := range exportHeaders {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheet, cell, header)
	}
	headerRange := fmt.Sprintf("A1:%s1", mustCellName(len(exportHeaders), 1))
	f.SetCellStyle(sheet, "A1", headerRange, headerStyle)
	f.SetRowHeight(sheet, 1, 30)

	for i, order := range orders {
		row := i + 2
		writeOrderRow(f, sheet, row, order)
	}

	for col := range exportHeaders {
		colName, _ := excelize.ColumnNumberToName(col + 1)
		f.SetColWidth(sheet, colName, colName, 20)
	}
	f.SetColWidth(sheet, "C", "C", 28) 
	if namaEkspedisiCol, err := excelize.ColumnNumberToName(51); err == nil {
		f.SetColWidth(sheet, namaEkspedisiCol, namaEkspedisiCol, 22) 
	}

	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func mustCellName(col, row int) string {
	name, _ := excelize.CoordinatesToCellName(col, row)
	return name
}


func writeOrderRow(f *excelize.File, sheet string, row int, order models.Order) {
	values := make([]interface{}, len(exportHeaders))

	values[0] = order.KodeOrder
	values[1] = order.KategoriOrder

	if order.Instansi != nil {
		values[2] = order.Instansi.NamaInstansi
		values[9] = strPtr(order.Instansi.KotaKab)
		values[10] = strPtr(order.Instansi.Provinsi)
		values[8] = strPtr(order.Instansi.Alamat)
		values[7] = strPtr(order.Instansi.NPWP)
	}

	if order.PIC != nil {
		values[3] = order.PIC.NamaPIC
		values[4] = strPtr(order.PIC.NoHP)
		values[5] = strPtr(order.PIC.Email)
		values[6] = strPtr(order.PIC.NIK)
	}

	var totalQty int64
	var totalHargaPPN float64
	for _, item := range order.Items {
		totalQty += item.Qty
		totalHargaPPN += item.Subtotal
	}
	if len(order.Items) > 0 {
		values[11] = totalQty
		values[12] = totalHargaPPN
	}

	values[13] = strPtr(order.PeriodeLangganan)
	values[14] = strPtr(order.Status)
	values[15] = strPtr(order.StatusOdoo)

	if order.Manual != nil {
		values[16] = strPtr(order.Manual.BulanPengirimanSPH)
		values[17] = strPtr(order.Manual.NomorSuratPenawaranHarga)
		values[20] = strPtr(order.Manual.NomorFormulirBerlangganan)
		values[21] = strPtr(order.Manual.NomorKontrakBerlangganan)
	}

	if order.Procurement != nil {
		values[18] = strPtr(order.Procurement.NomorSuratPenyampaianDaftarHarga)
		values[19] = strPtr(order.Procurement.NomorFormulirPembelian)
		values[22] = strPtr(order.Procurement.NoPOKUT)
		values[26] = strPtr(order.Procurement.NoInvoiceKUT)
		values[27] = datePtr(order.Procurement.TanggalInvoiceKUT)
	}

	values[23] = datePtr(order.TanggalPO)
	values[24] = datePtr(order.TanggalBAST)
	values[25] = strPtr(order.NoBAST)
	values[28] = strPtr(order.NoInvoiceInaproc)
	values[29] = strPtr(order.NSFP)

	if len(order.Payments) > 0 {
		p := order.Payments[0]
		values[30] = datePtr(p.TanggalUangMasuk)
		values[31] = p.JumlahUangMasuk
		values[32] = strPtr(p.Rekening)
	}

	values[33] = strPtr(order.KodeBayar)
	values[34] = strPtr(order.Keterangan)

	if order.Pricing != nil {
		p := order.Pricing
		values[35] = strPtr(p.TipeTimbangan)
		values[36] = p.HargaProduk
		values[37] = p.HargaPPN
		values[38] = p.HargaOngkirKUT
		values[39] = p.HargaPPNOngkir
		values[40] = p.TotalHargaOngkir
		values[41] = p.TotalHargaJual
		values[44] = p.HargaProdukReseller
		values[45] = p.HargaPPNReseller
		values[46] = p.HargaOngkirReseller
		values[47] = p.HargaPPNOngkirReseller
		values[48] = p.TotalHargaOngkirReseller
		values[49] = p.TotalHargaReseller
	}

	if len(order.Shipments) > 0 {
		sh := order.Shipments[0]
		if sh.Wilayah != nil {
			values[42] = sh.Wilayah.Provinsi + " - " + sh.Wilayah.KotaKab
		}
		values[43] = sh.Berat
		if sh.Ekspedisi != nil {
			values[50] = sh.Ekspedisi.NamaEkspedisi
		}
		values[51] = strPtr(sh.Resi)
		values[52] = datePtr(sh.TanggalDiterima)
	}

	for col, value := range values {
		if value == nil {
			continue
		}
		cell, _ := excelize.CoordinatesToCellName(col+1, row)
		f.SetCellValue(sheet, cell, value)
	}
}

func strPtr(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func datePtr(v *models.Date) string {
	if v == nil {
		return ""
	}
	return string(*v)
}