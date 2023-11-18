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

func (s cashierService) NewCashier(request NewCashierRequest) (*CashierResponse, error) {
	//Validate
	if request.Name == "" {
		return nil, errs.NewValidationError("Name not null")
	}
	// if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
	// 	return nil, errs.NewValidationError("account type should be saving or checking")
	// }

	cashier := repository.Cashier{
		Name:  request.Name,
		C1000: request.C1000,
		C500:  request.C500,
		C100:  request.C100,
		C50:   request.C50,
		C20:   request.C20,
		C10:   request.C10,
		C5:    request.C5,
		C1:    request.C1,
		C025:  request.C025,
	}

	newCas, err := s.casRepo.Create(cashier)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := CashierResponse{
		ID:    newCas.ID,
		Name:  newCas.Name,
		C1000: newCas.C1000,
		C500:  newCas.C500,
		C100:  newCas.C100,
		C50:   newCas.C50,
		C20:   newCas.C20,
		C10:   newCas.C10,
		C5:    newCas.C5,
		C1:    newCas.C1,
		C025:  newCas.C025,
	}

	return &response, nil
}

func (s cashierService) ProcessTransaction(request ProcessTransactionRequest) (*ProcessTransactionResponse, error) {
	//Validate
	response := ProcessTransactionResponse{
		CustomerChange: 0,
	}
	return &response, nil
}
