package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/NursultanNurgaliyev/Assignment1/Bank"
)

func runner(BankAccount BankAccount, scanner bufio.Scanner) {
	for {
		fmt.Println("\nBank account system")
		fmt.Println("Options:")
		fmt.Println("1. Deposit – will trigger Deposit method")
		fmt.Println("2. Withdraw – will trigger Withdraw method")
		fmt.Println("3. Get balance – will trigger GetBalance method")
		fmt.Println("4. Exit – finish program execution.")
		fmt.Print("\nInput number: ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print("Input amount: ")
			scanner.Scan()
			amountText := scanner.Text()
			amount, err := strconv.ParseFloat(amountText, 64)
			if err != nil {
				fmt.Println("Invalid amount. Please enter a valid number.")
				continue
			}
			BankAccount.Deposit(amount)
		case "2":
			fmt.Print("Input amount: ")
			scanner.Scan()
			amountText := scanner.Text()
			amount, err := strconv.ParseFloat(amountText, 64)
			if err != nil {
				fmt.Println("Invalid amount. Please enter a valid number.")
				continue
			}
			BankAccount.Withdraw(amount)
		case "3":
			BankAccount.GetBalance()
		case "4":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func main() {
	BankAccount := NewBankAccount("1", "Nursultan Nurgaliyev")
	scanner := bufio.NewScanner(os.Stdin)

	runner(*BankAccount, *scanner)
}
