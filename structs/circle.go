package structs

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}
