package main

import (
	"fmt"
)

func main() {
	// 数组声明形式：var 数组名 [数组长度]数组类型，数组类型也可以是数组，因此可以构成数组类型的数组，即多维数组
	// 声明一个二维数组
	var arrArray [3][2]int
	fmt.Printf("arrArray零值%v \n", arrArray[0][0])

	// 使用三个...来定义数组长度时代表其长度根据数据初始化时的个数来决定
	var arrLen = [...]string{"hello"}
	fmt.Println("arrLen长度==", len(arrLen))

	var initArr = [...]int{1, 99}

	for i, i2 := range initArr {
		fmt.Println("initArr索引", i)
		fmt.Println("initArr值", i2)
	}

	var all [2]interface{}
	all[0] = "字符串string"
	all[1] = 88
	for i, v := range all {
		fmt.Println("all 的i=", i)
		fmt.Println("all 的v=", v)
	}
	type User struct {
		name string
		age  int
	}
	var users [2]User
	users[1] = User{
		age:  20,
		name: "hongjun",
	}
	for i, v := range users {
		fmt.Println("users 的i=", i)
		fmt.Println("users 的v=", v)
	}

	// 声明一个名为array类型为int,长度为10的数组
	var array [10]int
	for i := 0; i < len(array); i++ {
		// 从0开始给数组赋值
		array[i] = i
	}
	fmt.Printf("数组array%v \n", array)
	fmt.Printf("数组array最后一个值=%v \n", array[len(array)-1])

	// 由于数组是定长的(数组的长度是其类型的一部分，因此数组不能改变大小)
	// 因此在Go中更多使用的是切片，(切片是长度可以变化的数组)切片和数组又息息相关
	var arrSlice []string
	if arrSlice == nil {
		fmt.Println("切片arrSlice是空的")
	}
	// arrSlice[0] = "hello" // 这一行编译通过，但是运行报错，切片只是声明了但没有初始化
	// 本质上切片的底层实现是一个指向数组的指针，再没有存入一个具体的数组之前，就是nil
	// arrSlice = []string{""}  直接是以这种方式简单方式就可以使用
	// 或者通过内置函数make初始化 make([]T,len,cap)
	// 长度必须定义，容量可以省略并且容量必须大于等于长度
	arrSlice = make([]string, 12)
	arrSlice[0] = "hello"
	arrSlice[1] = "word"
	fmt.Printf("切片arrSlice%v \n", arrSlice)
	fmt.Printf("切片arrSlice长度%v \n", len(arrSlice))
	fmt.Printf("切片arrSlice容量%v \n", cap(arrSlice))

	// 切片可以截取
	// 切片通过两个下标来截取，一个上界一个下界两者之间以冒号分割
	// splice[begin:end],半开区间，包含begin,排除end
	// 有多种写法  splice[:high]  从下标0开始到下标end-1的元素
	// 有多种写法  splice[begin:]  从下标begin开始所有的元素
	// 有多种写法  splice[:]  从下标0开始到下标最大值位置的元素(即所有元素)
	var srcarr = []string{"见到你", "很高兴"}
	fmt.Printf("最开始srcarr的值%v \n", srcarr)
	newarr := srcarr[:1]
	// 这里由于newarr是经过引用srcarr得到的，那么newarr改变了之后srcarr也改变了，其本质上切片就是由指针指向数组
	newarr[0] = ""
	fmt.Printf("newarr的值%v \n", newarr)
	fmt.Printf("引用了切片srcarr后的值%v \n", srcarr)
	// 切片的复制 使用内置函数copy
	strArr := make([]string, 2)
	strArr[0] = "go"
	fmt.Printf("拷贝之前strArr=%v \n", strArr)
	// 将arrSlice的元素拷贝给到strArr(会覆盖strArr原有的元素),拷贝个数由两者长度最小值确定
	i := copy(strArr, arrSlice)
	fmt.Printf("总共拷贝了%d 个元素 \n", i)
	fmt.Printf("拷贝之后strArr=%v \n", strArr)
	fmt.Printf("拷贝之后arrSlice=%v \n", arrSlice)

	// 切片的追加 使用内置函数append
	var arr1 = []string{"hello", ",", "go"}
	arr2 := []string{"language"}
	// 将切片arr2的元素追加到arr1后面
	strings := append(arr1, arr2...)
	fmt.Printf("arr1追加一个切片arr2之后的值%v \n", strings)
	// arr1再追加元素
	newArr1 := append(strings, "is", "good")
	fmt.Printf("strings追加单个字符串元素之后的值%v \n", newArr1)
	learn1(12)

	fmt.Println("分割线map-------------------------------------------------")
	maps()
	mapSlice()

}

func learn1(lens int) {
	var s []int
	var factor int
	fmt.Printf("切片s原始长度= %v \n", len(s))
	fmt.Printf("factor= %v \n", factor)
	factor = lens
	s = make([]int, factor)
	fmt.Printf("切片s扩展后长度= %v \n", factor)

}
