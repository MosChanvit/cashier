package handler

import (
	"github.com/labstack/echo"
)

type Handler interface {
	GetCashiers(c echo.Context) error
	GetCashier(c echo.Context) error
	NewCashier(c echo.Context) error
	ProcessTransaction(c echo.Context) error
}
