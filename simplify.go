package visvalingam

import (
	"container/heap"
)

func Simplify(points []Point, threshold float64) []Point {
	areaPoints := make([]*areaPoint, len(points))
	for i, p := range points {
		areaPoints[i] = newAreaPoint(p)
	}

	var q queue
	triangles := make([]*triangle, 0, len(areaPoints)-2)

	for i := 1; i < len(areaPoints)-1; i++ {
		t := newTriangle(areaPoints[i-1], areaPoints[i], areaPoints[i+1])
		t.b.area = area(t.a, t.b, t.c)
		if t.b.area > 0 {
			triangles = append(triangles, t)
			heap.Push(&q, t)
		}
	}

	for i, t := range triangles {
		if i > 0 {
			t.prev = triangles[i-1]
		}
		if i < len(triangles)-1 {
			t.next = triangles[i+1]
		}
	}

	var maxArea float64

	for len(q) > 0 {
		t := heap.Pop(&q).(*triangle)

		if t.b.area < maxArea {
			t.b.area = maxArea
		} else {
			maxArea = t.b.area
		}

		if t.prev != nil {
			t.prev.next = t.next
			t.prev.c = t.c
			t.prev.b.area = area(t.prev.a, t.prev.b, t.prev.c)
			heap.Fix(&q, t.prev.index)
		} else {
			t.a.area = t.b.area
		}

		if t.next != nil {
			t.next.prev = t.prev
			t.next.a = t.a
			t.next.b.area = area(t.next.a, t.next.b, t.next.c)
			heap.Fix(&q, t.next.index)
		} else {
			t.c.area = t.b.area
		}
	}

	var result []Point
	for _, p := range areaPoints {
		if p.area >= threshold {
			result = append(result, p.Point)
		}
	}
	return result
}
