package models

import "time"

type Order struct {
	ID uint `json:"id" gorm:"primaryKey"`

	KodeOrder     string `json:"kode_order" gorm:"not null;uniqueIndex"`
	JenisOrder    string `json:"jenis_order" gorm:"not null"`
	KategoriOrder string `json:"kategori_order" gorm:"not null"`

	InstansiID *uint `json:"instansi_id"`
	PicID      *uint `json:"pic_id"`

	Status     *string `json:"status"`
	StatusOdoo *string `json:"status_odoo"`

	NSFP             *string `json:"nsfp"`
	NoBAST           *string `json:"no_bast"`
	KodeBayar        *string `json:"kode_bayar"`
	NoInvoiceInaproc *string `json:"no_invoice_inaproc"`
	Keterangan       *string `json:"keterangan"`

	TanggalPO        *Date   `json:"tanggal_po"`
	NoPO             *string `json:"no_po"`
	PeriodeLangganan *string `json:"periode_langganan"`
	TanggalBAST      *Date   `json:"tanggal_bast"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Instansi *MasterInstansi `json:"instansi,omitempty" gorm:"foreignKey:InstansiID"`
	PIC      *MasterPIC      `json:"pic,omitempty" gorm:"foreignKey:PicID"`

	Items       []OrderItem       `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	Pricing     *OrderPricing     `json:"pricing,omitempty" gorm:"foreignKey:OrderID"`
	Manual      *OrderManual      `json:"manual,omitempty" gorm:"foreignKey:OrderID"`
	Procurement *OrderProcurement `json:"procurement,omitempty" gorm:"foreignKey:OrderID"`

	Payments  []Payment       `json:"payments,omitempty" gorm:"foreignKey:OrderID"`
	Shipments []Shipment      `json:"shipments,omitempty" gorm:"foreignKey:OrderID"`
	Documents []OrderDocument `json:"order_documents,omitempty" gorm:"foreignKey:OrderID"`
	SPJ       []SPJ           `json:"spj,omitempty" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "orders"
}