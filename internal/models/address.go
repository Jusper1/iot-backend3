package models

type AddressOverride struct {
	Alamat   *string `json:"alamat"`
	KotaKab  *string `json:"kota_kab"`
	Provinsi *string `json:"provinsi"`
}
type PicOverride struct {
	NamaPIC *string `json:"nama_pic_override"`
	NoHP    *string `json:"no_hp_override"`
}