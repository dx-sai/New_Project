package main

import (
	"CalculatorApp/calculator"
	"fmt"
)

func main() {
	a := 15.0
	b := 20.0

	adds := calculator.Add(a, b)
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, adds)

	subs := calculator.Subtract(a, b)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, subs)

	multips := calculator.Multiply(a, b)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, multips)

	divs := calculator.Divide(a, b)

	fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, divs)
}
