# Go-Decimal Basic Arithmetic Calculator

## What it is ...
This is a simple arithmetic calculator in go language, this program used Decimal package [https://github.com/shopspring/decimal] for Arbitrary-precision fixed-point decimal numbers in go.

## Why I built this is ...
### 0.1 + 0.2 != 0.3 (?????????)

In computer programming, there is a popular observation which 0.1 + 0.2 != 0.3, becasue of the way how computers represent numbers, 
I built this program to learn how I can implement simple four operation calculator that can calculate float64 type of number which is user input. 
(* This program can calculate any number user enter, not only limited to float64)

The core reason for this discrepancy is the inability to presicely represent certain decimal fractions in binary(base-2) format.

### Binary Representations of Decimals:

Computers natively operate in binary. While integers have exact binary representations, many decimal fractions, like 0.1, 0.2, and 0.3, do not have exact finite binary representations.
Similar to how 1/3 in decimal is a repeating decimal (0.333...), 0.1 in binary is a repeating binary fraction(0.0001100110011...).

Because 0.1, 0.2, and 0.3 are repeating in binary, they cannot be stored with perfect precision in the finite number of bits allocated for floating-point numbers. They are stored as the cloeset possible approximation.

### Accumulation of Small Errors:

When you add 0.1 and 0.2, the computer is actually adding their slightly inaccurate binary approximations. These small inaccuracies accumulate, leading to a sum that is very close to, but not exactly equal to, 0.3. 

For instance, in many languages, **0.1 + 0.2 evalutates to 0.30000000000000004**


### Implications:

Comparison of Floating-Point Numbers: Directly comparing floating-point numbers for equality using == can let to unexpected false results. Instead, it is recommended to check if the absoulute difference between the two numbers is less than a small tolerance(epsilon).
**Financial Calculations**: For applications requiring high precision, like financial calculations, specialized data type that handle arbitrary-precision decimal numbers(e.g., golang decimal package in this case) are often used to avoid these floating-point inaccuracies.


## How to use it ...
To install: 
```bash
git clone https://github.com/jayk0001/go-decimal.git
cd go-decimal
```

To run: 
```bash
go run main.go
```

The program will prompt 3 questions that taking user input in terminal following by enter key pressed to move next questions, enter desired number and operation in each question.
1) Enter the first number: 
2) Enter the second number: 
3) Enter the operations:
Then result will be shown

For instance, 
```bash
$ go run main.go 
Enter the first number: 0.1 # and press enter
Enter the second number: 0.2 # and press enter
Enter the operation(+, -, /, *): + # and press enter
RESULT: 0.3 
```
<img width="332" height="73" alt="Screenshot 2025-11-11 at 21 22 23" src="https://github.com/user-attachments/assets/be38e51a-18e7-4148-95b4-dfd0048bcedf" />


## How it works ... 
1) I create a reader for standard input
2) Read input until a newLine character is encountered
3) Remove the newLine character and any surrounding whitespace, then create decimal number using user entered string input, I used decimal.NewFromString() method from decimal package.
4) Repeat it for the second number user entered
5) When user enters operation, it goes to swith + case statement and use built-in arithmetic method in decimal type which already exists in user input1 that created from step2, then calculate and return result.


