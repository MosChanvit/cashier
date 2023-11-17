package handler

import (
	"cashier/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type cashierHandler struct {
	custSrv service.CashierService
}

func NewCashierHandler(custSrv service.CashierService) cashierHandler {
	return cashierHandler{custSrv: custSrv}
}

func (h cashierHandler) GetCashiers(w http.ResponseWriter, r *http.Request) {
	cashiers, err := h.custSrv.GetCashiers()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(cashiers)
}

func (h cashierHandler) GetCashier(w http.ResponseWriter, r *http.Request) {
	cashierID, _ := strconv.Atoi(mux.Vars(r)["cashierID"])

	cashier, err := h.custSrv.GetCashier(cashierID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(cashier)
}
