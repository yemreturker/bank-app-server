package main

import (
	"bank-app-server/src/controllers"
	"bank-app-server/src/db"
	"bank-app-server/src/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Veritabanına bağlan
	db.ConnectDatabase()

	// Router oluştur
	r := mux.NewRouter()

	// Endpoint’ler
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.Handle("/password/reset", middlewares.AuthMiddleware(http.HandlerFunc(controllers.PasswordResetHandler))).Methods("POST")
	r.Handle("/token/refresh", middlewares.AuthMiddleware(http.HandlerFunc(controllers.TokenRefreshHandler))).Methods("POST")

	r.Handle("/balance/{accountNumber}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.BalanceHandler))).Methods("GET")
	r.Handle("/transfer", middlewares.AuthMiddleware(http.HandlerFunc(controllers.TransferHandler))).Methods("POST")
	r.Handle("/account", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateAccountHandler))).Methods("POST")
	r.Handle("/accounts", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ListAccountsHandler))).Methods("GET")
	r.Handle("/transactions", middlewares.AuthMiddleware(http.HandlerFunc(controllers.TransactionHistoryHandler))).Methods("GET")
	r.Handle("/account/summary", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AccountSummaryHandler))).Methods("GET")
	r.Handle("/account/deposit", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DepositHandler))).Methods("POST")
	r.Handle("/account/withdraw", middlewares.AuthMiddleware(http.HandlerFunc(controllers.WithdrawHandler))).Methods("POST")

	// Server başlat
	log.Println("Server çalışıyor: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}