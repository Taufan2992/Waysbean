package models

import "time"

// User model struct
type Product struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price" form:"price" gorm:"type: int"`
	Image       string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func (ProductResponse) TableName() string {
	return "products"
}
