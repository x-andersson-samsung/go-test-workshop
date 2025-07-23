package calculator

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableCalculator_Add(t *testing.T) {
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

func TestTableCalculator_Sub(t *testing.T) {
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

func TestTableCalculator_Div(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		cases := []struct {
			a, b int
			want int
		}{
			{0, 1, 0},
			{1, 1, 1},
			{1, -1, -1},
			{-1, 1, -1},
			{-1, -1, 1},
			{2, 2, 1},
			{3, 2, 1},
			{4, 2, 2},
		}
		for _, c := range cases {
			t.Run(fmt.Sprintf("%d / %d", c.a, c.b), func(t *testing.T) {
				got, gotErr := (Calculator{}).Div(c.a, c.b)

				require.NoError(t, gotErr)
				require.Equal(t, got, c.want)
			})
		}
	})

	// Remember to check error paths
	t.Run("error_div_by_zero", func(t *testing.T) {
		_, gotErr := (Calculator{}).Div(1, 0)
		require.ErrorIs(t, gotErr, ErrDivByZero)
	})
}
