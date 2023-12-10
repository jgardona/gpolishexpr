# polishexpr

A poorly implemented polish reverse expression evaluator.

The list of supported functions are:

**Arithmetic functions**: +, -, *, /

* **sin**: sine

* **cos**: cosine

* **ln**: natural logarithm

* **exp**: exponent

* **sqrt**: square root

Arguments for these functions could be as usual constants, written as numbers, as variables, written as $var_number ($0, for example). The variable number is zero based index of variables vector.

## Usage

* **Sums 2+2 and equals 4**
  
```go
ev := NewPolishEvaluator("$0 $1 +", []float64{2.0, 2.0})
result, err := ev.Evaluate()
assert.Equal(t, 4.0, result)
assert.NotNil(t, err)
assert.Equal(t, err, ErrBadFunction)
```
* **Evaluate 2 + 3 - 5**

```go
ev := NewPolishEvaluator("$0 $1 + 5 -", []float64{2.0, 3.0})
result, err := ev.Evaluate()
assert.Equal(t, 0.0, result)
assert.NotNil(t, err)
assert.Equal(t, err, ErrBadFunction)
```
