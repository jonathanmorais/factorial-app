package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Num struct {
	Number *big.Int `json:"number"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/api/{num}", getFactorial).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Tudo Suave")

}

func getFactorial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //pega variaveis "path parameter"
	w.WriteHeader(http.StatusOK)
	number, _ := vars["num"]
	_, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	} else if err == nil {
		w.Write([]byte(fmt.Sprint("Is Empty")))
	}
	n, _ := strconv.ParseInt(number, 10, 64)
	calcFactorial := (factorial(big.NewInt(n)))
	w.Write([]byte(fmt.Sprintf("%v", calcFactorial)))
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}
