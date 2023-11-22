package handler

import (
	"cashier/logs"
	"cashier/service"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type cashierHandler struct {
	castSrv service.CashierService
}

func NewCashierHandler(castSrv service.CashierService) Handler {
	return cashierHandler{castSrv: castSrv}
}

func (s cashierHandler) GetCashiers(c echo.Context) error {

	res, err := s.castSrv.GetCashiers()
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusBadRequest, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) GetCashier(c echo.Context) error {

	name := c.QueryParam(`name`)

	if name == "" {
		return c.JSONPretty(http.StatusBadRequest, "required name", "")
	}

	res, err := s.castSrv.GetCashier(name)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusBadRequest, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) NewCashier(c echo.Context) error {

	req := service.NewCashierRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.Name == "" {
		return c.JSONPretty(http.StatusBadRequest, "id_cashier is required", "")
	}

	res, err := s.castSrv.NewCashier(req)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusBadRequest, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) ProcessTransaction(c echo.Context) error {

	req := service.ProcessTransactionRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.Name == "" {
		return c.JSONPretty(http.StatusBadRequest, "id_cashier is required", "")
	}

	change := req.CustomerPaid - req.ProductPrice
	if change < 0 {
		return c.JSONPretty(http.StatusBadRequest, "Not paying enough for the product", "")
	}

	res, err := s.castSrv.ProcessTransaction(req)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusOK, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) CalXYZ(c echo.Context) error {

	rawNumbers := c.QueryParams().Get("numbers")
	values := strings.Split(rawNumbers, ",")

	res := s.castSrv.CalXYZ(values)
	// if err != nil {
	// 	logs.Error(err)
	// 	return c.JSONPretty(http.StatusOK, nil, "")
	// }

	return c.String(http.StatusOK, res)

}
