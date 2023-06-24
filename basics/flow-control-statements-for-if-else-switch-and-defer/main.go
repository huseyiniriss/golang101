package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	forLoop()
	forIsGosWhile()
	ifStatement()
	ifStatementWithShortStatement()
	ifAndElse()
	switchStatement()
	switchEvaluationOrder()
	switchWithNoCondition()
	deferStatement()
	stackingDefers()
}

// for loop
func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// for loop continued
// for is Go's "while"
func forIsGosWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// for loop continued
// forever
// for {
// }
func forever() {
	for {
	}
}

// if statement
func ifStatement() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}
}

// if statement with a short statement
func ifStatementWithShortStatement() {
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}
}

// if and else
func ifAndElse() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
}

// switch
func switchStatement() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// switch evaluation order
func switchEvaluationOrder() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// switch with no condition
func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// defer
func deferStatement() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// stacking defers
func stackingDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
