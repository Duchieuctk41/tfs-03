package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type Charge struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptMail"`
	Description  string `json:"description"`
	Source       string `json:"source"`
}

func main() {
	// This is your real test secret API key.
	stripe.Key = "sk_test_51JcxcTGFEhJjjEOsJwqy6uUp6ZapjeRJDCEb1K3NnJ7ehdNum0jFBUSDftKrjKRRRtdShsrDR5KQDzzUmlEiPy2X00CxuuyLiF"

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/checkout", handleCreatePaymentIntent).Methods("POST")

	addr := "localhost:3000"
	log.Printf("Listening on %s ...", addr)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	var data Charge
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	params := &stripe.ChargeParams{
		Amount:       stripe.Int64(data.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(data.Description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(data.ReceiptEmail),
	}

	params.SetSource(data.Source)
	_, err := charge.New(params)
	if err != nil {
		fmt.Fprintf(w, "cannot pay")
		return
	}

	fmt.Fprintf(w, "success pay")
	return
}
