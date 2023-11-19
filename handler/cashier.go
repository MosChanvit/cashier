package handler

import (
	"cashier/logs"
	"cashier/service"
	"net/http"
	"strconv"

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
		return c.JSONPretty(http.StatusOK, res, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) GetCashier(c echo.Context) error {

	id := c.QueryParam(`id`)
	idReq, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	res, err := s.castSrv.GetCashier(idReq)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusNoContent, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) NewCashier(c echo.Context) error {

	req := service.NewCashierRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.IdCashier == "" {
		return c.JSONPretty(http.StatusBadRequest, "id_cashier is required", "")
	}

	res, err := s.castSrv.NewCashier(req)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusOK, res, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) ProcessTransaction(c echo.Context) error {

	req := service.ProcessTransactionRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.IdCashier == "" {
		return c.JSONPretty(http.StatusBadRequest, "id_cashier is required", "")
	}

	res, err := s.castSrv.ProcessTransaction(req)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusOK, res, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}
