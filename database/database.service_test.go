package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductReviewData(t *testing.T) {
	data := GetProductReviewData("./reviews-data-sample.json")

	expectedResult := []ProductReviews{
		{
			ID:        1,
			ProductID: 2,
			Rating:    5,
		},
	}

	assert.EqualValues(t, data, &expectedResult)
}
