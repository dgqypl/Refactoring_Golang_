package refactoring

import "math"

// Circumference Simple Mechanics: 对于命名不合适的函数直接重新命名
func Circumference(radius float64) float64 {
	return 2 * math.Pi * radius
}
