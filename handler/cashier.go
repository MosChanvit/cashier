package handler

import (
	"cashier/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type cashierHandler struct {
	castSrv service.CashierService
}

func NewCashierHandler(castSrv service.CashierService) cashierHandler {
	return cashierHandler{castSrv: castSrv}
}

func (h cashierHandler) GetCashiers(w http.ResponseWriter, r *http.Request) {
	cashiers, err := h.castSrv.GetCashiers()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(cashiers)
}

// func (h cashierHandler) GetCashier(w http.ResponseWriter, r *http.Request) {
// 	cashierID, _ := strconv.Atoi(mux.Vars(r)["cashierID"])

// 	cashier, err := h.castSrv.GetCashier(cashierID)
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(cashier)
// }

// TransferConfirm implements handlers.API
func (s cashierHandler) GetCashier(c echo.Context) error {

	id := c.Param("id")
	idReq, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	res, err := s.castSrv.GetCashier(idReq)
	if err != nil {
		return c.JSONPretty(http.StatusOK, res, "")
	}

	return c.JSONPretty(http.StatusOK, res, "")

}
