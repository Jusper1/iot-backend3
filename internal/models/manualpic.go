package models

type ManualPICRequest struct {
	NamaPIC string  `json:"nama_pic"`
	NoHP    *string `json:"no_hp"`
	NIK     *string `json:"nik"`
	Email   *string `json:"email"`
}