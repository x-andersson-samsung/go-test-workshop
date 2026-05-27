package calculator

//Add following functions in accordance to TDD principles:
//
//Sub(a, b float64) float64 - subtracts b from a
//Div(a, b float64) (float64, error) - divides a by b
//
//1. Write tests for all methods (they should fail)
//2. Implement the functions
//3. Run tests to verify implementation
//4. Refactor if needed

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	out := a
	for range b {
		out++
	}
	return out
}
