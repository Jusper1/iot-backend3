package models

import "time"

type OrderItem struct {
	ID uint `json:"id" gorm:"primaryKey"`

	OrderID  uint `json:"order_id" gorm:"not null"`
	ProdukID uint `json:"produk_id" gorm:"not null"`

	Qty      int64   `json:"qty" gorm:"not null"`
	Harga    float64 `json:"harga" gorm:"not null"`
	PPN      float64 `json:"ppn" gorm:"not null"`
	Subtotal float64 `json:"subtotal" gorm:"not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Produk *MasterProduk `json:"produk,omitempty" gorm:"foreignKey:ProdukID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}