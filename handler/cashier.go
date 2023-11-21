package handler

import (
	"cashier/logs"
	"cashier/service"
	"encoding/json"
	"fmt"
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

	res2B, _ := json.Marshal(req)
	fmt.Println(string(res2B))

	if req.Name == "" {
		return c.JSONPretty(http.StatusBadRequest, "id_cashier is required", "")
	}

	res, err := s.castSrv.ProcessTransaction(req)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusOK, err, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}

func (s cashierHandler) CalValue(c echo.Context) error {

	rawNumbers := c.QueryParams().Get("numbers")
	fmt.Println(rawNumbers)
	values := strings.Split(rawNumbers, ",")

	// var numbers []string
	// for _, rawNumber := range values {
	// 	fmt.Println(string(rawNumber))

	// 	number, _ := strconv.Atoi(string(rawNumber))
	// 	// if err != nil {
	// 	// 	fmt.Println(err)

	// 	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid number format"})
	// 	// }
	// 	numbers = append(numbers, number)
	// }

	fmt.Println(values)

	res, err := s.castSrv.CalValue(values)
	if err != nil {
		logs.Error(err)
		return c.JSONPretty(http.StatusOK, res, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}
