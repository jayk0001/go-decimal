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

	// Read input until a newLine character is encountered
	input1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Remove the newline character and any surrounding whitespace
	first, _ := decimal.NewFromString(strings.TrimSpace(input1))

	fmt.Print("Enter the second number: ")

	// Read input until a newLine character is encountered
	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Remove the newline character and any surrounding whitespace
	second, _ := decimal.NewFromString(strings.TrimSpace(input2))

	fmt.Print("Enter the operation(+, -, /, *): ")

	// Read input until a newLine character is encountered
	input3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// Remove the newline character and any surrounding whitespace
	operation := strings.TrimSpace(input3)

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
