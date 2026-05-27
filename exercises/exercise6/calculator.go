package calculator

import "errors"

//1. Add testify to exercise6:
//
//# In exercise6 directory
//go get github.com/stretchr/testify
//
//2. Convert Calculator tests to use testify:
//
//Replace if/error checks with assert/require functions
//Create a test suite for calculator
//Add setup method that creates calculator instance
//
//3. Create a mock logger:
//
//Add Logger interface with Log(operation string) method
//Create mock using testify/mock
//Verify calculator logs operations

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
