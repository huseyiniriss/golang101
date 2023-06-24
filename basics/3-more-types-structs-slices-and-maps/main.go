package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	pointers()
	structs()
	structsFields()
	structsPointersToStructs()
	structsStructLiterals()
	arrays()
	slices()
	slicesAreLikeReferencesToArrays()
	slicesLiterals()
	slicesDefaults()
	slicesLengthAndCapacity()
	nilSlices()
	createASliceWithMake()
	slicesOfSlices()
	appendingToASlice()
	rangeExample()
	rangeContinued()
	maps()
	mapLiterals()
	mapLiteralsContinued()
	mutatingMaps()
	functionValues()
	functionClosures()
}

// pointers
func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// structs
func structs() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
}

// structs fields
func structsFields() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// structs pointers to structs
func structsPointersToStructs() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// structs struct literals
func structsStructLiterals() {
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
}

// arrays
func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a)          // [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]
}

// slices
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // [3 5 7]
	fmt.Println(s)
}

// slices are like references to arrays
func slicesAreLikeReferencesToArrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2] // [John Paul]
	b := names[1:3] // [Paul George]
	fmt.Println(names, a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names) // [John XXX George Ringo]
}

// slices literals
func slicesLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// slices defaults
func slicesDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3 5 7]
	fmt.Println(s)

	s = s[:2] // [3 5]
	fmt.Println(s)

	s = s[1:] // [5]
	fmt.Println(s)
}

// slices length and capacity
func slicesLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}

	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// nil slices
func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// create a slice with make
func createASliceWithMake() {
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)

	c := b[:2] // len(c)=2, cap(c)=5
	printSlice2("c", c)

	d := c[2:5] // len(d)=3, cap(d)=3
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// slices of slices
func slicesOfSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// appending to a slice
func appendingToASlice() {
	var s []int
	printSlice3(s) // len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0) // len=1 cap=1 [0]
	printSlice3(s)

	// The slice grows as needed.
	s = append(s, 1) // len=2 cap=2 [0 1]
	printSlice3(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4) // len=5 cap=6 [0 1 2 3 4]
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// range
func rangeExample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// range returns two values: the index and the value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// range continued
func rangeContinued() {
	pow := make([]int, 10)
	// you can skip the index or value by assigning to _
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// you can skip the index or value by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// maps
func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

// map literals
func mapLiterals() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

// map literals continued
func mapLiteralsContinued() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967}, // if the top-level type is just a type name, you can omit it from the elements of the literal.
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

// mutating maps
func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // The value: 48

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}

// function values
func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4)) // 5
}

// function closures
func functionClosures() {
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // 0 1 3 6 10 15 21 28 36 45
			neg(-2*i), // 0 -2 -6 -12 -20 -30 -42 -56 -72 -90
		)
	}
}
