package function

import (
	"fmt"
	"time"
)

type Invoice struct {
	Customer string
	Orders   []Order
	DueDate  time.Time
}

type Order struct {
	Amount int
}

func PrintOwing(invoice Invoice) {
	outstanding := 0
	fmt.Println("***********************")
	fmt.Println("**** Customer Owes ****")
	fmt.Println("***********************")

	// calculate outstanding
	for _, v := range invoice.Orders {
		outstanding += v.Amount
	}

	// record due date
	today := time.Now()
	invoice.DueDate = today.AddDate(0, 0, 30)

	// print details
	fmt.Println("name: ", invoice.Customer)
	fmt.Println("amount: ", outstanding)
	fmt.Println("due: ", invoice.DueDate.Format("2006-01-02"))
}
