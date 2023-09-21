package service

type bank struct {
	accountDetails map[int]int
}

func New(details map[int]int) *bank {
	return &bank{
		accountDetails: details,
	}
}

func (b *bank) CreateAccount() {

}

func (b *bank) Deposit(accNo int, amount int) {

}

func (b *bank) Withdraw(accNo int, amount int) {

}

func (b *bank) DisplayBalance(acctNo int) {

}
func (b *bank) DisplayBankBalance() {

}