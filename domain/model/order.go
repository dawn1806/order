package model

import "time"

type Order struct{
	ID int64 `gorm:"primary_key;not null;auto_increment" json:"id"`
	OrderCode string `gorm:"unique_index;not null" json:"order_code"`
	PayStatus int32 `json:"pay_status"`
	ShipStatus int32 `json:"ship_status"`
	Price float64 `json:"price"`
	OrderDetail []OrderDetail `gorm:"ForeignKey:OrderID" json:"order_detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

