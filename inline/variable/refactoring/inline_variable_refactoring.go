package inlineVariableRefactoring

type Order struct {
	BasePrice int
}

func SomeFunction(order Order) bool {
	return order.BasePrice > 1000
}
