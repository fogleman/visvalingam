package visvalingam

type Point struct {
	X, Y float64
}

type areaPoint struct {
	Point
	area float64
}

func newAreaPoint(p Point) *areaPoint {
	return &areaPoint{p, 0}
}
