package main

import (
	"fmt"
	"simplebankapp/pkg/domain"
	"simplebankapp/pkg/impl"
	"simplebankapp/pkg/model"
)

func main() {
	acc := &model.Account{
		Number:  "123-456-789",
		Balance: 1000000.0,
	}

	accOther := &model.Account{
		Number:  "987-654-321",
		Balance: 1000000.0,
	}

	transactions := []domain.Transaction{
		impl.Deposit{Amount: 250000.0},
		impl.Deposit{Amount: -250000.0},
		impl.Deposit{Amount: 500000.0},
		impl.Withdrawal{Amount: -500000.0},
		impl.Withdrawal{Amount: 500000.0},
	}

	transactionsTrf := []domain.TransactionTrf{
		impl.Transfer{Source: acc, Beneficiary: accOther, Amount: 250000.0},
		impl.Transfer{Source: acc, Beneficiary: accOther, Amount: -250000.0},
	}

	for _, tx := range transactions {
		fmt.Println("Melakukan:", tx.Description())
		err := tx.Apply(acc)
		if err != nil {
			fmt.Println("Gagal")
		} else {
			fmt.Println("Berhasil")
		}
		acc.PrintBalance()
	}

	for _, tx := range transactionsTrf {
		fmt.Println("Melakukan:", tx.Description())
		err := tx.Apply(acc, accOther)
		if err != nil {
			fmt.Println("Gagal")
		} else {
			fmt.Println("Berhasil")
		}
		acc.PrintBalance()
		accOther.PrintBalance()
	}
}
