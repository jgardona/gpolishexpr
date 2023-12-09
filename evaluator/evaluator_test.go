package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolishEvaluator(t *testing.T) {
	t.Run("must do 3 + 2 + 1.5 equals 6.5", func(t *testing.T) {
		ev := NewPolishEvaluator("3 2 + $0 +", []float64{1.5})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.Equal(t, 6.5, result)
	})

	t.Run("must do 3 - 2 equals 1", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $0 -", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.Equal(t, 1.0, result)
	})

	t.Run("subtraction must fail with invalid syntax error", func(t *testing.T) {
		ev := NewPolishEvaluator("3$0 -", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "strconv.ParseFloat: parsing \"3$0\": invalid syntax")
	})

	t.Run("addition must fail with cant parse integer invalid syntax", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $a +", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "strconv.ParseInt: parsing \"a\": invalid syntax")
	})

	t.Run("addition must fail with function not implemented", func(t *testing.T) {
		ev := NewPolishEvaluator("3 $0 ]", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Equal(t, 0.0, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "function not implemented")
	})

	t.Run("division results of 10 and 2 must be equal 5", func(t *testing.T) {
		ev := NewPolishEvaluator("10 $0 /", []float64{2.0})
		result, err := ev.Evaluate()
		assert.Nil(t, err)
		assert.Equal(t, 5.0, result)
	})

	t.Run("test cant divide by zero", func(t *testing.T) {
		ev := NewPolishEvaluator("10 $0 /", []float64{0.0})
		result, err := ev.Evaluate()
		assert.NotNil(t, err)
		assert.Equal(t, "cant divide by zero", err.Error())
		assert.Equal(t, 0.0, result)
	})
}
