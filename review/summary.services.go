package review

import (
	"product_rating/database"
	"product_rating/float_utils"
)

var summaryReturnData DefaultResponse

func CalculateSummary(reviewsData *[]database.ProductReviews) *DefaultResponse {
	if (DefaultResponse{}) != summaryReturnData {
		return &summaryReturnData
	}
	ratingTotal := 0
	totalReviews := len(*reviewsData)
	summaryReturnData = DefaultResponse{
		TotalReviews:   totalReviews,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	}
	for _, reviewData := range *reviewsData {
		ratingTotal += reviewData.Rating
		IncrementStarInReturnData(&summaryReturnData, reviewData.Rating)
	}
	if ratingTotal > 0 && totalReviews > 0 {
		summaryReturnData.AverageRatings = float_utils.ToFixed(float64(ratingTotal)/float64(totalReviews), 1)
	}
	return &summaryReturnData
}
