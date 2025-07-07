package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World! Welcome to GoLang")
	// futureValue()
	stringFormatting()
}

func futureValue() {
	//variables and types
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

func stringFormatting() {
	fmt.Print("Hello Hi Namaskara!\n")

	fmt.Println("Hello Hi")
	fmt.Println("Namaskara")

	fmt.Printf("%v %v %v\n", "Hello", "hi", "Namaskara\n")

	text := `Using specifiers %.0f and %.2f`
	fmt.Print(text)
	a := 10.6
	b := 20.43
	fmt.Printf("Numbers are %.0f and %.2f\n", a, b)

	fmt.Println("String concatenation using +")
	s1 := "Hey!"
	s2 := "How're you doing?"
	fmt.Printf("%v\n", (s1 + s2))

	//Multiple lines String using backticks - ` `
	fmt.Print(`Using Backticks.....
	...........
	..........
	Multiple Lines String!`)

}
