package oop

import "testing"

func TestPerimeter(t *testing.T) {
    areaTests := []struct{
        shape Shape
        hasPerimeter float64
    } {
        {Rectangle{10.0, 10.0}, 40.0},
        {Circle{10.0}, 62.83185307179586},
    }

    for _, tt := range areaTests {
        got := tt.shape.Perimeter()

        if got != tt.hasPerimeter {
            t.Errorf("%T: %v has perimeter %g got %g", tt.shape, tt.shape, tt.hasPerimeter, got)
        }
    }
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 11, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
                t.Errorf("%#v has area %g got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
