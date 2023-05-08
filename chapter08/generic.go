package main

// SumIntsOrFloats 累加和，入参数据为 int64/float64 类型 -> 泛型
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Number 接口约束,也称作声明类型约束
type Number interface {
	int
}

// Index 返回指定数据的索引，如果不存在返回 -1
// comparable 用作参数类型的约束
func Index[T comparable, V Number](s []T, x T) V {
	// func Index[T comparable](s []T, x T) any {
	for i, v := range s {
		if v == x {
			// return i
			return V(i)
		}
	}
	return -1
}
