package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{
		sideLength: 10,
	}

	t := triangle{
		base:   10,
		height: 25,
	}

	printArea(sq)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("Area = ", s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
