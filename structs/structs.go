package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Item         Item      `gorm:"foreignKey:OrderID;references:ID"`
}

type Item struct {
	gorm.Model
	ItemCode    int64  `json:"item_code"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
