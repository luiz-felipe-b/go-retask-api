package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}
