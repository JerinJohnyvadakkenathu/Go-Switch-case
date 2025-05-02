package main

import "fmt"

func main() {
	var Num1 int
	var Num2 int
	var Operator string
	var Result int
	fmt.Println("Enter the First value")
	fmt.Scan(&Num1)
	fmt.Println("Enter the second value")
	fmt.Scan(&Num2)
	fmt.Println("Enter which operator")
	fmt.Scan(&Operator)
	switch Operator {
	case "+":
		Result = Num1 + Num2
	case "-":
		Result = Num1 - Num2
	case "*":
		Result = Num1 * Num2
	case "/":
		Result = Num1 / Num2
	}
	fmt.Println(Result)
}
