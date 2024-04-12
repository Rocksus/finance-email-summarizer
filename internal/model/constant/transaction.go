package constant

type TransactionType int

const (
	TransactionTypeDebit TransactionType = iota
	TransactionTypeCredit
)
