// Package inlineVariable
// Motivation:
// Variables provide names for expressions within a function, and as such they are usually a Good Thing.
// But sometimes, the name doesn’t really communicate more than the expression itself.
package inlineVariable

type Order struct {
	BasePrice int
}

func SomeFunction(order Order) bool {
	basePrice := order.BasePrice
	return basePrice > 1000
}
