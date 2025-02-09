package main

import (
	"fmt"

	"github.com/Rocksus/fundtract/internal/platform/storage/sqlite"
	"github.com/labstack/echo/v4"
)

func InitApp() *echo.Echo {
	e := echo.New()

	_ = sqlite.New()

	return e
}

func main() {
	app := InitApp()
	app.Logger.Fatal(app.Start(fmt.Sprint(":", 14045)))
}
