package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func EuclidianDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {
	pointOne := NewPoint(5.6, 6.4)
	pointTwo := NewPoint(2.3, 4.4)
	fmt.Println(EuclidianDistance(pointOne, pointTwo))
}
