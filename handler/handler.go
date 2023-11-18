package handler

import (
	"cashier/errs"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type API interface {
	GetCashier(c echo.Context) error
}

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}
