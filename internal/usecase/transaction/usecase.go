package transaction

import (
	"context"
	"time"
)

type Usecase interface {
	ProcessTransactions(ctx context.Context, startFrom time.Time)
}

type transactionUsecase struct {
}

func NewUsecase() Usecase {
	return &transactionUsecase{}
}
