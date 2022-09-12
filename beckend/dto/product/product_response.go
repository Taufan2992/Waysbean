package productdto

type ProductResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
