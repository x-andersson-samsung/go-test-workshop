package calculator

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	out := a
	for range b {
		out++
	}
	return out
}
