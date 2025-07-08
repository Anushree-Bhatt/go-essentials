package main

import (
	"errors"
	"fmt"
	"os"
)

const taxRate = 30

func calculateProfit() {
	revenue := userInput("revenue")
	expenses := userInput("expenses")

	err := validateInputs(revenue, expenses)
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, eat, ebt_profit := calculations(revenue, expenses)
	writeToFile(ebt, eat, ebt_profit)

	fmt.Printf("Earnings before Tax: %.3f\n", ebt)
	fmt.Printf("Tax rate: 30 percent. Earnings after paying Tax: %.3f\n", eat)
	fmt.Printf("Ratio (EBT/profit): %.3f\n", ebt_profit)
}

func validateInputs(revenue float64, expenses float64) error {
	var msg string

	if revenue <= 0 && expenses <= 0 {
		msg = "revenue and expense"
	} else if expenses <= 0 {
		msg = "expenses"
	} else if revenue <= 0 {
		msg = "revenue"
	} else {
		return nil
	}

	return errors.New(msg + " cannot be negative or zero. Try again later")

}

func writeToFile(ebt, eat, ebt_profit float64) {
	data := fmt.Sprintf("ebt:%.3f\neat:%.3f\nratio:%.3f", ebt, eat, ebt_profit)
	os.WriteFile("profit_calculations.txt", []byte(data), 0644)
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
