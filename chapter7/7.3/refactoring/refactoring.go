package refactoring

import "fmt"

type Order struct {
	Priority Priority
}

type Priority struct {
	value string
}

func NewPriority(value string) Priority {
	if includes(value) {
		return Priority{value: value}
	}
	panic("illegal value!")
}

func includes(value string) bool {
	for _, v := range legalValues() {
		if v == value {
			return true
		}
	}
	return false
}

func legalValues() []string {
	return []string{"low", "normal", "high", "rush"}
}

func (p *Priority) index() int {
	for i, v := range legalValues() {
		if p.value == v {
			return i
		}
	}
	return -1
}

func (p *Priority) Equals(other Priority) bool {
	return p.index() == other.index()
}

func (p *Priority) HigherThan(other Priority) bool {
	return p.index() > other.index()
}

func (p *Priority) LowerThan(other Priority) bool {
	return p.index() < other.index()
}

func SomeInvoker(orders []Order) {
	for i, v := range orders {
		if v.Priority.HigherThan(NewPriority("normal")) {
			// do some logic
			fmt.Println(i, v)
		}
	}
}
