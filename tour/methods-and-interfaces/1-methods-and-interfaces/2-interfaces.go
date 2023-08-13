package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	// interfaces
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	// interfaces are implemented implicitly
	var i I = &T{"hello"}
	i.M()

	// interface values
	var i2 I
	i2 = &T{"Hello"}
	describe(i2)
	i2.M()

	i2 = F(math.Pi)
	describe(i2)
	i2.M()

	// interface values with nil underlying values
	var i3 I
	var t *T
	i3 = t
	describe(i3)
	i3.M()

	i3 = &T{"hello"}
	describe(i3)
	i3.M()

	// nil interface values
	var i4 I
	describe(i4)
	//i4.M()

	// empty interface
	var i5 interface{}
	describe2(i5)

	i5 = 42
	describe2(i5)

	i5 = "hello"
	describe2(i5)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
