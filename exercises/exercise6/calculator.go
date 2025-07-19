package calculator

import "errors"

var (
	DivByZero = errors.New("division by zero")
)

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Sub(a, b int) int {
	return a - b
}

func (c Calculator) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, DivByZero
	}
	return a / b, nil
}
