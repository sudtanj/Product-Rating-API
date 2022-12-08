package database

type ProductReviews struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Rating    int `json:"rating"`
}
