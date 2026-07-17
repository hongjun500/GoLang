package main

// ============ 1. 泛型函数：类型参数写在函数名后的方括号里 ============

// SumIntsOrFloats 使用"内联联合约束" int64 | float64 直接约束 V。
//   - K comparable：map 的 key 必须可比较（这是 map key 的固有要求）。
//   - V int64 | float64：值只能是这两种类型之一。
//
// 类型实参通常无需手写，由编译器从实参自动推断
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V // 泛型里用 var 取零值；不能写 sum := 0，因为 0 不一定是 V 的合法字面量
	for _, v := range m {
		sum += v
	}
	return sum
}

// Number 接口约束,也称作声明类型约束,它的类型集合是 {int64, float64}
// 把约束抽出来复用：约束本质是一个"类型集合"接口
// ~ 表示"底层类型"：加上 ~ 后，type MyInt64 int64 这类命名类型也满足约束
type Number interface {
	~int64 | ~float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// Index 返回指定数据 x 在 s 中的下标，如果不存在返回 -1
// 只要求 T 可比较（== 需要 comparable）；不必、也不应把返回值做成泛型
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// Map 把 []T 通过函数 f 转换为 []U，T、U 是两个互相独立的类型参数。
// 这是生产中泛型最高频的落地形态（配合 Filter/Reduce）。
func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Stack 是类型安全的泛型栈，元素类型为 T。
type Stack[T any] struct {
	items []T
}

// Push 注意：方法【不能】再声明自己的类型参数，只能使用接收者上的 T。
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop 返回栈顶元素与是否成功；空栈时返回 T 的零值 + false。
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}
