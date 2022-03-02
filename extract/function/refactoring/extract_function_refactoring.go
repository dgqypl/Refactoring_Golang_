package extractFunctionRefactoring

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
	printBanner()
	outstanding := calculateOutstanding(invoice)
	recordDueDate(&invoice)
	printDetails(invoice, outstanding)
}

func printBanner() {
	fmt.Println("***********************")
	fmt.Println("**** Customer Owes ****")
	fmt.Println("***********************")
}

func recordDueDate(invoice *Invoice) {
	today := time.Now()
	invoice.DueDate = today.AddDate(0, 0, 30)
}

func calculateOutstanding(invoice Invoice) int {
	result := 0
	for _, v := range invoice.Orders {
		result += v.Amount
	}
	return result
}

func printDetails(invoice Invoice, outstanding int) {
	fmt.Println("name: ", invoice.Customer)
	fmt.Println("amount: ", outstanding)
	fmt.Println("due: ", invoice.DueDate.Format("2006-01-02"))
}
