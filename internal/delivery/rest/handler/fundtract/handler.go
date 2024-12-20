package fundtract

import (
        "github.com/Rocksus/fundtract/internal/delivery/rest"
        "github.com/Rocksus/fundtract/internal/usecase/fundtract"
        "github.com/labstack/echo/v4"
)

type fundtractHandler struct {
        fu fundtract.Usecase
}

func NewFundtractHandler(fu fundtract.Usecase) rest.API {
        return &fundtractHandler{
                fu: fu,
        }
}

func (h *fundtractHandler) RegisterRoutes(router *echo.Group) {
        router.POST("/transactions", h.handleInsertTransaction)
        router.GET("/accounts/:account_id/transactions", h.handleListAccountTransaction)
        router.POST("/accounts", h.handleInsertUserAccount)
        router.GET("/accounts", h.handleListUserAccount)
        router.POST("/users", h.handleInsertUser)
}
