package float_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRound(t *testing.T) {
	assert.Equal(t, Round(2.921), 3)
	assert.Equal(t, Round(2.2), 2)
}

func TestToFixed(t *testing.T) {
	assert.Equal(t, ToFixed(2.921, 1), 2.9)
	assert.Equal(t, ToFixed(2.951, 1), float64(3))
	assert.Equal(t, ToFixed(2.921, 2), 2.92)
	assert.Equal(t, ToFixed(2.921, 4), 2.9210)
}
