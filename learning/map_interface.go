package learning

import (
	"math"

	"github.com/bubblepays/log"
)

//Caliculations Addition of 2 numbers & strings using interfaces
type Caliculations interface {
	addStrings() string
	add() int
}
type Addition struct {
	p, q       int
	str1, str2 string
}

func (ad Addition) addStrings() string {
	return ad.str1 + ad.str2
}
func (ad Addition) add() int {
	return ad.p + ad.q
}
func AddTwoVal(cal Caliculations) {
	log.Infof("addition of 2 integers is : %v ", cal.add())
	log.Infof("addition of 2 strings is : %v ", cal.addStrings())
}
func Implimentation() {
	sum := Addition{
		p:    10,
		q:    20,
		str1: "alekhya",
		str2: "sheelam",
	}
	AddTwoVal(sum)
}

//Shapes interface Recatangle struct using interfaces
type Shapes interface {
	area() float32
}
type Rectanle struct {
	length, breadth float32
}

func (srtu Rectanle) area() float32 {
	return 2 * srtu.length * srtu.breadth
}
func ShapeCal(s Shapes) {
	log.Infof("area is : %v ", s.area())
}
func ShapeImpl() {
	rect := Rectanle{length: 8, breadth: 9}
	ShapeCal(rect)
}

//MapFunc Map example program
func MapFunc() {
	subjectMarks := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}
	// Access values using keys
	log.Infof("Golang marks: %v ", subjectMarks["Golang"])
	log.Infof("Java marks: %v ", subjectMarks["Java"])
	log.Infof("Python marks: %v ", subjectMarks["Python"])
}

//Shape1 map interface combination
type Shape1 interface {
	area() float32
}
type Rectanle1 struct {
	length, breadth float32
}

func (srtu Rectanle1) area() float32 {
	return 2 * srtu.length * srtu.breadth
}
func MapInterfaceImpl() {
	sMap := make(map[string]Shape1)
	sMap["rectangle"] = Rectanle1{length: 10, breadth: 20}
	rect := sMap["rectangle"]
	log.Infof("area of rectangle is : %v ", rect.area())
}

//ShapeIn Map inteface combination
type ShapeIn interface {
	Area() float32
}
type Circle struct {
	radius float32
}
type Rectangle struct {
	length, height float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) Area() float32 {
	return r.length * r.height
}
func MapInterfaceEx() {
	shapes := make(map[string]ShapeIn)
	circle := Circle{radius: 4}
	//circle.
	rectangle := Rectangle{length: 5, height: 5}
	shapes["rectangle"] = rectangle
	shapes["circle"] = circle
	for shapeName, shape := range shapes {
		area := shape.Area()
		log.Infof("key : %v , value : %v ", shapeName, area)
	}
}
