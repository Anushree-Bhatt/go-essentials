package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balance_sheet = "balance.txt"

func readBalFromFile() (float64, error) {
	data, err := os.ReadFile(balance_sheet)

	if err != nil {
		custom_err := errors.New("file not found, hence defaulting balance to 0")
		return 0, custom_err
	}

	bal_txt := string(data)
	bal, err := strconv.ParseFloat(bal_txt, 64)

	if err != nil {
		custom_err := errors.New("file doesn't contain proper data, hence defaulting balance to 0")
		return 0, custom_err
	}

	return bal, err
}

func writeBalToFile(bal float64) {
	bal_txt := fmt.Sprint(bal)
	err := os.WriteFile(balance_sheet, []byte(bal_txt), 0644)

	if err != nil {
		fmt.Print("Writing file error:", err)
	}
}

func bank_using_switch(balance float64) float64 {
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
			return balance
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

	balance_data, err := readBalFromFile()
	if err != nil {
		fmt.Println(err)
		panic("Stop running....")
	}
	bal := bank_using_switch(balance_data)
	writeBalToFile(bal)

	fmt.Println("Thank you for visiting!")
}
