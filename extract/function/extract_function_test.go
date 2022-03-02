package function

import (
	"testing"
	"time"
)

func TestPrintOwing(t *testing.T) {
	var orders []Order
	invoice := Invoice{
		Customer: "",
		Orders:   orders,
		DueDate:  time.Time{},
	}
	PrintOwing(invoice)
}
