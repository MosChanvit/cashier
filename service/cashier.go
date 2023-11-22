package service

type NewCashierRequest struct {
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

type CashierAllResponse struct {
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

type ProcessTransactionRequest struct {
	Name         string  `json:"name"`
	ID           int     `json:"id"`
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
	C1000          int     `json:"1000"`
	C500           int     `json:"500"`
	C100           int     `json:"100"`
	C50            int     `json:"50"`
	C20            int     `json:"20"`
	C10            int     `json:"10"`
	C5             int     `json:"5"`
	C1             int     `json:"1"`
	C025           int     `json:"0.25"`
	CustomerChange float64 `json:"customer_change"`
}

type CashierResponse struct {
	Cashier             CashierAllResponse    `json:"cashier"`
	TransactionResponse []TransactionResponse `json:"transaction_response"`
}

type TransactionResponse struct {
	C1000          int     `json:"1000"`
	C500           int     `json:"500"`
	C100           int     `json:"100"`
	C50            int     `json:"50"`
	C20            int     `json:"20"`
	C10            int     `json:"10"`
	C5             int     `json:"5"`
	C1             int     `json:"1"`
	C025           int     `json:"0.25"`
	ProductPrice   float64 `json:"product_price"`
	CustomerPaid   float64 `json:"customer_paid"`
	CustomerChange float64 `json:"customer_change"`
}

type CashierService interface {
	GetCashiers() ([]CashierAllResponse, error)
	GetCashier(string) (*CashierResponse, error)
	NewCashier(request NewCashierRequest) (*CashierAllResponse, error)
	ProcessTransaction(request ProcessTransactionRequest) (*ProcessTransactionResponse, error)
	CalXYZ(num []string) string
}
