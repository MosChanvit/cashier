package service

type NewCashierRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type CashierResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	C1000 int    `json:"c1000"`
	C500  int    `json:"c500"`
	C100  int    `json:"c100"`
	C50   int    `json:"c50"`
	C20   int    `json:"c20"`
	C10   int    `json:"c10"`
	C5    int    `json:"c5"`
	C1    int    `json:"c1"`
	C025  int    `json:"c025"`
}

//go:generate mockgen -destination=../mock/mock_service/mock_account_service.go bank/service AccountService
type CashierService interface {
	GetCashiers() ([]CashierResponse, error)
	GetCashier(int) (*CashierResponse, error)
}
