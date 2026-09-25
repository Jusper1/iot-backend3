package rules

import "errors"

const (
	JenisInaproc = "inaproc"
	JenisManual  = "manual"

	KategoriIoTInaproc       = "iot_inaproc"
	KategoriIoTManual        = "iot_manual"
	KategoriTimbanganInaproc = "timbangan_inaproc"
	KategoriTimbanganManual  = "timbangan_manual"
	KategoriRCW360           = "rcw_360"
	KategoriRCW800W          = "rcw_800w"
)

var AllKategori = []string{
	KategoriIoTInaproc,
	KategoriIoTManual,
	KategoriTimbanganInaproc,
	KategoriTimbanganManual,
	KategoriRCW360,
	KategoriRCW800W,
}

var kategoriTimbanganRCW = map[string]bool{
	KategoriTimbanganInaproc: true,
	KategoriTimbanganManual:  true,
	KategoriRCW360:           true,
	KategoriRCW800W:          true,
}

func IsKategoriValid(kategori string) bool {
	for _, k := range AllKategori {
		if k == kategori {
			return true
		}
	}
	return false
}

func IsKategoriTimbanganRCW(kategori string) bool {
	return kategoriTimbanganRCW[kategori]
}

var (
	ErrKategoriTidakDikenal = errors.New("kategori_order tidak dikenal")
	ErrEntityTidakBerlaku   = errors.New("entitas ini tidak berlaku untuk kategori order ini")
)

type ChildEntity string

const (
	EntityOrderManual ChildEntity = "order_manual"

	EntityOrderPricing ChildEntity = "order_pricing"

	EntityShipment ChildEntity = "shipments"
)

func IsChildEntityAllowed(kategori string, entity ChildEntity) (bool, error) {
	if !IsKategoriValid(kategori) {
		return false, ErrKategoriTidakDikenal
	}
	switch entity {
	case EntityOrderManual:
		return kategori == KategoriIoTManual, nil
	case EntityOrderPricing, EntityShipment:
		return kategoriTimbanganRCW[kategori], nil
	default:
		return true, nil
	}
}

func ValidateChildEntity(kategori string, entity ChildEntity) error {
	allowed, err := IsChildEntityAllowed(kategori, entity)
	if err != nil {
		return err
	}
	if !allowed {
		return ErrEntityTidakBerlaku
	}
	return nil
}

var procurementFieldsTimbanganOnly = map[string]bool{
	"no_po_kut":                             true,
	"nomor_surat_penyampaian_daftar_harga": true,
	"nomor_formulir_pembelian":              true,
}

func IsProcurementFieldAllowed(kategori string, column string) bool {
	if procurementFieldsTimbanganOnly[column] {
		return kategoriTimbanganRCW[kategori]
	}
	return true
}

var (
	orderFieldsLanggananOnly = map[string]bool{
		"periode_langganan": true,
	}

	orderFieldsExceptIotManual = map[string]bool{
		"kode_bayar":         true,
		"no_invoice_inaproc": true,
	}
)

func IsOrderFieldAllowed(kategori string, column string) bool {
	if orderFieldsLanggananOnly[column] {
		return kategori == KategoriIoTInaproc || kategori == KategoriIoTManual
	}
	if orderFieldsExceptIotManual[column] {
		return kategori != KategoriIoTManual
	}
	return true
}
