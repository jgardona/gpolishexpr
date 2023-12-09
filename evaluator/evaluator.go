package evaluator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/jgardona/polishexpr/utils"
)

type PollishEvaluator struct {
	expression string
	variables  []float64
}

func NewPolishEvaluator(expr string, vars []float64) PollishEvaluator {
	return PollishEvaluator{expression: expr, variables: vars}
}

func (pe PollishEvaluator) Evaluate() (float64, error) {
	tokens := strings.Split(pe.expression, " ")
	arguments := []float64{}

	for _, token := range tokens {
		if unicode.IsDigit(rune(token[0])) {
			if result, err := strconv.ParseFloat(token, 64); err != nil {
				return 0.0, err
			} else {
				arguments = utils.Push(arguments, result)
			}
		} else if strnum, ok := strings.CutPrefix(token, "$"); ok {
			if index, err := strconv.ParseInt(strnum, 10, 64); err != nil {
				return 0.0, err
			} else {
				arguments = utils.Push(arguments, pe.variables[index])
			}
		} else {
			var value float64
			value, arguments = utils.Pop(arguments)
			switch token {
			case "+":
				var old_value float64
				old_value, arguments = utils.Pop(arguments)
				arguments = utils.Push(arguments, old_value+value)
			case "-":
				var old_value float64
				old_value, arguments = utils.Pop(arguments)
				arguments = utils.Push(arguments, old_value-value)
			case "*":
				var old_value float64
				old_value, arguments = utils.Pop(arguments)
				arguments = utils.Push(arguments, old_value*value)
			case "/":
				var old_value float64
				old_value, arguments = utils.Pop(arguments)
				if value == 0 {
					return 0.0, errors.New("cant divide by zero")
				}
				arguments = utils.Push(arguments, old_value/value)
			default:
				return 0.0, errors.New("function not implemented")
			}
		}
	}

	if len(arguments) != 1 {
		return 0.0, errors.New("the solution must have only one answere")
	} else {
		return arguments[0], nil
	}
}
