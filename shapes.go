package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{1.2}
	tri := triangle{1.2, 3.4}
	printArea(sq)
	printArea(tri)
}

func (s square) getArea() float64 {
	return s.side * s.side
}
func (t triangle) getArea() float64 {
	return t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
