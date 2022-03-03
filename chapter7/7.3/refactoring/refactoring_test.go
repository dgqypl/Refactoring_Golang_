package refactoring

import (
	"testing"
)

func TestSomeInvoker(t *testing.T) {
	var orders = []Order{
		{Priority: Priority{value: "normal"}},
		{Priority: Priority{value: "rush"}},
		{Priority: Priority{value: "normal"}},
		{Priority: Priority{value: "low"}},
		{Priority: Priority{value: "high"}},
	}
	SomeInvoker(orders)
}
