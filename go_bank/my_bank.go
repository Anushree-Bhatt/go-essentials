package main

import (
	"fmt"
)

func bank_using_switch(balance float64) {
	for {
		ch := display_start()

		switch ch {
		case 1:
			fmt.Println("Your balance:", balance)
		case 2:
			var to_deposit float64
			fmt.Print("Enter the amount to deposit:")
			fmt.Scan(&to_deposit)

			if to_deposit <= 0 {
				fmt.Print("Invalid amount. You cannot deposit this amount")
				continue
			}

			balance = balance + to_deposit
			fmt.Println("Your amount is deposited. New balance:", balance)
		case 3:
			var to_withdraw float64
			fmt.Print("Enter the amount to withdraw:")
			fmt.Scan(&to_withdraw)

			if to_withdraw <= 0 {
				fmt.Println("Invalid amount. You cannot withdraw this amount")
				continue
			}
			if to_withdraw > balance {
				fmt.Println("Your withdrawal amount exceeds current balance!")
				continue
			}

			balance = balance - to_withdraw
			fmt.Println("You've withdrawn the amount. New balance:", balance)
		default:
			return

		}

	}
}

func bank_using_if(balance float64) {
	for {
		ch := display_start()

		if ch == 1 {
			fmt.Println("Your balance:", balance)
		} else if ch == 2 {
			var to_deposit float64
			fmt.Print("Enter the amount to deposit:")
			fmt.Scan(&to_deposit)

			if to_deposit <= 0 {
				fmt.Print("Invalid amount. You cannot deposit this amount")
				continue
			}

			balance = balance + to_deposit
			fmt.Println("Your amount is deposited. New balance:", balance)
		} else if ch == 3 {
			var to_withdraw float64
			fmt.Print("Enter the amount to withdraw:")
			fmt.Scan(&to_withdraw)

			if to_withdraw <= 0 {
				fmt.Println("Invalid amount. You cannot withdraw this amount")
				continue
			}
			if to_withdraw > balance {
				fmt.Println("Your withdrawal amount exceeds current balance!")
				continue
			}

			balance = balance - to_withdraw
			fmt.Println("You've withdrawn the amount. New balance:", balance)
		} else {
			return
		}
	}
}

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

func main() {
	fmt.Println("Welcome to Go Bank!")
	// bank_using_if(1000)
	bank_using_switch(1000)
	fmt.Println("Thank you for visiting!")
}
