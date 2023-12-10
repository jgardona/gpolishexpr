package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	value := Pop(&data)
	assert.Equal(t, []float64{0, 1, 2}, data)
	assert.Equal(t, 3.0, value)
}

func TestPush(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	Push(&data, 4)
	assert.Equal(t, []float64{0, 1, 2, 3, 4}, data)
}

func TestReturnEmptyValue(t *testing.T) {
	data := []float64{}
	value := Pop(&data)
	assert.Equal(t, 0.0, value)
}
