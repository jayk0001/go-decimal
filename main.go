package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	qty := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString("0.08875")

	subTotal := price.Mul(qty)

	preTax := subTotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("SubTotal: ", subTotal)
	fmt.Println("Pre-tax: ", preTax)
	fmt.Println("Taxes: ", total.Sub(preTax))
	fmt.Println("Total: ", total)
	fmt.Println("Tax rate: ", total.Sub(preTax).Div(preTax))
}
