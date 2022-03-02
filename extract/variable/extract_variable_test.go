package variable

import (
	"testing"
)

func TestPrice(t *testing.T) {
	order := Order{
		Quantity:  50,
		ItemPrice: 20,
	}
	expected := 1100.0
	if ans := Price(order); ans != expected {
		t.Fatalf("%+v expected %f, but %f got",
			order, expected, ans)
	}
}
