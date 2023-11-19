package service

type NewCashierRequest struct {
	IdCashier string `json:"id_cashier"`
	Name      string `json:"name"`
	C1000     int    `json:"c1000"`
	C500      int    `json:"c500"`
	C100      int    `json:"c100"`
	C50       int    `json:"c50"`
	C20       int    `json:"c20"`
	C10       int    `json:"c10"`
	C5        int    `json:"c5"`
	C1        int    `json:"c1"`
	C025      int    `json:"c025"`
}

type CashierResponse struct {
	IdCashier string `json:"id_cashier"`
	Name      string `json:"name"`
	C1000     int    `json:"c1000"`
	C500      int    `json:"c500"`
	C100      int    `json:"c100"`
	C50       int    `json:"c50"`
	C20       int    `json:"c20"`
	C10       int    `json:"c10"`
	C5        int    `json:"c5"`
	C1        int    `json:"c1"`
	C025      int    `json:"c025"`
}

type ProcessTransactionRequest struct {
	ID           int     `json:"id"`
	IdCashier    string  `json:"id_cashier"`
	C1000        int     `json:"c1000"`
	C500         int     `json:"c500"`
	C100         int     `json:"c100"`
	C50          int     `json:"c50"`
	C20          int     `json:"c20"`
	C10          int     `json:"c10"`
	C5           int     `json:"c5"`
	C1           int     `json:"c1"`
	C025         int     `json:"c025"`
	ProductPrice float64 `json:"product_price"`
	CustomerPaid float64 `json:"customer_paid"`
}
type ProcessTransactionResponse struct {
	C1000          int     `json:"c1000"`
	C500           int     `json:"c500"`
	C100           int     `json:"c100"`
	C50            int     `json:"c50"`
	C20            int     `json:"c20"`
	C10            int     `json:"c10"`
	C5             int     `json:"c5"`
	C1             int     `json:"c1"`
	C025           int     `json:"c025"`
	CustomerChange float64 `json:"customer_change"`
}

//go:generate mockgen -destination=../mock/mock_service/mock_account_service.go bank/service AccountService
type CashierService interface {
	GetCashiers() ([]CashierResponse, error)
	GetCashier(int) (*CashierResponse, error)
	NewCashier(request NewCashierRequest) (*CashierResponse, error)
	ProcessTransaction(request ProcessTransactionRequest) (*ProcessTransactionResponse, error)
}
