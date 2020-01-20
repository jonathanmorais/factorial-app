package main

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

type Num struct {
	Number *big.Int `json:"number"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/api/facto", getFactorial).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
	http.Handle("/", r)

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Tudo Suave")

}

func getFactorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var number Num
	n := big.NewInt(Num)
	_ = json.NewDecoder(r.Body).Decode(&number)
	calcFactorial := (factorial(big.NewInt(number)))
	json.NewEncoder(w).Encode(calcFactorial)
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}
