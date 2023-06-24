package main

import "fmt"

type Calculator struct {
	Num1 int
	Num2 int
}

func (c Calculator) Add() int {
	return c.Num1 + c.Num2
}

func (c Calculator) Sub() int {
	return c.Num1 - c.Num2
}

func (c Calculator) Mul() int {
	return c.Num1 * c.Num2
}

func (c Calculator) Div() int {
	return c.Num1 / c.Num2
}

func main() {
	for {
		fmt.Println("--- CALCULATOR ---")
		fmt.Println("1. Add")
		fmt.Println("2. Sub")
		fmt.Println("3. Mul")
		fmt.Println("4. Div")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		condition := 0
		fmt.Scanf("%d", &condition)
		switch condition {
		case 1:
			c := Calculator{}
			fmt.Print("Enter num1: ")
			fmt.Scanf("%d", &c.Num1)
			fmt.Print("Enter num2: ")
			fmt.Scanf("%d", &c.Num2)
			fmt.Println("Result: ", c.Add())
		case 2:
			c := Calculator{}
			fmt.Print("Enter num1: ")
			fmt.Scanf("%d", &c.Num1)
			fmt.Print("Enter num2: ")
			fmt.Scanf("%d", &c.Num2)
			fmt.Println("Result: ", c.Sub())
		case 3:
			c := Calculator{}
			fmt.Print("Enter num1: ")
			fmt.Scanf("%d", &c.Num1)
			fmt.Print("Enter num2: ")
			fmt.Scanf("%d", &c.Num2)
			fmt.Println("Result: ", c.Mul())
		case 4:
			c := Calculator{}
			fmt.Print("Enter num1: ")
			fmt.Scanf("%d", &c.Num1)
			fmt.Print("Enter num2: ")
			fmt.Scanf("%d", &c.Num2)
			fmt.Println("Result: ", c.Div())
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
