package review

import (
	"github.com/stretchr/testify/assert"
	"product_rating/database"
	"testing"
)

func TestCalculateSummary(t *testing.T) {
	data := []database.ProductReviews{
		{
			ID:        1,
			ProductID: 2,
			Rating:    5,
		},
		{
			ID:        2,
			ProductID: 1,
			Rating:    3,
		},
		{
			ID:        3,
			ProductID: 3,
			Rating:    4,
		},
		{
			ID:        4,
			ProductID: 4,
			Rating:    4,
		},
		{
			ID:        5,
			ProductID: 4,
			Rating:    1,
		},
		{
			ID:        6,
			ProductID: 3,
			Rating:    1,
		},
		{
			ID:        7,
			ProductID: 4,
			Rating:    3,
		},
	}

	result := CalculateSummary(&data)

	assert.EqualValues(t, result, &DefaultResponse{
		TotalReviews:   7,
		AverageRatings: 3,
		FiveStar:       1,
		FourStar:       2,
		ThreeStar:      2,
		TwoStar:        0,
		OneStar:        2,
	})

	// reset cache
	summaryReturnData = DefaultResponse{}

	var emptyData []database.ProductReviews

	assert.EqualValues(t, CalculateSummary(&emptyData), &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	})
}
