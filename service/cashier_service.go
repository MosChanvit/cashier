package service

import (
	"cashier/errs"
	"cashier/logs"
	"cashier/repository"
	"database/sql"
	"fmt"
	"strconv"
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
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	Cashier := CashierResponse{
		IdCashier: cashier.IdCashier,
		Name:      cashier.Name,
		C1000:     cashier.C1000,
		C500:      cashier.C500,
		C100:      cashier.C100,
		C50:       cashier.C50,
		C20:       cashier.C20,
		C10:       cashier.C10,
		C5:        cashier.C5,
		C1:        cashier.C1,
		C025:      cashier.C025,
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
			IdCashier: cashier.IdCashier,
			Name:      cashier.Name,
			C1000:     cashier.C1000,
			C500:      cashier.C500,
			C100:      cashier.C100,
			C50:       cashier.C50,
			C20:       cashier.C20,
			C10:       cashier.C10,
			C5:        cashier.C5,
			C1:        cashier.C1,
			C025:      cashier.C025,
		}
		cashiersResponses = append(cashiersResponses, cashierResponse)
	}

	return cashiersResponses, nil
}

func (s cashierService) NewCashier(request NewCashierRequest) (*CashierResponse, error) {
	//Validate
	if request.IdCashier == "" {
		return nil, errs.NewValidationError("id_cashier not null")
	}
	if request.Name == "" {
		return nil, errs.NewValidationError("name not null")
	}

	cashier := repository.Cashier{
		IdCashier: request.IdCashier,
		Name:      request.Name,
		C1000:     request.C1000,
		C500:      request.C500,
		C100:      request.C100,
		C50:       request.C50,
		C20:       request.C20,
		C10:       request.C10,
		C5:        request.C5,
		C1:        request.C1,
		C025:      request.C025,
	}

	newCas, err := s.casRepo.Create(cashier)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := CashierResponse{
		IdCashier: newCas.IdCashier,
		Name:      newCas.Name,
		C1000:     newCas.C1000,
		C500:      newCas.C500,
		C100:      newCas.C100,
		C50:       newCas.C50,
		C20:       newCas.C20,
		C10:       newCas.C10,
		C5:        newCas.C5,
		C1:        newCas.C1,
		C025:      newCas.C025,
	}

	return &response, nil
}

func (s cashierService) ProcessTransaction(req ProcessTransactionRequest) (*ProcessTransactionResponse, error) {
	// c := context.Background()
	cashier, err := s.casRepo.GetById(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	Notes := map[float64]int{
		1000: cashier.C1000,
		500:  cashier.C500,
		100:  cashier.C100,
		50:   cashier.C50,
		20:   cashier.C20,
		10:   cashier.C10,
		5:    cashier.C5,
		1:    cashier.C5,
		0.25: cashier.C025,
	}

	// return changeNotes

	customerNotes := make(map[float64]int)
	for _, denomination := range []float64{1000, 500, 100, 50, 20, 10, 5, 1, 0.25} {
		numNotes, err := strconv.Atoi(fmt.Sprintf("note_%s", strconv.FormatFloat(denomination, 'f', -1, 64)))
		if err != nil {
			return nil, nil
		}
		customerNotes[denomination] = numNotes
	}

	// Process the transaction

	// Add the customer's notes to the cashier's desk
	for note, numNotes := range customerNotes {
		Notes[note] += numNotes
	}

	// Calculate change
	change := req.CustomerPaid - req.ProductPrice
	changeNotes := make(map[float64]int)

	for note, limit := range Notes {
		if change >= note {
			numNotes := int(change / note)
			if numNotes > limit {
				numNotes = limit
			}

			changeNotes[note] = numNotes
			change -= float64(numNotes) * note
			Notes[note] -= numNotes
		}
	}

	// Display remaining notes/coins
	fmt.Println("\nRemaining Notes/Coins:")
	for note, numNotes := range Notes {
		fmt.Printf("$%.2f: %d notes/coins\n", note, numNotes)
	}

	//Validate
	response := ProcessTransactionResponse{
		CustomerChange: change,
	}
	return &response, nil
}

// func ProcessTransaction(productPrice, customerPaid float64, customerNotes map[float64]int) map[float64]int {
// 	// Add the customer's notes to the cashier's desk
// 	for note, numNotes := range customerNotes {
// 		c.Notes[note] += numNotes
// 	}

// 	// Calculate change
// 	change := customerPaid - productPrice
// 	changeNotes := make(map[float64]int)

// 	for note, limit := range c.Notes {
// 		if change >= note {
// 			numNotes := int(change / note)
// 			if numNotes > limit {
// 				numNotes = limit
// 			}

// 			changeNotes[note] = numNotes
// 			change -= float64(numNotes) * note
// 			c.Notes[note] -= numNotes
// 		}
// 	}

// 	// Display remaining notes/coins
// 	fmt.Println("\nRemaining Notes/Coins:")
// 	for note, numNotes := range c.Notes {
// 		fmt.Printf("$%.2f: %d notes/coins\n", note, numNotes)
// 	}

// 	return changeNotes
// }
