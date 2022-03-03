package rpwo

import "fmt"

type Order struct {
	Priority string
}

func SomeInvoker(orders []Order) {
	for i, v := range orders {
		if v.Priority == "high" || v.Priority == "rush" {
			// do some logic
			fmt.Println(i, v)
		}
	}
}
