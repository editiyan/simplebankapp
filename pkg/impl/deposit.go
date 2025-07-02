package impl

import (
	"fmt"
	"simplebankapp/pkg/domain"
	"simplebankapp/pkg/model"
)

type Deposit struct {
	Amount float64
}

func (d Deposit) Apply(acc *model.Account) error {
	if d.Amount <= 0 {
		return fmt.Errorf("Jumlah Deposit harus > 0")
	}
	acc.Balance += d.Amount
	return nil
}

func (d Deposit) Description() string {
	return fmt.Sprintf("Deposit sebesar %.2f", d.Amount)
}

var _ domain.Transaction = Deposit{}
