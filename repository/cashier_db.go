package repository

import "github.com/jmoiron/sqlx"

type cashierRepositoryDB struct {
	db *sqlx.DB
}

func NewCashierRepositoryDB(db *sqlx.DB) CashierRepository {
	return cashierRepositoryDB{db: db}
}
func (r cashierRepositoryDB) GetAll() ([]Cashier, error) {
	cashier := []Cashier{}
	query := "select * from Cashier"
	err := r.db.Select(&cashier, query)
	if err != nil {
		return nil, err
	}
	return cashier, nil
}

func (r cashierRepositoryDB) GetById(id int) (*Cashier, error) {
	cashier := Cashier{}
	query := "select * from Cashier where id =?"
	err := r.db.Get(&cashier, query, id)
	if err != nil {
		return nil, err
	}
	return &cashier, nil
}
