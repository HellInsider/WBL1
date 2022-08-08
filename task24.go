package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func New(x float64, y float64) Point {
	return Point{x, y}
}

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func task24() {
	fmt.Println("------------Task 24----------------")
	p1 := New(-4, 3)
	p2 := New(0, 0)
	fmt.Println("Distance : ", distance(p2, p1))
}
