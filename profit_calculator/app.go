package main

import (
	"fmt"
)

func calculateProfit() {
	var revenue float64
	var expenses float64
	const taxRate = 30

	fmt.Print("Enter the company's revenue:")
	fmt.Scan(&revenue)

	fmt.Print("Enter the company's expenses:")
	fmt.Scan(&expenses)

	ebt := revenue - expenses          //Earnings before tax
	eat := ebt - (ebt * taxRate / 100) //profit
	ebt_to_profit := ebt / eat         //Ratio of Earnings before tax to Profit

	fmt.Println("Earnings before Tax:", ebt)
	fmt.Println("Tax rate: 30%. Earnings after paying Tax:", eat)
	fmt.Println("Ratio (EBT/profit):", ebt_to_profit)

}

func main() {
	fmt.Println("Welcome to Profit Calculatore!")
	calculateProfit()
}
