package main_test

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Cashier represents the automatic cashier system.
type Cashier struct {
	Notes map[float64]int
}

// NewCashier initializes a new cashier instance.
func NewCashier() *Cashier {
	return &Cashier{
		Notes: map[float64]int{
			1000: 10,
			500:  20,
			100:  15,
			50:   20,
			20:   30,
			10:   20,
			5:    20,
			1:    20,
			0.25: 50,
		},
	}
}

// ProcessTransaction calculates the change and displays the remaining notes/coins.
func (c *Cashier) ProcessTransaction(productPrice, customerPaid float64, customerNotes map[float64]int) map[float64]int {
	// Add the customer's notes to the cashier's desk
	for note, numNotes := range customerNotes {
		c.Notes[note] += numNotes
	}

	// Calculate change
	change := customerPaid - productPrice
	changeNotes := make(map[float64]int)

	for note, limit := range c.Notes {
		if change >= note {
			numNotes := int(change / note)
			if numNotes > limit {
				numNotes = limit
			}

			changeNotes[note] = numNotes
			change -= float64(numNotes) * note
			c.Notes[note] -= numNotes
		}
	}

	// Display remaining notes/coins
	fmt.Println("\nRemaining Notes/Coins:")
	for note, numNotes := range c.Notes {
		fmt.Printf("$%.2f: %d notes/coins\n", note, numNotes)
	}

	return changeNotes
}

func main_test() {
	e := echo.New()

	// Create a cashier instance
	cashier := NewCashier()

	e.POST("/processTransaction", func(c echo.Context) error {
		// Parse request data
		productPrice, err := strconv.ParseFloat(c.FormValue("productPrice"), 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid product price")
		}

		customerPaid, err := strconv.ParseFloat(c.FormValue("customerPaid"), 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid customer paid amount")
		}

		// Parse customer notes from the request form
		customerNotes := make(map[float64]int)
		for _, denomination := range []float64{1000, 500, 100, 50, 20, 10, 5, 1, 0.25} {
			numNotes, err := strconv.Atoi(c.FormValue(fmt.Sprintf("note_%s", strconv.FormatFloat(denomination, 'f', -1, 64))))
			if err != nil {
				return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid number of %s notes", strconv.FormatFloat(denomination, 'f', -1, 64)))
			}
			customerNotes[denomination] = numNotes
		}

		// Process the transaction
		change := cashier.ProcessTransaction(productPrice, customerPaid, customerNotes)

		// Return the change as JSON
		return c.JSON(http.StatusOK, change)
	})

	e.Start(":8080")
}
