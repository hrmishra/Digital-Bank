package service

import (
	"fmt"
)

type bank struct {
	accountDetails map[int]int
}

func New() *bank {
	return &bank{
		accountDetails: make(map[int]int),
	}
}

func (b *bank) CreateAccount(acctNo int) {
	b.accountDetails[acctNo] = 0
	fmt.Printf("New account created with account number %d\n", acctNo)
}

func (b *bank) Deposit(acctNo int, amount int) {
	if balance, ok := b.accountDetails[acctNo]; ok {
		b.accountDetails[acctNo] = balance + amount
		fmt.Printf("%d deposited to account number %d\n", amount, acctNo)
	} else {
		fmt.Printf("Account number %d not found\n", acctNo)
	}
}

func (b *bank) Withdraw(acctNo int, amount int) {
	if balance, ok := b.accountDetails[acctNo]; ok {
		if balance >= amount {
			b.accountDetails[acctNo] = balance - amount
			fmt.Printf("%d withdrawn from account number %d\n", amount, acctNo)
		} else {
			fmt.Println("Withdrawal limit exceeded")
		}
	} else {
		fmt.Printf("Account number %d not found\n", acctNo)
	}
}

func (b *bank) DisplayBalance(acctNo int) {
	if balance, ok := b.accountDetails[acctNo]; ok {
		fmt.Printf("Balance for account number %d: %d\n", acctNo, balance)
	} else {
		fmt.Printf("Account number %d not found\n", acctNo)
	}
}
func (b *bank) DisplayBankBalance() {
	totalBal := 0
	for _, bal := range b.accountDetails {
		totalBal = totalBal + bal
	}

	fmt.Printf("Total bank balance: %d\n", totalBal)
}
