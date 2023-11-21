package repository

import (
	"cashier/logs"

	"github.com/jmoiron/sqlx"
)

type cashierRepositoryDB struct {
	db *sqlx.DB
}

func NewCashierRepositoryDB(db *sqlx.DB) CashierRepository {
	return cashierRepositoryDB{db: db}
}
func (r cashierRepositoryDB) GetAll() ([]Cashier, error) {
	cashier := []Cashier{}
	query := "select * from cashier"
	err := r.db.Select(&cashier, query)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return cashier, nil
}

func (r cashierRepositoryDB) GetByIdCashier(idCashier string) (*Cashier, error) {
	cashier := Cashier{}
	query := "select * from cashier where id_cashier = ?"
	err := r.db.Get(&cashier, query, idCashier)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &cashier, nil
}

func (r cashierRepositoryDB) Create(cas Cashier) (*Cashier, error) {
	query := `INSERT INTO cashier.cashier
	(id_cashier, name, c1000, c500, c100, c50, c20, c10, c5, c1, c025)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	result, err := r.db.Exec(
		query,
		cas.IdCashier,
		cas.Name,
		cas.C1000,
		cas.C500,
		cas.C100,
		cas.C50,
		cas.C20,
		cas.C10,
		cas.C5,
		cas.C1,
		cas.C025,
	)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	cas.ID = int(id)

	return &cas, nil
}

func (r cashierRepositoryDB) AddShoppingList(shop ShoppingList) (*ShoppingList, error) {
	query := `INSERT INTO cashier.shopping_list
	(c1000, c500, c100, c50, c20, c10, c5, c1, c025, product_price, customer_paid, customer_change, cashier_id)
	VALUES(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);`
	result, err := r.db.Exec(
		query,
		shop.C1000,
		shop.C500,
		shop.C100,
		shop.C50,
		shop.C20,
		shop.C10,
		shop.C5,
		shop.C1,
		shop.C025,
		shop.ProductPrice,
		shop.CustomerPaid,
		shop.CustomerChange,
		shop.CashierId,
	)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	shop.ID = int(id)

	return &shop, nil
}

func (r cashierRepositoryDB) Update(cas Cashier) (*Cashier, error) {
	query := `UPDATE cashier.cashier
	SET  c1000 = ?, c500 = ?, c100 = ?, c50 = ?, c20 = ?, c10 = ?, c5 = ?, c1 = ?, c025 = ?, id_cashier=''
	WHERE id_cashier = ? ;`
	result, err := r.db.Exec(
		query,
		cas.C1000,
		cas.C500,
		cas.C100,
		cas.C50,
		cas.C20,
		cas.C10,
		cas.C5,
		cas.C1,
		cas.C025,
		cas.IdCashier,
	)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	cas.ID = int(id)

	return &cas, nil
}
