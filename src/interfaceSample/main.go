// interfaceSample project main.go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	len, width float32
}

func (r *rect) area() float32 {
	return r.len * r.width
}

func (r *rect) perim() float32 {
	return 2 * (r.len + r.width)
}

type circle struct {
	radius float32
}

func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r *circle) perim() float32 {
	return 2 * math.Pi * r.radius
}

func show(name string, parmas interface{}) {
	switch parmas.(type) {
	case geometry:
		fmt.Printf("area of %v is %v \n", name, parmas.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, parmas.(geometry).perim())
	default:
		fmt.Println("wrong type!")
	}
}
func main() {
	r := &rect{
		len:   1,
		width: 2,
	}
	show("rect", r)

	c := &circle{
		radius: 1,
	}
	show("cricle", c)

	show("test", "")
}
