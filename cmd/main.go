package main

import (
	"cashier/handler"
	"cashier/logs"
	"cashier/repository"
	"cashier/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()

	db := initDatabase()
	createTable(db)

	cashierRepositoryDB := repository.NewCashierRepositoryDB(db)
	cashierService := service.NewCashierRepoService(cashierRepositoryDB)
	cashierHandler := handler.NewCashierHandler(cashierService)

	router := mux.NewRouter()

	router.HandleFunc("/cashiers", cashierHandler.GetCashiers).Methods(http.MethodGet)
	router.HandleFunc("/cashier/{cashier:[0-9]+}", cashierHandler.GetCashier).Methods(http.MethodGet)

	logs.Info("Cashier service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
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
