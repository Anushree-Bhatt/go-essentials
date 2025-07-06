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
	principle := 1000
	rate_of_interest := 5.5
	var years float64 = 10 // explicit type assignment

	futureValue := float64(principle) * math.Pow(1+rate_of_interest/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Expected returns:", futureValue)
	fmt.Println("Actual returns:", realFutureValue)

}
