package fundtract

import (
        "net/http"
        "strconv"
        "time"

        "github.com/Rocksus/fundtract/internal/model"
        "github.com/Rocksus/fundtract/internal/model/constant"
        "github.com/labstack/echo/v4"
)

type InsertTransactionRequest struct {
        AccountID         int64                      `json:"account_id" validate:"required"`
        TransactionName   string                     `json:"transaction_name" validate:"required"`
        TransactionAmount int64                      `json:"transaction_amount" validate:"required"`
        TransactionType   constant.TransactionType   `json:"transaction_type" validate:"required"`
        CategoryID        int64                      `json:"category_id" validate:"required"`
        Notes             string                     `json:"notes"`
}

type InsertUserAccountRequest struct {
        AccountName string `json:"account_name" validate:"required"`
        Currency    string `json:"currency" validate:"required,len=3"`
}

type InsertUserRequest struct {
        Username string `json:"username" validate:"required"`
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
}

// @Summary Insert a new transaction
// @Description Create a new transaction for a user account
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body InsertTransactionRequest true "Transaction details"
// @Success 200 {object} model.AccountTransaction
// @Router /transactions [post]
func (h *fundtractHandler) handleInsertTransaction(c echo.Context) error {
        var req InsertTransactionRequest
        if err := c.Bind(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        if err := c.Validate(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        // TODO: Get userID from context
        userID := int64(1)

        tx, err := h.fu.InsertTransaction(c.Request().Context(), model.InsertTransactionRequest{
                AccountID:         req.AccountID,
                UserID:            userID,
                TransactionName:   req.TransactionName,
                TransactionAmount: req.TransactionAmount,
                TransactionType:   req.TransactionType,
                CategoryID:        req.CategoryID,
                CreatedAt:         time.Now(),
                Notes:             req.Notes,
        })
        if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, tx)
}

// @Summary List account transactions
// @Description Get a list of transactions for a specific account
// @Tags transactions
// @Accept json
// @Produce json
// @Param account_id path int true "Account ID"
// @Param limit query int false "Number of items per page" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} model.AccountTransaction
// @Router /accounts/{account_id}/transactions [get]
func (h *fundtractHandler) handleListAccountTransaction(c echo.Context) error {
        accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
        if err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, "invalid account_id")
        }

        limit, _ := strconv.Atoi(c.QueryParam("limit"))
        if limit <= 0 {
                limit = 10
        }

        offset, _ := strconv.Atoi(c.QueryParam("offset"))
        if offset < 0 {
                offset = 0
        }

        txs, pagination, err := h.fu.ListAccountTransaction(c.Request().Context(), accountID, limit, offset)
        if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, map[string]interface{}{
                "transactions": txs,
                "pagination":   pagination,
        })
}

// @Summary Create a new user account
// @Description Create a new account for a user
// @Tags accounts
// @Accept json
// @Produce json
// @Param request body InsertUserAccountRequest true "Account details"
// @Success 200 {object} model.UserAccount
// @Router /accounts [post]
func (h *fundtractHandler) handleInsertUserAccount(c echo.Context) error {
        var req InsertUserAccountRequest
        if err := c.Bind(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        if err := c.Validate(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        // TODO: Get userID from context
        userID := int64(1)

        acc, err := h.fu.InsertUserAccount(c.Request().Context(), model.InsertUserAccountRequest{
                UserID:      userID,
                AccountName: req.AccountName,
                Currency:    req.Currency,
        })
        if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, acc)
}

// @Summary List user accounts
// @Description Get a list of accounts for a user
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {array} model.UserAccount
// @Router /accounts [get]
func (h *fundtractHandler) handleListUserAccount(c echo.Context) error {
        // TODO: Get userID from context
        userID := int64(1)

        accounts, err := h.fu.ListUserAccount(c.Request().Context(), userID)
        if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, accounts)
}

// @Summary Create a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body InsertUserRequest true "User details"
// @Success 200 {object} model.User
// @Router /users [post]
func (h *fundtractHandler) handleInsertUser(c echo.Context) error {
        var req InsertUserRequest
        if err := c.Bind(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        if err := c.Validate(&req); err != nil {
                return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }

        user, err := h.fu.InsertUser(c.Request().Context(), model.InsertUserRequest{
                Username: req.Username,
                Email:    req.Email,
                Password: req.Password,
        })
        if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }

        return c.JSON(http.StatusOK, user)
}
