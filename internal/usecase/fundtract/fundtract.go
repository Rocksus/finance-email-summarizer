package fundtract

import (
        "context"
        "time"

        "github.com/Rocksus/fundtract/internal/model"
        "github.com/google/uuid"
)

func (fu *fundtractUsecase) InsertTransaction(ctx context.Context, req model.InsertTransactionRequest) (model.AccountTransaction, error) {
        txID := uuid.New().String()
        tx := model.AccountTransaction{
                TransactionID:              txID,
                AccountID:                  req.AccountID,
                UserID:                     req.UserID,
                TransactionName:            req.TransactionName,
                MagnifiedTransactionAmount: req.TransactionAmount,
                TransactionType:            req.TransactionType,
                CategoryID:                 req.CategoryID,
                CreatedAt:                  req.CreatedAt,
                UpdatedAt:                  time.Now(),
                Notes:                      req.Notes,
                CreatedBy:                  "system", // TODO: Get from context
        }
        return tx, nil
}

func (fu *fundtractUsecase) ListAccountTransaction(ctx context.Context, accountID int64, limit, offset int) ([]model.AccountTransaction, model.PaginationData, error) {
        // TODO: Implement database query
        return []model.AccountTransaction{}, model.PaginationData{
                Limit:   limit,
                Offset:  offset,
                HasNext: false,
        }, nil
}

func (fu *fundtractUsecase) InsertUserAccount(ctx context.Context, req model.InsertUserAccountRequest) (model.UserAccount, error) {
        acc := model.UserAccount{
                AccountID:               req.AccountID,
                UserID:                  req.UserID,
                AccountName:             req.AccountName,
                MagnifiedBalanceSummary: 0,
                CreatedAt:               time.Now(),
                UpdatedAt:               time.Now(),
                Currency:                req.Currency,
        }
        return acc, nil
}

func (fu *fundtractUsecase) ListUserAccount(ctx context.Context, userID int64) ([]model.UserAccount, error) {
        // TODO: Implement database query
        return []model.UserAccount{}, nil
}

func (fu *fundtractUsecase) InsertUser(ctx context.Context, req model.InsertUserRequest) (model.User, error) {
        user := model.User{
                Username:  req.Username,
                Email:     req.Email,
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
        }
        return user, nil
}

import (
        "context"

        "github.com/Rocksus/fundtract/internal/model"
)

func (s *fundtractUsecase) InsertTransaction(ctx context.Context, req model.InsertTransactionRequest) (model.AccountTransaction, error) {
        return model.AccountTransaction{}, nil
}
func (s *fundtractUsecase) ListAccountTransaction(ctx context.Context, accountID int64, limit, offset int) ([]model.AccountTransaction, model.PaginationData, error) {
        return nil, model.PaginationData{}, nil
}
func (s *fundtractUsecase) InsertUserAccount(ctx context.Context, req model.InsertUserAccountRequest) (model.UserAccount, error) {
        return model.UserAccount{}, nil
}
func (s *fundtractUsecase) ListUserAccount(ctx context.Context, userID int64) ([]model.UserAccount, error) {
        return nil, nil
}
func (s *fundtractUsecase) InsertUser(ctx context.Context, req model.InsertUserRequest) (model.User, error) {
        return model.User{}, nil
}
