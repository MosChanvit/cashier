package main

import (
	"cashier/handler"
	"cashier/logs"
	"cashier/repository"
	"cashier/service"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	ctx := context.Background()

	db := initDatabase()
	createTable(db)

	cashierRepositoryDB := repository.NewCashierRepositoryDB(db)
	cashierService := service.NewCashierRepoService(cashierRepositoryDB)

	e := InitRouter(cashierService)

	run(ctx, e)
}

func InitRouter(CashierSvc service.CashierService) *echo.Echo {
	// init echo
	e := echo.New()

	handler := handler.NewCashierHandler(CashierSvc)
	e.GET("/health", healthCheck) // health check
	e.GET("/cashiers", handler.GetCashiers)
	e.GET("/cashier", handler.GetCashier)
	e.POST("/cashier", handler.NewCashier)
	e.POST("/pay", handler.ProcessTransaction)
	e.POST("/cal_xyz", handler.CalXYZ)

	return e
}

func run(ctx context.Context, e *echo.Echo) {
	serverPort := fmt.Sprintf(":%v", viper.GetString("app.port"))
	if err := e.Start(serverPort); err != nil {
		e.Logger.Fatal(ctx, "shutting down the server")
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	fmt.Println(dsn)

	fmt.Println(viper.GetString("db.driver"))
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func createTable(db *sqlx.DB) {
	// SQL statement to create the user table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cashier (
		id INT AUTO_INCREMENT PRIMARY KEY,
		id_cashier varchar(250)  NOT NULL,
		name varchar(250)  NOT NULL,
		c1000 INT(100) NOT NULL,
		c500 INT(100) NOT NULL,
		c100 INT(100) NOT NULL,
		c50 INT(100) NOT NULL,
		c20 INT(100) NOT NULL,
		c10 INT(100) NOT NULL,
		c5 INT(100) NOT NULL,
		c1 INT(100) NOT NULL,
		c025 INT(100) NOT NULL
	);`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	createTableSQL = `
	CREATE TABLE IF NOT EXISTS shopping_list (
		id INT AUTO_INCREMENT PRIMARY KEY,
		c1000 INT(100) NOT NULL,
		c500 INT(100) NOT NULL,
		c100 INT(100) NOT NULL,
		c50 INT(100) NOT NULL,
		c20 INT(100) NOT NULL,
		c10 INT(100) NOT NULL,
		c5 INT(100) NOT NULL,
		c1 INT(100) NOT NULL,
		c025 INT(100) NOT NULL,
		product_price FLOAT(24) NOT NULL,
		customer_paid FLOAT(24) NOT NULL,
		customer_change  FLOAT(24) NOT NULL,
		cashier_id INT NOT NULL
	);`
	// Execute the SQL statement
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	logs.Info("Table cashier & shopping_list Ready to use")
}

func healthCheck(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, echo.Map{"message": "Service is Running !!"}, "	")
}

// ///////////////////////
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
