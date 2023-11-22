package service

import (
	"cashier/logs"
	"cashier/repository"
	"database/sql"
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

func (s cashierService) GetCashier(name string) (*CashierResponse, error) {
	cashier, err := s.casRepo.GetByNameCashier(name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	Cashier := CashierAllResponse{
		Name:  cashier.Name,
		C1000: cashier.C1000,
		C500:  cashier.C500,
		C100:  cashier.C100,
		C50:   cashier.C50,
		C20:   cashier.C20,
		C10:   cashier.C10,
		C5:    cashier.C5,
		C1:    cashier.C1,
		C025:  cashier.C025,
	}

	transactionRes, err := s.casRepo.GetAllShoppingList(cashier.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	transactionResponseAll := []TransactionResponse{}
	for _, v := range transactionRes {
		transactionResponse := TransactionResponse{
			C1000:          v.C1000,
			C500:           v.C500,
			C100:           v.C100,
			C50:            v.C50,
			C20:            v.C20,
			C10:            v.C10,
			C5:             v.C5,
			C1:             v.C1,
			C025:           v.C025,
			ProductPrice:   v.ProductPrice,
			CustomerPaid:   v.CustomerPaid,
			CustomerChange: v.CustomerChange,
		}
		transactionResponseAll = append(transactionResponseAll, transactionResponse)
	}

	cashierResponse := CashierResponse{
		Cashier:             Cashier,
		TransactionResponse: transactionResponseAll,
	}

	return &cashierResponse, nil
}

func (s cashierService) GetCashiers() ([]CashierAllResponse, error) {

	cashiers, err := s.casRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	cashiersResponses := []CashierAllResponse{}
	for _, cashier := range cashiers {
		cashierResponse := CashierAllResponse{
			Name:  cashier.Name,
			C1000: cashier.C1000,
			C500:  cashier.C500,
			C100:  cashier.C100,
			C50:   cashier.C50,
			C20:   cashier.C20,
			C10:   cashier.C10,
			C5:    cashier.C5,
			C1:    cashier.C1,
			C025:  cashier.C025,
		}
		cashiersResponses = append(cashiersResponses, cashierResponse)
	}

	return cashiersResponses, nil
}

func (s cashierService) NewCashier(request NewCashierRequest) (*CashierAllResponse, error) {

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
		return nil, err
	}

	response := CashierAllResponse{
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

func (s cashierService) ProcessTransaction(req ProcessTransactionRequest) (*ProcessTransactionResponse, error) {
	// c := context.Background()
	cashier, err := s.casRepo.GetByNameCashier(req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	change := req.CustomerPaid - req.ProductPrice

	changeNotes := make(map[float64]int)

	cashierSum := repository.Cashier{
		Name:  cashier.Name,
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

	customerChange := req.CustomerPaid - req.ProductPrice

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

	cashierSum.C1000 -= changeNotes[1000]
	cashierSum.C500 -= changeNotes[500]
	cashierSum.C100 -= changeNotes[100]
	cashierSum.C50 -= changeNotes[50]
	cashierSum.C20 -= changeNotes[20]
	cashierSum.C10 -= changeNotes[10]
	cashierSum.C5 -= changeNotes[5]
	cashierSum.C1 -= changeNotes[1]
	cashierSum.C025 -= changeNotes[0.25]

	///// ShoppingList
	shoppingList := repository.ShoppingList{
		C1000:          changeNotes[1000],
		C500:           changeNotes[500],
		C100:           changeNotes[100],
		C50:            changeNotes[50],
		C20:            changeNotes[20],
		C10:            changeNotes[10],
		C5:             changeNotes[5],
		C1:             changeNotes[1],
		C025:           changeNotes[0.25],
		ProductPrice:   req.ProductPrice,
		CustomerPaid:   req.CustomerPaid,
		CustomerChange: customerChange,
		CashierId:      cashier.ID,
	}

	_, err = s.casRepo.AddShoppingList(shoppingList)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	/////Update
	_, err = s.casRepo.Update(cashierSum)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	// changeNotes
	return &response, nil
}

func (s cashierService) CalXYZ(setStr []string) string {
	res := ""
	// var numbers []int
	for i, rawNumber := range setStr {
		// fmt.Println(string(rawNumber))

		_, err := strconv.Atoi(string(rawNumber))
		if err != nil {
			fmt.Println(string(rawNumber), " has possible values ")
			res += string(rawNumber) + " has possible values \n"

			var numbers []int
			i1, errBef := strconv.Atoi(string(setStr[i-1]))
			i2, errAft := strconv.Atoi(string(setStr[i+1]))

			if errBef == nil && errAft == nil {
				for i := i1 + 1; i < i2; i++ {
					numbers = append(numbers, i)
				}
			} else if errBef == nil && errAft != nil {
				i3, _ := strconv.Atoi(string(setStr[i+2]))
				for i := i1 + 1; i < i3-1; i++ {
					numbers = append(numbers, i)
				}
			} else if errBef != nil && errAft == nil {

				i3, _ := strconv.Atoi(string(setStr[i-2]))

				for i := i3 + 2; i < i2; i++ {
					numbers = append(numbers, i)
				}
			}

			for _, v := range numbers {
				num := strconv.Itoa(v)
				res += num + ", "
			}
			res += "\n"
		}

	}
	return res
}
