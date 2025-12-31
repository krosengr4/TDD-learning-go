package main

import "testing"

// Now that Shape is an interface with the Area() func, we can test all at the same time
func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 4, Height: 6}, hasArea: 24},
		{name: "Square", shape: Square{Height: 5}, hasArea: 25},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.shape.Area()
			exp := tt.hasArea

			if act != exp {
				t.Errorf("%#v: Got %g when expecting %g", tt.shape, act, exp)
			}
		})

	}
}


