package service

import (
	"cashier/errs"
	"cashier/logs"
	"cashier/repository"
	"database/sql"
)

type cashierService struct {
	casRepo repository.CashierRepository
}

func NewCashierRepoService(casRepo repository.CashierRepository) CashierService {
	return cashierService{casRepo: casRepo}
}

func (s cashierService) GetCashier(id int) (*CashierResponse, error) {
	cashier, err := s.casRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("cashier not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	Cashier := CashierResponse{
		ID:   cashier.ID,
		Name: cashier.Name,
	}

	return &Cashier, nil
}
func (s cashierService) GetCashiers() ([]CashierResponse, error) {

	cashiers, err := s.casRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	cashiersResponses := []CashierResponse{}
	for _, cashier := range cashiers {
		cashierResponse := CashierResponse{
			ID:   cashier.ID,
			Name: cashier.Name,
		}
		cashiersResponses = append(cashiersResponses, cashierResponse)
	}

	return cashiersResponses, nil
}
