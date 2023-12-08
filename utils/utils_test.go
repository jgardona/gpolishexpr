package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	value, data := Pop(data)
	assert.Equal(t, []float64{0, 1, 2}, data)
	assert.Equal(t, 3.0, value)
}

func TestPush(t *testing.T) {
	data := []float64{0, 1, 2, 3}
	data = Push(data, 4)
	assert.Equal(t, []float64{0, 1, 2, 3, 4}, data)
}
