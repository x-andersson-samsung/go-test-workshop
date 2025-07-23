package calculator

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, -1, 0},
		{math.MaxInt, 1, math.MinInt},
		{math.MaxInt, math.MinInt, -1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d + %d", c.a, c.b), func(t *testing.T) {
			got := (Calculator{}).Add(c.a, c.b)
			if got != c.want {
				t.Fatalf("calculator should return %d, got %d", c.want, got)
			}
		})
	}
}

func TestCalculator_Sub(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{0, 1, -1},
		{1, 0, 1},
		{1, -1, 2},
		{math.MaxInt, math.MinInt, -1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d - %d", c.a, c.b), func(t *testing.T) {
			got := (Calculator{}).Sub(c.a, c.b)
			if got != c.want {
				t.Fatalf("calculator should return %d, got %d", c.want, got)
			}
		})
	}
}

func TestCalculator_Div(t *testing.T) {
	cases := []struct {
		a, b    int
		want    int
		wantErr error
	}{
		{0, 1, 0, nil},
		{1, 1, 1, nil},
		{1, -1, -1, nil},
		{-1, 1, -1, nil},
		{-1, -1, 1, nil},
		{1, 0, 0, DivByZero},
		{2, 2, 1, nil},
		{3, 2, 1, nil},
		{4, 2, 2, nil},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d / %d", c.a, c.b), func(t *testing.T) {
			got, gotErr := (Calculator{}).Div(c.a, c.b)
			if got != c.want {
				t.Fatalf("calculator should return %d, got %d", c.want, got)
			}
			if !errors.Is(gotErr, c.wantErr) {
				t.Fatalf("calculator should return error `%v`, got `%v`", c.wantErr, gotErr)
			}
		})
	}
}
