package refactoring

type Driver struct {
	NumberOfLateDeliveries int
}

func Rating(driver Driver) int {
	if driver.NumberOfLateDeliveries > 5 {
		return 2
	}
	return 1
}
