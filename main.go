package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Loan struct {
	Principal float64 `json:"principal"`
	Rate      float64 `json:"rate"`
	Tenure    int     `json:"tenure"`
}

func calculateEMI(p, r float64, n int) float64 {
	monthlyRate := r / (12 * 100)
	emi := (p * monthlyRate * (1 + monthlyRate) * float64(n)) /
		((1 + monthlyRate) * float64(n) - 1)
	return emi
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var loan Loan
	_ = json.NewDecoder(r.Body).Decode(&loan)
	emi := calculateEMI(loan.Principal, loan.Rate, loan.Tenure)
	json.NewEncoder(w).Encode(map[string]float64{"emi": emi})
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
