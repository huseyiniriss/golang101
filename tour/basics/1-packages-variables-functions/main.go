package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// var statement declares a list of variables
var c, python, java bool

func main() {
	//hello world
	fmt.Println("Hello World")

	//exported names
	fmt.Println(math.Pi)

	add(42, 13)
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	split(17)

	varDeclaration()
	varDeclarationWithInitializers()
	shortVarDeclaration()

	basicTypes()
	zeroValues()
	typeConversions()
	typeInference()
	constants()
	numericConstants()
}

// func name (param type, param type) return type
func add(x, y int) int {
	return x + y
}

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var declaration in function
func varDeclaration() {
	var i int
	fmt.Println(i, c, python, java)
}

// variable declaration with initializers
func varDeclarationWithInitializers() {
	var i, j int = 1, 2
	fmt.Println(i, j)
}

// short variable declaration
func shortVarDeclaration() {
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(k, c, python, java)
}

// basic types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
//
//	// represents a Unicode code point
//
// float32 float64
// complex64 complex128
func basicTypes() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// zero values
// 0 0 false ""
func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// type conversions
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// or
// i := 42
// f := float64(i)
// u := uint(f)
func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// type inference
// v := 42 // change me!
// fmt.Printf("v is of type %T\n", v)
func typeInference() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

// Constants
// const Pi = 3.14
// const World = "世界"
// const Truth = true
func constants() {
	const Pi = 3.14
	fmt.Println("Hello", Pi)
}

// Numeric Constants
// const (
//
//	// Create a huge number by shifting a 1 bit left 100 places.
//	// In other words, the binary number that is 1 followed by 100 zeroes.
//	Big = 1 << 100
//	// Shift it right again 99 places, so we end up with 1<<1, or 2.
//	Small = Big >> 99
//
// )
func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

// fmt.Println(needInt(Small))
// fmt.Println(needFloat(Small))
// fmt.Println(needFloat(Big))
func numericConstants() {
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
