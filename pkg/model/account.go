package model

import "fmt"

type Account struct {
	Number  string
	Balance float64
}

func (acc *Account) PrintBalance() {
	fmt.Printf("saldo rekening %s : %.2f\n", acc.Number, acc.Balance)
}
