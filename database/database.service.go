package database

import (
	"encoding/json"
	"os"
)

var productReviews []ProductReviews

func GetProductReviewData(filePath string) *[]ProductReviews {
	if len(productReviews) != 0 {
		return &productReviews
	}
	reviewFilesBytes, _ := os.ReadFile(filePath)
	json.Unmarshal(reviewFilesBytes, &productReviews)
	return &productReviews
}
