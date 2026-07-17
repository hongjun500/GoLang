package main

// SumInts 泛型出现之前：每种元素类型都得复制粘贴一个几乎一样的函数，
// 这正是泛型要解决的"重复代码 + 缺乏类型安全"的痛点。
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
// SumFloats32 注意 float32 精度有限，累加浮点数可能出现 15.540000000000001 之类的误差
func SumFloats32(m map[string]float32) float32 {
	var s float32
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
