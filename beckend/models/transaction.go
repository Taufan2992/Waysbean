package models

import "time"

type Transaction struct {
	ID        int                      `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                      `json:"user_id"`
	User      UsersTransactionResponse `json:"user"`
	Carts     []Cart                   `json:"carts"`
	Status    string                   `json:"status"  gorm:"type:varchar(25)"`
	Total     int                      `json:"total"`
	CreatedAt time.Time                `json:"-"`
	UpdatedAt time.Time                `json:"-"`
}

type TransactionResponse struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	Status string               `json:"status" gorm:"type: varchar(255)"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
