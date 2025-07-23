package calculator

import (
	"errors"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	c := Calculator{}
	if c.Add(1, 1) != 2 {
		t.Error("1 + 1 = 2")
	}
	if c.Add(0, 1) != 1 {
		t.Error("0 + 1 = 1")
	}
	if c.Add(1, 0) != 1 {
		t.Error("1 + 0 = 1")
	}
}

func TestCalculator_Sub(t *testing.T) {
	c := Calculator{}
	if c.Sub(1, 1) != 0 {
		t.Error("1 - 1 = 0")
	}
	if c.Sub(0, 1) != -1 {
		t.Error("0 - 1 = -1")
	}
	if c.Sub(1, 0) != 1 {
		t.Error("1 - 0 = 1")
	}
}

func TestCalculator_Div(t *testing.T) {
	c := Calculator{}
	if out, err := c.Div(1, 1); out != 1 || err != nil {
		t.Error("1 / 1 = 1")
	}
	if out, err := c.Div(0, 1); out != 0 || err != nil {
		t.Error("0 / 1 = 0")
	}
	// Remember to check error paths
	if out, err := c.Div(1, 0); out != 0 || !errors.Is(err, ErrDivByZero) {
		t.Error("expected 0 and ErrDivByZero")
	}
}
