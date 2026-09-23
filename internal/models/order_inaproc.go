package models

import "time"

type OrderInaproc struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID uint `json:"order_id" gorm:"not null"`

	NoInvoiceInaproc *string `json:"no_invoice_inaproc"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OrderInaproc) TableName() string {
	return "order_inaproc"
}