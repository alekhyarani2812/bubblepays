package learning

import (
	"fmt"
	"math"
)

// Addition of 2 numbers & strings using interfaces
type Caliculations interface {
	addStrings() string
	add() int
}
type Addition struct {
	p, q       int
	str1, str2 string
}

func (stu Addition) addStrings() string {
	return stu.str1 + stu.str2
}
func (ad Addition) add() int {
	return ad.p + ad.q
}
func AddTwoVal(cal Caliculations) {
	fmt.Println("addition of 2 integers is : ", cal.add())
	fmt.Println("addition of 2 strings is : ", cal.addStrings())
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

// Shape interface Recatangle struct using interfaces
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
	fmt.Println("area is : ", s.area())
}
func ShapeImpl() {
	rect := Rectanle{length: 8, breadth: 9}
	ShapeCal(rect)
}

// Map example program
func MapFunc() {
	subjectMarks := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}
	// Access values using keys
	fmt.Println("Golang marks:", subjectMarks["Golang"])
	fmt.Println("Java marks:", subjectMarks["Java"])
	fmt.Println("Python marks:", subjectMarks["Python"])
}

// map interface combination
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
	fmt.Println("area of rectangle is : ", rect.area())
}

//Map inteface combination
type ShapeIn interface{
	Area() float32
}
type Circle struct{
	radius float32
}
type Rectangle struct{
	length, height float32
}
func (c Circle) Area() float32{
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) Area() float32{
	return r.length * r.height
}
func MapInterfaceEx(){
	shapes := make(map[string]ShapeIn)
	circle := Circle{radius: 4}
	//circle.
	rectangle := Rectangle{length: 5, height: 5}
	shapes["rectangle"] = rectangle
	shapes["circle"] = circle	
	for shapeName, shape := range shapes{
		area := shape.Area()
		fmt.Println("key, value", shapeName,area)
	}
}