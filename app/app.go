package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/da-n/banking/domain"
	"github.com/da-n/banking/service"
	"github.com/gorilla/mux"
)

func envCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

// Start bootstraps the mux and starts server.
func Start() {
	envCheck()

	router := mux.NewRouter()

	// Wiring.
	dbClient := getDbClient()
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)

	ah := AccountHandlers{service.NewAccountService(accountRepositoryDb)}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	// Define routes.
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", ah.MakeTransaction).Methods(http.MethodPost)

	// Start server.
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
