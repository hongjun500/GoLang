package main

import "fmt"

func main() {
	ints := map[string]int64{"first": 6, "second": 8}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	// 泛型出现前：每种类型一个函数
	fmt.Printf("非泛型: SumInts=%v, SumFloats64=%v\n", SumInts(ints), SumFloats64(floats))

	// 泛型：一个函数搞定，类型实参自动推断
	fmt.Printf("泛型(内联约束): %v, %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("泛型(复用约束): %v, %v\n", SumNumbers(ints), SumNumbers(floats))

	// 也可以显式指定类型实参（一般不需要，仅在无法推断时使用）
	fmt.Printf("显式类型实参: %v\n", SumNumbers[string, int64](ints))

	// Index：返回值天然是 int
	strs := []string{"Java", "Go", "JS"}
	nums := []int{6, 8, 7}
	fmt.Printf("Index: %v, %v\n", Index(strs, "Go"), Index(nums, 7))

	// Map：切片元素类型转换 []string -> []int
	lengths := Map(strs, func(s string) int { return len(s) })
	fmt.Printf("Map(取长度): %v\n", lengths)

	// 泛型类型 Stack[int]：类型安全，取出来无需断言
	var st Stack[int]
	st.Push(1)
	st.Push(2)
	if top, ok := st.Pop(); ok {
		fmt.Printf("Stack.Pop: %v\n", top)
	}
}
