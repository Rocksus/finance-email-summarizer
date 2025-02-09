package user

import (
	"context"

	"github.com/Rocksus/fundtract/internal/model"
)

type Usecase interface {
	GetUserById(ctx context.Context, userId int) model.UserData
	Login(ctx context.Context)
}

type userUsercase struct {
}

func NewUsecase() Usecase {
	return &userUsercase{}
}
