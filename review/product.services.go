package review

import (
	"product_rating/database"
	"product_rating/float_utils"
)

var productReturnData = make(map[int]*DefaultResponse)

func CalculateProduct(reviewsData *[]database.ProductReviews, productId int) *DefaultResponse {
	if _, ok := productReturnData[productId]; ok {
		return productReturnData[productId]
	}
	ratingTotal := 0
	totalReviews := 0
	productReturnData[productId] = &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	}
	for _, reviewData := range *reviewsData {
		if reviewData.ProductID != productId {
			continue
		}
		ratingTotal += reviewData.Rating
		totalReviews++
		IncrementStarInReturnData(productReturnData[productId], reviewData.Rating)
	}
	if totalReviews > 0 && ratingTotal > 0 {
		productReturnData[productId].TotalReviews = totalReviews
		productReturnData[productId].AverageRatings = float_utils.ToFixed(float64(ratingTotal)/float64(totalReviews), 1)
	}
	return productReturnData[productId]
}
