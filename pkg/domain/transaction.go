package domain

import "simplebankapp/pkg/model"

type Transaction interface {
	Apply(acc *model.Account) error
	Description() string
}

type TransactionTrf interface {
	Apply(accS *model.Account, accB *model.Account) error
	Description() string
}
