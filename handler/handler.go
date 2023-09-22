package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"digitalbank/commands"
)

var accountNumber int

type handler struct {
	opts Operations
}

func NewHandler(o Operations) *handler {
	return &handler{
		opts: o,
	}
}

func (h *handler) BankOperations(input string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error opening input file: %v", err)
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		command := strings.Fields(line)

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case commands.CreateAccount:
			accountNumber++
			h.opts.CreateAccount(accountNumber)
		case commands.DepositMoney:
			if len(command) != 3 {
				fmt.Println("Invalid deposit command")
				continue
			}
			acctNo, amount := parseArgs(command)
			h.opts.Deposit(acctNo, amount)
		case commands.WithdrawMoney:
			if len(command) != 3 {
				fmt.Println("Invalid withdraw command")
				continue
			}
			acctNo, amount := parseArgs(command)
			h.opts.Withdraw(acctNo, amount)
		case commands.DisplayBalance:
			if len(command) != 2 {
				fmt.Println("Invalid display_balance command")
				continue
			}
			acctNo := parseAccountNumber(command[1])
			h.opts.DisplayBalance(acctNo)
		case commands.DisplayBankBalance:
			h.opts.DisplayBankBalance()
		default:
			fmt.Printf("Invalid command: %s\n", command[0])
		}
	}

	if err = scan.Err(); err != nil {
		fmt.Printf("Error reading input command file: %v\n", err)
		return
	}

}

func parseArgs(args []string) (int, int) {
	actNo := parseAccountNumber(args[1])
	amount := parseAmount(args[2])
	return actNo, amount
}

func parseAccountNumber(accountStr string) int {
	actNo, _ := strconv.Atoi(accountStr)
	return actNo
}

func parseAmount(amountStr string) int {
	amount, _ := strconv.Atoi(amountStr)
	return amount
}
