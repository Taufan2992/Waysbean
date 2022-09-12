package models

import "time"

// User model struct
type Cart struct {
	ID            int               `json:"id" gorm:"PRIMARY_KEY"`
	UserID        int               `json:"user_id"`
	User          UserBuyerResponse `json:"user"`
	ProductID     int               `json:"product_id"`
	Product       ProductResponse   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty           *int              `json:"qty" form:"qty"`
	Subamount     *int              `json:"subamount"`
	TransactionID *int              `json:"transaction_id"`
	Stockproduct  *int              `json:"stockproduct"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"-"`
}

type CartResponse struct {
	ID        int             `json:"id"`
	UserID    int             `json:"user_id"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Qty       int             `json:"qty" form:"qty"`
	Subamount int             `json:"subamount"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

func (CartResponse) TableName() string {
	return "carts"
}
