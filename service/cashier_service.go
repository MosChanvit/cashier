package service

import (
	"cashier/errs"
	"cashier/logs"
	"cashier/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type cashierService struct {
	casRepo repository.CashierRepository
}

func NewCashierRepoService(casRepo repository.CashierRepository) CashierService {
	return cashierService{casRepo: casRepo}
}

func (s cashierService) GetCashier(IdCashier string) (*CashierResponse, error) {
	cashier, err := s.casRepo.GetByIdCashier(IdCashier)
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
		return nil, err
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
	cashier, err := s.casRepo.GetByIdCashier(req.IdCashier)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}
	log.Println("cashier")

	res2B, _ := json.Marshal(cashier)
	fmt.Println(string(res2B))

	change := req.CustomerPaid - req.ProductPrice
	changeNotes := make(map[float64]int)

	cashierSum := repository.Cashier{
		C1000: cashier.C1000 + req.C1000,
		C500:  cashier.C500 + req.C500,
		C100:  cashier.C100 + req.C100,
		C50:   cashier.C50 + req.C50,
		C20:   cashier.C20 + req.C20,
		C10:   cashier.C10 + req.C10,
		C5:    cashier.C5 + req.C5,
		C1:    cashier.C5 + req.C1,
		C025:  cashier.C025 + req.C025,
	}

	res2B, _ = json.Marshal(cashierSum)
	fmt.Println(string(res2B))

	coins := map[float64]int{
		1000: cashier.C1000 + req.C1000,
		500:  cashier.C500 + req.C500,
		100:  cashier.C100 + req.C100,
		50:   cashier.C50 + req.C50,
		20:   cashier.C20 + req.C20,
		10:   cashier.C10 + req.C10,
		5:    cashier.C5 + req.C5,
		1:    cashier.C5 + req.C1,
		0.25: cashier.C025 + req.C025,
	}

	var keys []float64
	for k := range coins {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	for _, k := range keys {
		log.Println("note", k, "limit", coins[k])
		if change >= k {
			numNotes := int(change / k)
			if numNotes > coins[k] {
				numNotes = coins[k]
			}

			changeNotes[k] = numNotes
			change -= float64(numNotes) * k
			coins[k] -= numNotes
		}
	}

	log.Println("changeNotes")
	log.Println(changeNotes)

	res2B, _ = json.Marshal(changeNotes)
	fmt.Println(string(res2B))
	// changeNotes

	customerChange := req.CustomerPaid - req.ProductPrice
	log.Println("cashcustomerChangeier")
	log.Println(cashier)
	log.Println(cashier)
	log.Println(customerChange)

	// //Validate
	response := ProcessTransactionResponse{
		CustomerChange: customerChange,
		C1000:          changeNotes[1000],
		C500:           changeNotes[500],
		C100:           changeNotes[100],
		C50:            changeNotes[50],
		C20:            changeNotes[20],
		C10:            changeNotes[10],
		C5:             changeNotes[5],
		C1:             changeNotes[1],
		C025:           changeNotes[0.25],
	}
	res2B, _ = json.Marshal(response)
	fmt.Println(string(res2B))

	log.Println("cashier +++")
	res2B, _ = json.Marshal(cashierSum)
	fmt.Println(string(res2B))

	cashierSum.C1000 -= changeNotes[1000]
	cashierSum.C500 -= changeNotes[500]
	cashierSum.C100 -= changeNotes[100]
	cashierSum.C50 -= changeNotes[50]
	cashierSum.C20 -= changeNotes[20]
	cashierSum.C10 -= changeNotes[10]
	cashierSum.C5 -= changeNotes[5]
	cashierSum.C1 -= changeNotes[1]
	cashierSum.C025 -= changeNotes[0.25]

	log.Println("cashierSum ---")

	res2B, _ = json.Marshal(cashierSum)
	fmt.Println(string(res2B))

	// changeNotes
	return &response, nil
}

func (s cashierService) CalValue(setStr []string) ([]string, error) {

	var numbers []int
	for _, rawNumber := range setStr {
		fmt.Println(string(rawNumber))

		int, err := strconv.Atoi(string(rawNumber))
		if err != nil {
			fmt.Println(err)

			// return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid number format"})
		}

		fmt.Println(string(rawNumber))

		numbers = append(numbers, int)
	}

	for i, v := range numbers {

		if v == 0 && numbers[i-1] != 0 && numbers[i+1] != 0 && numbers[i+2] != 0 {
			g1 := numbers[i+1] - numbers[i-1]
			g2 := numbers[i+2] - numbers[i+1]
			fmt.Println(g2-g1, v)
		}
	}

	// order[n+1] - order[n-1]
	// (order[n+2] - order[n+1]) - (order[n+1] - order[n-1])

	log.Println(numbers)
	return setStr, nil
	// return []int{}, nil

	// (order[8+2] - order[8+1]) - (order[8+1] - order[8-1])
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
