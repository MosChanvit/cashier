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
	query := "select * from cashier"
	err := r.db.Select(&cashier, query)
	if err != nil {
		return nil, err
	}
	return cashier, nil
}

func (r cashierRepositoryDB) GetById(id int) (*Cashier, error) {
	cashier := Cashier{}
	query := "select * from cashier where id =?"
	err := r.db.Get(&cashier, query, id)
	if err != nil {
		return nil, err
	}
	return &cashier, nil
}

func (r cashierRepositoryDB) Create(cas Cashier) (*Cashier, error) {
	query := `INSERT INTO cashier.cashier
	(name, c1000, c500, c100, c50, c20, c10, c5, c1, c025)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
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
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	cas.ID = int(id)

	return &cas, nil
}
