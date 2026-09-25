package models

type OrderExportFilter struct {
	KodeOrder    string 
	NamaInstansi string 
	Status       string 

	TanggalPODari   string 
	TanggalPOSampai string

	TanggalUangMasukDari   string 
	TanggalUangMasukSampai string

	Provinsi string 
	KategoriOrder []string
}