package refactoring

import (
	"encoding/json"
	"fmt"
	"math"
)

type Reading struct {
	Customer string `json:"customer"`
	Quantity int    `json:"quantity"`
	Month    int    `json:"month"`
	Year     int    `json:"year"`
}

func (r *Reading) BaseRate() int {
	// do some logic
	return 0
}

func (r *Reading) BaseCharge() int {
	return r.BaseRate() * r.Quantity
}

func (r *Reading) TaxThreshold() int {
	// do some logic
	return 0
}

func (r *Reading) TaxableCharge() float64 {
	return math.Max(0, float64(r.BaseCharge()-r.TaxThreshold()))
}

func getReading() Reading {
	data := `{"Customer": "ivan", "Quantity": 10, "Month": 5, "Year": 2017}`
	var reading Reading
	json.Unmarshal([]byte(data), &reading)
	fmt.Println(reading)
	return reading
}

func Invoker1() {
	reading := getReading()
	baseCharge := reading.BaseCharge()
	fmt.Println(baseCharge)
}

func Invoker2() {
	reading := getReading()
	taxableCharge := reading.TaxableCharge()
	fmt.Println(taxableCharge)
}
