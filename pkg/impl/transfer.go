package impl

import (
	"fmt"
	"simplebankapp/pkg/domain"
	"simplebankapp/pkg/model"
)

type Transfer struct {
	Source      *model.Account
	Beneficiary *model.Account
	Amount      float64
}

func (t Transfer) Apply(accSource *model.Account, accBene *model.Account) error {
	if t.Amount <= 0 {
		return fmt.Errorf("Jumlah Transfer harus > 0")
	}
	accSource.Balance -= t.Amount
	accBene.Balance += t.Amount
	return nil
}

func (d Transfer) Description() string {
	return fmt.Sprintf("Transfer sebesar %.2f dari rekening %v ke rekening %v", d.Amount, d.Source.Number, d.Beneficiary.Number)
}

var _ domain.TransactionTrf = Transfer{}
