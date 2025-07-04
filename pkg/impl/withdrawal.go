package impl

import (
	"fmt"
	"simplebankapp/pkg/domain"
	"simplebankapp/pkg/model"
)

type Withdrawal struct {
	domain.BaseTransaction
	Amount          float64
	model.AuditInfo //composition reuse
}

func (w Withdrawal) Apply(acc *model.Account) error {
	if w.Amount <= 0 {
		return fmt.Errorf("jumlah Withdrawal harus > 0")
	}
	acc.Balance -= w.Amount
	return nil
}

func (w Withdrawal) Description() string {
	return fmt.Sprintf("Withdrawal sebesar %.2f", w.Amount)
}

var _ domain.Transaction = Withdrawal{}
