package rpwo

import (
	"testing"
)

func TestSomeInvoker(t *testing.T) {
	var orders = []Order{
		{Priority: "normal"},
		{Priority: "rush"},
		{Priority: "normal"},
		{Priority: "low"},
		{Priority: "high"},
	}
	SomeInvoker(orders)
}
