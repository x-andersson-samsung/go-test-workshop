package calculator

import "testing"

func TestCalculator_Add(t *testing.T) {
	c := Calculator{}
	c.Add(1, 1)
	if c.Add(1, 1) != 2 {
		t.Error("calculator should return 2")
	}
}
