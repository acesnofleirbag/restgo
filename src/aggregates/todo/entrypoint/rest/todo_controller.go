package rest

import (
	"api/src/aggregates/todo/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @@@: Dependency injection
func Store(ctx echo.Context) error {
    services.Store()

    return ctx.String(http.StatusOK, "Store todo")
}

func Show(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "Show")
}

func Index(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "Index")
}

func Update(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "Update")
}

func Destroy(ctx echo.Context) error {
    return ctx.String(http.StatusOK, "Destroy")
}
