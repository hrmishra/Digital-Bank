package handler

type Operations interface {
	CreateAccount(acctNo int)
	Deposit(acctNo int, amount int)
	Withdraw(acctNo int, amount int)
	DisplayBalance(acctNo int)
	DisplayBankBalance()
}
