package domain

import "time"

type BaseTransaction struct {
	ByUser string
	AtTime time.Time
	Notes  string
}

func (b BaseTransaction) AuditTrail() string {
	return "[" + b.ByUser + "] at " + b.AtTime.Format(time.RFC3339) + ", dengan catatan : " + b.Notes
}
