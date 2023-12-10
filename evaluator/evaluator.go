package evaluator

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/jgardona/polishexpr/utils"
)

var (
	ErrWrongSolution  = errors.New("the solution must have only one answere")
	ErrBadFunction    = errors.New("function not implemented")
	ErrDivisionByZero = errors.New("cant divide by zero")
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
	arguments := make([]float64, 0, 3)

	for _, token := range tokens {
		if unicode.IsDigit(rune(token[0])) {
			if result, err := strconv.ParseFloat(token, 64); err != nil {
				return 0.0, err
			} else {
				utils.Push(&arguments, result)
			}
		} else if strnum, ok := strings.CutPrefix(token, "$"); ok {
			if index, err := strconv.ParseInt(strnum, 10, 64); err != nil {
				return 0.0, err
			} else {
				utils.Push(&arguments, pe.variables[index])
			}
		} else {

			value := utils.Pop(&arguments)
			switch token {
			case "+":
				utils.Push(&arguments, utils.Pop(&arguments)+value)
			case "-":
				utils.Push(&arguments, utils.Pop(&arguments)-value)
			case "*":
				utils.Push(&arguments, utils.Pop(&arguments)*value)
			case "/":
				if value == 0 {
					return 0.0, ErrDivisionByZero
				}
				utils.Push(&arguments, utils.Pop(&arguments)/value)
			case "sin":
				utils.Push(&arguments, math.Sin(value))
			case "cos":
				utils.Push(&arguments, math.Cos(value))
			case "ln":
				utils.Push(&arguments, math.Log(value))
			case "exp":
				utils.Push(&arguments, math.Exp(value))
			case "sqrt":
				utils.Push(&arguments, math.Sqrt(value))

			default:
				return 0.0, ErrBadFunction
			}
		}
	}

	if len(arguments) != 1 {
		return 0.0, ErrWrongSolution
	} else {
		return arguments[0], nil
	}
}
