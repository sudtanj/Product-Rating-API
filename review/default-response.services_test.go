package review

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementStarInReturnData(t *testing.T) {
	returnData := &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	}

	IncrementStarInReturnData(returnData, 6)
	assert.EqualValues(t, returnData, &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        0,
	})

	IncrementStarInReturnData(returnData, 1)
	assert.EqualValues(t, returnData, &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        1,
	})

	IncrementStarInReturnData(returnData, 5)
	assert.EqualValues(t, returnData, &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       1,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        1,
	})

	IncrementStarInReturnData(returnData, 1)
	assert.EqualValues(t, returnData, &DefaultResponse{
		TotalReviews:   0,
		AverageRatings: 0,
		FiveStar:       1,
		FourStar:       0,
		ThreeStar:      0,
		TwoStar:        0,
		OneStar:        2,
	})
}
