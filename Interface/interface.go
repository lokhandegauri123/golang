package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float32
}
type Circle struct{
	radius float32
}

type Rectangle struct{
	width ,  height float32
}

func (c Circle) area() float32{
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float32{
	return r.width * r.height
}

func main(){
	c := Circle{radius: 5.0}
	r := Rectangle{width: 6.0, height: 7.0}

	fmt.Println(c.area()) 
	fmt.Println(r.area())
}