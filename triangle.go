package visvalingam

import "math"

type triangle struct {
	a, b, c *areaPoint
	index   int
	next    *triangle
	prev    *triangle
}

func newTriangle(a, b, c *areaPoint) *triangle {
	return &triangle{a, b, c, -1, nil, nil}
}

func area(a, b, c *areaPoint) float64 {
	return math.Abs((a.X-c.X)*(b.Y-a.Y) - (a.X-b.X)*(c.Y-a.Y))
}
