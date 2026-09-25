package rules


var JenisKertasOptions = []string{
	"A4",
	"F4",
}

var TipeTimbanganOptions = []string{
	"Digital 150 Kg",
}

var KodeBayarOptions = []string{
	"KB-001",
}

var StatusPesananOptions = []string{
	"pending",
	"baru",
	"diproses",
	"Dikirim",
	"selesai",
	"batal",
}

var StatusOdooOptions = []string{
	"draft",
	"confirmed",
}

var JenisFileOptions = []string{
	"PDF",
}


func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func IsValidJenisKertas(value string) bool    { return contains(JenisKertasOptions, value) }
func IsValidTipeTimbangan(value string) bool  { return contains(TipeTimbanganOptions, value) }
func IsValidKodeBayar(value string) bool      { return contains(KodeBayarOptions, value) }
func IsValidStatusPesanan(value string) bool  { return contains(StatusPesananOptions, value) }
func IsValidStatusOdoo(value string) bool     { return contains(StatusOdooOptions, value) }
func IsValidJenisFile(value string) bool      { return contains(JenisFileOptions, value) }