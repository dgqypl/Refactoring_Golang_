package refactoring

import "math"

type Order struct {
	Quantity  int
	ItemPrice int
}

func (o *Order) Price() float64 {
	return o.basePrice() -
		o.quantityDiscount() +
		o.shipping()
}

func (o *Order) basePrice() float64 {
	return float64(o.Quantity * o.ItemPrice)
}
func (o *Order) quantityDiscount() float64 {
	return math.Max(0, float64(o.Quantity-500)) * float64(o.ItemPrice) * 0.05
}
func (o *Order) shipping() float64 {
	return math.Min((o.basePrice())*0.1, 100)
}
