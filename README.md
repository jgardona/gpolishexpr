# polishexpr

A poorly implemented polish reverse expression evaluator.

The list of supported functions are:

**Arithmetic functions**: +, -, *, /
**sin**: sine
**cos**: cosine
**ln**: natural logarithm
**exp**: exponent
**sqrt**: square root
Arguments for these functions could be as usual constants, written as numbers, as variables, written as $var_number ($0, for example). The variable number is zero based index of variables vector.

## Usage

* **Sums 2+2 and equals 4**
```go
let pe = PolishEvaluator::new("$0 2 +", &[2.0]);
let result = pe.evaluate()?;
assert_eq!(4f64, result);
```
* **Evaluate 2 + 3 - 5**

```go
let pe = PolishEvaluator::new("$0 $1 + $3 -", &[2.0, 3.0, 5.0]);
let result = pe.evaluate()?;
assert_eq!(0f64, result);
```
