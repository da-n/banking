package app

import (
	"log"
	"net/http"

	"github.com/da-n/banking/domain"
	"github.com/da-n/banking/service"
	"github.com/gorilla/mux"
)

// Start bootstraps the mux and starts server.
func Start() {

	router := mux.NewRouter()

	// Wiring.
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Define routes.
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Start server.
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
