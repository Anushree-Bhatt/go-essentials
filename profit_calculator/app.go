package main

import (
	"fmt"
)

const taxRate = 30

func calculateProfit() {
	revenue := userInput("revenue")
	expenses := userInput("expenses")

	ebt, eat, ebt_profit := calculations(revenue, expenses)

	fmt.Printf("Earnings before Tax: %.3f\n", ebt)
	fmt.Printf("Tax rate: 30 percent. Earnings after paying Tax: %.3f\n", eat)
	fmt.Printf("Ratio (EBT/profit): %.3f\n", ebt_profit)
}

func calculations(revenue, expenses float64) (float64, float64, float64) {
	ebt := revenue - expenses          //Earnings before tax
	eat := ebt - (ebt * taxRate / 100) //profit - earnings after tax
	ebt_to_profit := ebt / eat

	return ebt, eat, ebt_to_profit
}

func userInput(info string) float64 {
	var input float64
	fmt.Printf("Enter the Company's %v:", info)
	fmt.Scan(&input)
	return input
}

func main() {
	fmt.Println("Welcome to Profit Calculator!")
	calculateProfit()
}
