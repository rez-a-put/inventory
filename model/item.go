package model

import (
	"time"
)

// Item : struct to hold item data. Based for migration into table items
type Item struct {
	ID          int64      `json:"id" gorm:"type:bigint;primary_key;auto_increment"`
	SKU         string     `json:"sku" gorm:"type:varchar(25)" binding:"required"`
	Name        string     `json:"name" gorm:"type:varchar(100)" binding:"required"`
	Description string     `json:"description" gorm:"type:text"`
	UnitPrice   float64    `json:"unit_price" gorm:"type:decimal(10,2)" binding:"required"`
	Status      int8       `json:"status" gorm:"type:tinyint;default:1;comment:1=Active,2=Inactive"`
	CreatedAt   time.Time  `json:"created_at,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP()"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamp"`
}
