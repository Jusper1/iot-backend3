package models

import "time"

type MasterInstansi struct {
	ID uint `json:"id" gorm:"primaryKey"`

	NamaInstansi	string `json:"nama_instansi" gorm:"not null"`
	NPWP *			string `json:"npwp"`
	Provinsi 		*string `json:"provinsi"`
	KotaKab 		*string `json:"kota_kab"`
	Alamat 			*string `json:"alamat"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	PICs []MasterPIC `json:"pics,omitempty" gorm:"foreignKey:InstansiID"`
}

func (MasterInstansi) TableName() string {
	return "master_instansi"
}