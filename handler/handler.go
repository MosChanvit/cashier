package handler

import (
	"github.com/labstack/echo"
)

type Handler interface {
	GetCashiers(c echo.Context) error
	GetCashier(c echo.Context) error
	NewCashier(c echo.Context) error
}

// func handleError(w http.ResponseWriter, err error) {
// 	switch e := err.(type) {
// 	case errs.AppError:
// 		w.WriteHeader(e.Code)
// 		fmt.Fprintln(w, e)
// 	case error:
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, e)
// 	}
// }
