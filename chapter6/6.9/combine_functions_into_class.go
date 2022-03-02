package cfic

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

func getReading() Reading {
	data := `{"customer": "ivan", "quantity": 10, "month": 5, "year": 2017}`
	var reading Reading
	json.Unmarshal([]byte(data), &reading)
	return reading
}

func baseRate(month int, year int) int {
	// do some logic
	return 0
}

func Invoker1() {
	reading := getReading()
	baseCharge := baseRate(reading.Month, reading.Year) * reading.Quantity
	fmt.Println(baseCharge)
}

func Invoker2() {
	reading := getReading()
	base := baseRate(reading.Month, reading.Year) * reading.Quantity
	taxableCharge := math.Max(0, float64(base-taxThreshold(reading.Year)))
	fmt.Println(taxableCharge)
}

func taxThreshold(year int) int {
	// do some logic
	return 0
}

func Invoker3() {
	reading := getReading()
	basicChargeAmount := calculateBaseCharge(reading)
	fmt.Println(basicChargeAmount)
}

func calculateBaseCharge(reading Reading) int {
	return baseRate(reading.Month, reading.Year) * reading.Quantity
}
