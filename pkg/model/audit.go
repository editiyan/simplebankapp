package model

import "time"

type AuditInfo struct {
	ByUser string
	AtTime time.Time
	Notes  string
}

func (a AuditInfo) Summary() string {
	return "[" + a.ByUser + "] at " + a.AtTime.Format(time.RFC3339) + ", dengan catatan : " + a.Notes
}
