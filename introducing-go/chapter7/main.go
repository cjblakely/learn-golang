package main

import (
	"math"
	"fmt"
)

func main() {
	//var rx1, ry1 float64 = 0, 0
	//var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5

	//fmt.Println(rectangeArea(rx1, ry1, rx2, ry2))
	//fmt.Println(circleArea(cx, cy, cr))

	c := Circle{0, 0, 5}
	//fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

//func rectangeArea(x1, y1, x2, y2 float64) float64 {
//	l := distance(x1, y1, x1, y2)
//	w := distance(x1, y1, x2, y1)
//	return l * w
//}

//func circleArea(x, y, r float64) float64 {
//	return math.Pi * r*r
//}
//func circleArea(c *Circle) float64 {
//	return math.Pi * c.r*c.r
//}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//function cannot contain two variadic parameters
//func totalArea(circles ...Circle, rectangles ...Rectangle) float64 {
//	var total float64
//	for _, c := range circles {
//		total += c.area()
//	}
//	for _, r := range rectangles {
//		total += r.area()
//	}
//	return total
//}

//no good if we add more shapes
//func totalArea(circles []Circle, rectangles []Rectangle) float64 {
//		var total float64
//		for _, c := range circles {
//			total += c.area()
//		}
//		for _, r := range rectangles {
//			total += r.area()
//		}
//		return total
//}

//interface
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

//interface as field
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
