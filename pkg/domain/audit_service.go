package domain

type Logger interface {
	Log()
}

type Notifier interface {
	Notify()
}

type AuditService interface {
	Logger
	Notifier
}
