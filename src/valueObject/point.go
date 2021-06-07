package valueObject

type Point struct {
	x, y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}
