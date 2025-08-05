package mobilebanking

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const balanceFile = "mobile-banking/balance.txt"

func StartBanking() {

	balance := getBalance()

	for {

		fmt.Println("What you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Input: ")

		input := getInput()

		switch input {

		case "1":
			fmt.Printf("Your account balance is %.2f\n", balance)

		case "2":
			fmt.Print("Your deposit: ")
			depositAmt := getInput()
			amount, err := strconv.ParseFloat(depositAmt, 64)
			if err != nil || amount < 0 {
				fmt.Println("Invalid deposit amount.")

			}
			balance += amount
			fmt.Printf("Balance updated! Amount: %.2f\n", balance)
			saveBalance(balance)

		case "3":
			fmt.Print("Withdrawal amount: ")
			withdrawAmt := getInput()
			amount, err := strconv.ParseFloat(withdrawAmt, 64)
			if err != nil || amount < 0 {
				fmt.Println("Invalid withdraw amount.")
			}
			if amount > balance {
				fmt.Println("Insufficient funds.")
			} else {
				balance -= amount
				fmt.Printf("Balance updated! Amount: %.2f\n", balance)
				saveBalance(balance)
			}

		case "4":
			fmt.Println("Thanks for banking with us")

		default:
			fmt.Println("Invalid choice.")
		}

		if input == "4" {
			break
		}
	}
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func getBalance() float64 {
	data, err := os.ReadFile(balanceFile)
	//check(err)

	if err != nil {
		fmt.Println("Error while reading file:", err)
		return 0
	}

	res, err := strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
	if err != nil {
		fmt.Println("Error while parsing:", err)
		return 0
	}
	return res
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func saveBalance(balance float64) {
	file, err := os.Create(balanceFile)
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%.2f", balance)
	file.Close()
}
