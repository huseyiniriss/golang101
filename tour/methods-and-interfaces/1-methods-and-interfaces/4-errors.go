package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return "at " + e.When.String() + ", " + e.What
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}
func main() {
	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Exercise: Errors
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	return 0, nil
}
