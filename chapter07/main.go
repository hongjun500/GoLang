package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type StudentInterface interface {
	PrintAge()
	PrintName()
}

func main() {

	arr := make([]int, 5)
	arr = []int{43, 2435, 677, 43, 1}
	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(value []int) {
	// flag := true
	for i := 0; i < len(value)-1; i++ {
		for j := 0; j < len(value)-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
				// flag = false
			}
		}
		/*if flag == true {
			break
		}*/
	}
}

func (s Student) PrintAge() {
	fmt.Println(s.Age)
}
func (s Student) PrintName() {

}
func PrintAge(s Student) {

}
