package web

import (
	"html/template"
	"net/http"

	"github.com/Rocksus/fundtract/internal/usecase/fundtract"
	"github.com/labstack/echo/v4"
)

type webHandler struct {
	fu       fundtract.Usecase
	template *template.Template
}

func NewWebHandler(fu fundtract.Usecase) *webHandler {
	tmpl := template.Must(template.ParseGlob("internal/delivery/rest/templates/*.html"))
	return &webHandler{
		fu:       fu,
		template: tmpl,
	}
}

func (h *webHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.handleHome)
	e.GET("/login", h.handleLogin)
	e.GET("/dashboard", h.handleDashboard)
	e.GET("/transactions/new", h.handleNewTransaction)
}

func (h *webHandler) handleHome(c echo.Context) error {
	// TODO: Check if user is logged in
	return c.Redirect(http.StatusFound, "/login")
}

func (h *webHandler) handleLogin(c echo.Context) error {
	data := map[string]interface{}{
		"Title": "Login",
	}
	return h.template.ExecuteTemplate(c.Response().Writer, "login.html", data)
}

func (h *webHandler) handleDashboard(c echo.Context) error {
	// TODO: Check if user is logged in
	userID := int64(1)

	accounts, err := h.fu.ListUserAccount(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Title":    "Dashboard",
		"Accounts": accounts,
	}
	return h.template.ExecuteTemplate(c.Response().Writer, "dashboard.html", data)
}

func (h *webHandler) handleNewTransaction(c echo.Context) error {
	// TODO: Check if user is logged in
	userID := int64(1)

	accounts, err := h.fu.ListUserAccount(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Accounts": accounts,
		"Currency": "USD", // TODO: Get from selected account
	}
	return h.template.ExecuteTemplate(c.Response().Writer, "transaction_form.html", data)
}