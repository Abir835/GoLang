package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Abir struct {
	Name string
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (a Abir) Abs() string {
	return a.Name + "a"
}

func main() {
	v := Vertex{3, 4}
	x := Abir{"x"}
	fmt.Println(x.Abs())
	fmt.Println(v.Abs())
}
