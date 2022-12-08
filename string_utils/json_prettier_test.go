package string_utils

import (
	"github.com/stretchr/testify/assert"
	"product_rating/review"
	"testing"
)

func TestToJSONPrettier(t *testing.T) {
	output := review.DefaultResponse{
		TotalReviews:   1,
		AverageRatings: 3,
		FiveStar:       0,
		FourStar:       0,
		ThreeStar:      1,
		TwoStar:        0,
		OneStar:        0,
	}

	result := ToJSONPrettier(output)

	assert.EqualValues(t, result, "{\n    \"total_reviews\": 1,\n    \"average_ratings\": 3,\n    \"5_star\": 0,\n    \"4_star\": 0,\n    \"3_star\": 1,\n    \"2_star\": 0,\n    \"1_star\": 0\n}")

}
