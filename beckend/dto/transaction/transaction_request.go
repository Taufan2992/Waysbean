package transactiondto

type TransactionRequest struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type UpdateTransactionRequest struct {
	Status    string `json:"status" gorm:"type: varchar(255)"`
	UserID    int    `json:"buyer_id"`
	ProductID int    `json:"product_id"`
}
