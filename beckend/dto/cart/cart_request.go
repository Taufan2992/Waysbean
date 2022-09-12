package cartdto

type CreateCart struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	Subamount    *int   `json:"subamount"`
	Status       string `json:"status"`
	Qty          *int   `json:"qty"`
	Stockproduct *int   `json:"stockproduct"`
}

type UpdateQtyRequest struct {
	Qty          int `json:"qty"`
	Subamount    int `json:"subamount"`
	Stockproduct int `json:"stockproduct"`
}

type UpdateCartRequest struct {
	TransactionID int `json:"transaction_id"`
}
