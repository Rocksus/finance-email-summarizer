package model

import (
	"time"

	"github.com/Rocksus/fundtract/internal/model/constant"
)

type User struct {
	UserID    int64
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InsertUserRequest struct {
	Username string
	Email    string
	Password string
}

type UserAuth struct {
	UserID       int64
	PasswordHash string
}

type UserAccount struct {
	AccountID               int64
	UserID                  int64
	AccountName             string
	MagnifiedBalanceSummary int64
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Currency                string
}

type InsertUserAccountRequest struct {
	AccountID   int64
	UserID      int64
	AccountName string
	Currency    string
}

type AccountTransaction struct {
	TransactionID              string
	AccountID                  int64
	UserID                     int64
	TransactionName            string
	MagnifiedTransactionAmount int64
	TransactionType            constant.TransactionType
	CategoryID                 int64
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	Notes                      string
	CreatedBy                  string
}

type AccountTransactionApproval struct {
	ApprovalID                 string
	AccountID                  int64
	UserID                     int64
	TransactionName            string
	MagnifiedTransactionAmount int64
	TransactionType            constant.TransactionType
	CategoryID                 int64
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	Notes                      string
	CreatedBy                  string
}

type CurrencyMagnifier struct {
	Currency  string
	Magnifier int
}

type UserAPIKey struct {
	APIKeyID      string
	APISecretHash string
	Identifier    string
	CreatedAt     time.Time
	DeletedAt     *time.Time
}

type InsertTransactionRequest struct {
	AccountID         int64
	UserID            int64
	TransactionName   string
	TransactionAmount int64
	TransactionType   constant.TransactionType
	CategoryID        int64
	CreatedAt         time.Time
	Notes             string
}

type PaginationData struct {
	Limit   int
	Offset  int
	HasNext bool
}
