package refactoring

import "math"

// deprecated
func Circum(radius float64) float64 {
	return Circumference(radius)
}

// Circumference Migration Mechanics: 适用于函数做为对外发布的API这种场景，即已经有客户端在使用。
// 将旧函数的函数体包装成一个合适名字的新函数，旧函数改为调用新函数，同时标记旧函数为deprecated。以后如果有用到该函数的地方，直接调用新函数。
// 同时通知使用旧函数的客户端，将调用旧函数的地方改为调用新函数，最后我们再删除旧函数。（当然，也有可能客户端一直使用旧函数）
func Circumference(radius float64) float64 {
	return 2 * math.Pi * radius
}

func OneOfExistedInvoker() {
	Circum(1.0)
}

func OneOfNewInvoker() {
	Circumference(1.0)
}
