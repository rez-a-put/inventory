package model

import (
	"time"
)

// User : struct to hold user data. Based for migration into table users
type User struct {
	ID        int64      `json:"id" gorm:"type:bigint;primary_key;auto_increment"`
	Name      string     `json:"name" gorm:"type:varchar(100)" binding:"required"`
	Email     string     `json:"email" gorm:"type:varchar(50)" binding:"required,email"`
	Password  string     `json:"password" gorm:"type:varchar(255)" binding:"required"`
	Phone     string     `json:"phone" gorm:"type:varchar(15)"`
	Address   string     `json:"address" gorm:"type:text"`
	Status    int8       `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamp"`
}

// Login : struct used when login
type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
