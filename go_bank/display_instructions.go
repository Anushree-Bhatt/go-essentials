package main

import "fmt"

func display_start() int {
	var ch int

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice:")
	fmt.Scan(&ch)

	return ch
}
