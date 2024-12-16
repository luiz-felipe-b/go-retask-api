package structs

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() (area float64) {
	area = t.Base * t.Height / 2
	return
}
