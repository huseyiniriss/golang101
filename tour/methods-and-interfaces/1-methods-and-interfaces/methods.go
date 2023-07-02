package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// methods are functions
	fmt.Println(Abs(v))

	// methods continued
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// pointer receivers
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v.Abs())

	// pointer receivers
	v3 := Vertex{3, 4}
	Scale(&v3, 10)
	fmt.Println(Abs(v3))

	// pointers and functions
	v4 := Vertex{3, 4}
	Scale(&v4, 10)
	fmt.Println(Abs(v4))

	// methods and pointer indirection
	v5 := Vertex{3, 4}
	v5.Scale(2)
	ScaleFunc(&v5, 10)

	v6 := &Vertex{4, 3}
	v6.Scale(3)
	ScaleFunc(v6, 8)

	fmt.Println(v, v6)

	// methods and pointer indirection (2)
	v7 := Vertex{3, 4}
	fmt.Println(v7.Abs())
	fmt.Println(AbsFunc(v7))

	v8 := &Vertex{4, 3}
	fmt.Println(v8.Abs())
	fmt.Println(AbsFunc(*v8))

	// choosing a value or pointer receiver
	v9 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v9, v9.Abs())
	v9.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v9, v9.Abs())
}
