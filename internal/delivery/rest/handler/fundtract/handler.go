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

func (ah *fundtractHandler) RegisterRoutes(router *echo.Group) {
}
