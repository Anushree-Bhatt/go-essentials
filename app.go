package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World! Welcome to GoLang Future value prediction for invested amount^ _ ^")
	futureValue()
}

func futureValue() {
	const inflationRate = 6.0
	var principle float64
	var rate_of_interest float64
	var years float64

	fmt.Print("Enter the amount you want to invest:")
	fmt.Scan(&principle)

	fmt.Print("At what rate of interest or expected returns:")
	fmt.Scan(&rate_of_interest)

	fmt.Print("For how many years:")
	fmt.Scan(&years)

	futureValue := principle * math.Pow(1+rate_of_interest/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Expected returns:", futureValue)
	fmt.Println("Actual returns:", realFutureValue)

}
