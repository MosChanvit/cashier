package repository

import (
	"cashier/logs"
	"log"

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

func (r cashierRepositoryDB) GetByNameCashier(name string) (*Cashier, error) {
	cashier := Cashier{}
	query := "select * from cashier where name = ?"
	err := r.db.Get(&cashier, query, name)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &cashier, nil
}

func (r cashierRepositoryDB) Create(cas Cashier) (*Cashier, error) {
	query := `INSERT INTO cashier.cashier
	(name, c1000, c500, c100, c50, c20, c10, c5, c1, c025)
	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	result, err := r.db.Exec(
		query,
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
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
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
	SET  c1000 = ?, c500 = ?, c100 = ?, c50 = ?, c20 = ?, c10 = ?, c5 = ?, c1 = ?, c025 = ?
	WHERE name = ? ;`
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
		cas.Name,
	)
	log.Println("up")
	log.Println(cas)
	log.Println(result)

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

func (r cashierRepositoryDB) GetAllShoppingList(id int) ([]ShoppingList, error) {
	shoppingList := []ShoppingList{}
	query := "select * from shopping_list where cashier_id = ?"
	err := r.db.Select(&shoppingList, query, id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return shoppingList, nil
}
