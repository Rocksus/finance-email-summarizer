package fundtract

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
