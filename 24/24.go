package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

// NewPoint Конструктор
func NewPoint(newX, newY int) *Point {
	return &Point{x: newX, y: newY}
}

// ComputeDistance Поиск дистанции
func (p *Point) ComputeDistance(p2 Point) float64 {
	xDelta, yDelta := p.x-p2.x, p.y-p2.y
	xSq, ySq := xDelta*xDelta, yDelta*yDelta
	return math.Sqrt(float64(xSq + ySq))
}

func main() {
	p1, p2 := NewPoint(1, 2), NewPoint(20, -4)
	distance := p1.ComputeDistance(*p2)
	fmt.Printf("Расстояние между точками = %0.3f ед", distance)
}
