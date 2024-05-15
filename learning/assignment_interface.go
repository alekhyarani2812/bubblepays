package learning

import "github.com/bubblepays/log"

type Triangle struct {
	height float64
	base   float64
}
type Square struct {
	sideLength float64
}
type Shape interface {
	getArea() float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t Triangle) getArea() float64 {
	return (t.height * t.base) / 2
}
func PrintArea(s Shape) {
	log.Infof("area is : %v ", s.getArea())
}
func ShapeImpls() {
	tri := Triangle{height: 10, base: 10}
	sqa := Square{sideLength: 10}
	PrintArea(tri)
	PrintArea(sqa)
}
