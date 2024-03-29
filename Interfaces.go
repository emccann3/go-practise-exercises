package main

import (
	"fmt"
	"math"
)

type geometry interface {
	thisarea() float64
	thisperim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) thisarea() float64 {
	return r.width * r.height
}
func (r rect) thisperim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) thisarea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) thisperim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.thisarea())
	fmt.Println(g.thisperim())
}
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
