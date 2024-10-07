package main

import (
	"GOLANG/calci"
	"fmt"
)

func main() {
	var a, b float64
	var option int
	var result float64

	fmt.Println("Select an option: 1. Add, 2. Subtract, 3. Multiply, 4. Divide")
	fmt.Print("Enter the option: ")
	fmt.Scanln(&option)
	fmt.Print("Enter the value of a: ")
	fmt.Scan(&a)
	fmt.Print("Enter the value of b: ")
	fmt.Scan(&b)
	if option == 1 {
		result = calci.Add(a, b)
		fmt.Printf("Result: %.2f\n", result)
	} else if option == 2 {
		result = calci.Subtract(a, b)
		fmt.Printf("Result: %.2f\n", result)
	} else if option == 3 {
		result = calci.Multiply(a, b)
		fmt.Printf("Result: %.2f\n", result)
	} else if option == 4 {
		result = calci.Divide(a, b)
		fmt.Printf("Result: %.2f\n", result)
	} else {
		fmt.Println("error")
	}
}
