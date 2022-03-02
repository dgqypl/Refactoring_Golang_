package variable

import "math"

type Order struct {
	Quantity  int
	ItemPrice int
}

func Price(order Order) float64 {
	return float64(order.Quantity*order.ItemPrice) -
		math.Max(0, float64(order.Quantity-500))*float64(order.ItemPrice)*0.05 +
		math.Min((float64(order.Quantity*order.ItemPrice))*0.1, 100)
}
