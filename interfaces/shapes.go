package main

import "math"

// Shape is anything that can tell us it's Area.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Height float64
}

func (s Square) Area() float64 {
	return s.Height * s.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
