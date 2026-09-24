package models

import "time"

type MasterPIC struct {
	ID uint `json:"id" gorm:"primaryKey"`

	InstansiID 	*uint `json:"instansi_id"`
	NamaPIC	 	string `json:"nama_pic" gorm:"not null"`
	Email 		*string `json:"email"`
	NIK 		*string `json:"nik"`
	NoHP 		*string `json:"no_hp"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Instansi *MasterInstansi `json:"instansi,omitempty" gorm:"foreignKey:InstansiID"`
}

func (MasterPIC) TableName() string {
	return "master_pic"
}