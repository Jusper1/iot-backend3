package models

import "time"

type Document struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID 		*uint `json:"order_id"`
	SPJID 			*uint `json:"spj_id"`
	JenisDokumen 	string `json:"jenis_dokumen" gorm:"not null"`
	NamaFile 		string `json:"nama_file" gorm:"not null"`
	FileURL 		*string `json:"file_url"`
	FilePath 	*string `json:"file_path"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Document) TableName() string {
	return "documents"
}