package review

import (
	"github.com/stretchr/testify/assert"
	"product_rating/database"
	"testing"
)

func TestCalculateProduct(t *testing.T) {
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

	assert.EqualValues(t, CalculateProduct(&data, 1), &DefaultResponse{
		TotalReviews:   1,
		AverageRatings: 3,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      1,
		TwoStar:        0,
		OneStar:        0,
	})

	assert.EqualValues(t, CalculateProduct(&data, 3), &DefaultResponse{
		TotalReviews:   2,
		AverageRatings: 2.5,
		FiveStar:       0,
		FourStar:       1,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        1,
	})

	assert.EqualValues(t, CalculateProduct(&data, 6), &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	})

	assert.EqualValues(t, CalculateProduct(&data, -5), &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	})
}
