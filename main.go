package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Print("Enter the first number: ")

	// Create a new reader for standard input
	reader := bufio.NewReader(os.Stdin)

	// Read input1 until a newLine character is encountered
	input1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character and any surrounding whitespace,
	// create decimal using user entered string input
	first, _ := decimal.NewFromString(strings.TrimSpace(input1))

	fmt.Print("Enter the second number: ")

	// Read input2 until a newLine character is encountered
	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character and any surrounding whitespace,
	// create decimal using user entered string input
	second, _ := decimal.NewFromString(strings.TrimSpace(input2))

	fmt.Print("Enter the operation(+, -, /, *): ")

	// Read input3 until a newLine character is encountered
	input3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Remove the newline character and any surrounding whitespace, this will be operation
	operation := strings.TrimSpace(input3)

	// Using switch + case handling in golang, perform correct calculate operation
	// from decimal library, return result after caculcate
	switch operation {
	case "+":
		res := first.Add(second)
		fmt.Println("RESULT:", res)
	case "-":
		res := first.Sub(second)
		fmt.Println("RESULT:", res)
	case "*":
		res := first.Mul(second)
		fmt.Println("RESULT:", res)
	case "/":
		res := first.Div(second)
		fmt.Println("RESULT:", res)
	}
}
