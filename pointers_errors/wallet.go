package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds!")

type Wallet struct {
	Balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.Balance {
		return ErrInsufficientFunds
	}

	w.Balance -= amount
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("WELCOME TO YOUR BITCOIN WALLET MANAGER!")
	fmt.Println(strings.Repeat("-", 40))

	w := Wallet{}

	for {
		fmt.Println("1 - Deposit Bitcoin\n2 - Withdraw Bitcoin\n3 - See Amount\n4 - Exit")
		scanner.Scan()
		userOption, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number:", err)
		}

		var userAmount int

		if userOption == 1 || userOption == 2 {
			fmt.Println("Enter the amount:")
			scanner.Scan()
			userAmount, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("invalid number:", err)
			}
		}

		switch userOption {
		case 1:
			w.Deposit(Bitcoin(userAmount))
		case 2:
			err := w.Withdraw(Bitcoin(userAmount))
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			fmt.Println("Your Bitcoin Amount:" + strconv.Itoa(int(w.Balance)))
		case 4:
			fmt.Println("\nSee you soon!")
			return
		default:
			fmt.Println("Error! Please enter a number that is listed.")
		}
	}

}
