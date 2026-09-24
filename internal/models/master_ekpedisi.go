package models

import "time"

type MasterEkspedisi struct {
	ID uint `json:"id" gorm:"primaryKey"`

	NamaEkspedisi 	string `json:"nama_ekspedisi" gorm:"not null"`
	Status 			string `json:"status" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MasterEkspedisi) TableName() string {
	return "master_ekspedisi"
}