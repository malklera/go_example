package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("interface:", g)
	fmt.Println(".area:", g.area())
	fmt.Println(".perim:", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func detectRect(g geometry) {
	//is the same as detectCircle, but for rect
	// interface.(struct) it can be unpacket to two variables, the actual struct
	// and a boolean if it is of that type, this is callled type assertion
	r, ok := g.(rect)
	if ok {
		fmt.Println("rect of width:", r.width)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

	detectRect(r)
	detectRect(c)
}
