package evaluator

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/jgardona/polishexpr/utils"
)

const (
	sum            = "+"
	subtraction    = "-"
	multiplication = "*"
	division       = "/"
	sine           = "sin"
	cosine         = "cos"
	logarithm      = "ln"
	exponential    = "exp"
	squareroot     = "sqrt"
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
			case sum:
				utils.Push(&arguments, utils.Pop(&arguments)+value)
			case subtraction:
				utils.Push(&arguments, utils.Pop(&arguments)-value)
			case multiplication:
				utils.Push(&arguments, utils.Pop(&arguments)*value)
			case division:
				if value == 0 {
					return 0.0, ErrDivisionByZero
				}
				utils.Push(&arguments, utils.Pop(&arguments)/value)
			case sine:
				utils.Push(&arguments, math.Sin(value))
			case cosine:
				utils.Push(&arguments, math.Cos(value))
			case logarithm:
				utils.Push(&arguments, math.Log(value))
			case exponential:
				utils.Push(&arguments, math.Exp(value))
			case squareroot:
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
