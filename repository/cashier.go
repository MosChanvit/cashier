package repository

type Cashier struct {
	ID        int    `db:"id"`
	IdCashier string `db:"id_cashier"`
	Name      string `db:"name"`
	C1000     int    `db:"c1000"`
	C500      int    `db:"c500"`
	C100      int    `db:"c100"`
	C50       int    `db:"c50"`
	C20       int    `db:"c20"`
	C10       int    `db:"c10"`
	C5        int    `db:"c5"`
	C1        int    `db:"c1"`
	C025      int    `db:"c025"`
}

type ShoppingList struct {
	ID             int     `db:"id"`
	C1000          int     `db:"c1000"`
	C500           int     `db:"c500"`
	C100           int     `db:"c100"`
	C50            int     `db:"c50"`
	C20            int     `db:"c20"`
	C10            int     `db:"c10"`
	C5             int     `db:"c5"`
	C1             int     `db:"c1"`
	C025           int     `db:"c025"`
	ProductPrice   float64 `json:"product_price"`
	CustomerPaid   float64 `json:"customer_paid"`
	CustomerChange float64 `json:"customer_change"`
	CashierId      int     `json:"cashier_id"`
}

//go:generate mockgen -destination=../mock/mock_repository/mock_account_repository.go bank/repository AccountRepository
type CashierRepository interface {
	// Create(Account) (*Account, error)
	GetAll() ([]Cashier, error)
	GetByIdCashier(string) (*Cashier, error)
	Create(Cashier) (*Cashier, error)
	Update(Cashier) (*Cashier, error)
	AddShoppingList(ShoppingList) (*ShoppingList, error)
}
