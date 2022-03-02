// Package inlineFunction
// Motivation:
// Sometimes, come across a function in which the body is as clear as the name.
// So, refactor the body of the code into something that is just as clear as the name.
package inlineFunction

type Driver struct {
	NumberOfLateDeliveries int
}

func Rating(driver Driver) int {
	if moreThanFiveLateDeliveries(driver) {
		return 2
	}
	return 1
}

func moreThanFiveLateDeliveries(driver Driver) bool {
	return driver.NumberOfLateDeliveries > 5
}
