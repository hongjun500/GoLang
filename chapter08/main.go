package main

import (
	"fmt"
)

func main() {

	ints := map[string]int64{
		"first":  6,
		"second": 8,
	}
	floats32 := map[string]float32{
		// "first":  6.66,
		// "second": 8.88,

		"first":  35.98,
		"second": 26.99,
	}
	floats64 := map[string]float64{
		// "first":  6.66,
		// "second": 8.88,
		// 这里计算的结果会是 15.540000000000001

		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("Non-generic sums: %v and floats32 = %v and floats64 = %v", SumInts(ints), SumFloats32(floats32), SumFloats64(floats64))
	fmt.Print("\n")
	fmt.Printf("generic sums: %v and %v", SumIntsOrFloats(ints), SumIntsOrFloats(floats64))

	fmt.Print("\n")

	myStrSplice := []string{"Java", "Go", "Js"}
	var myIntSplice []int
	myIntSplice = make([]int, 3)
	myIntSplice[0] = 6
	myIntSplice[1] = 8
	myIntSplice[2] = 7

	// 返回字符串切片和整数切片中指定数据的下标
	fmt.Printf("Index string result %v \nIndex int result %v", Index(myStrSplice, "Go"), Index(myIntSplice, 7))

}
