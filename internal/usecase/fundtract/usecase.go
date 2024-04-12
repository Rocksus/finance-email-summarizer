package fundtract

import (
	"context"

	"github.com/Rocksus/fundtract/internal/model"
)

type Usecase interface {
	InsertTransaction(ctx context.Context, req model.InsertTransactionRequest) (model.AccountTransaction, error)
	ListAccountTransaction(ctx context.Context, accountID int64, limit, offset int) ([]model.AccountTransaction, model.PaginationData, error)
	InsertUserAccount(ctx context.Context, req model.InsertUserAccountRequest) (model.UserAccount, error)
	ListUserAccount(ctx context.Context, userID int64) ([]model.UserAccount, error)
	InsertUser(ctx context.Context, req model.InsertUserRequest) (model.User, error)
}

type fundtractUsecase struct {
}

func NewUsecase() Usecase {
	return &fundtractUsecase{}
}
