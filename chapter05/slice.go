// @author hongjun500
// @date 2023/8/17 16:06
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 数组和切片

package main

import "fmt"

func arrFunc() {
	// 数组声明形式：var 数组名 [数组长度]数组类型，数组类型也可以是数组，因此可以构成数组类型的数组，即多维数组
	// 声明一个二维数组
	var arrArray [3][2]int
	fmt.Printf("arrArray 零值%v \n", arrArray[0][0])

	println("不定长数组部分------------------")

	// 使用三个...来定义数组长度时代表其长度根据数据初始化时的个数来决定
	var arrLen = [...]string{"hello"}
	fmt.Println("不定长的数组 arrLen 长度==", len(arrLen))

	var initArr = [...]int{1, 99}

	for i, v := range initArr {
		println("对不定长数组 initArr 进行遍历")
		fmt.Println("initArr 索引", i)
		fmt.Println("initArr 值", v)
	}

	var all [2]interface{}
	all[0] = "字符串string"
	all[1] = 88
	for i, v := range all {
		println("对不定长数组 all 进行遍历")
		fmt.Println("all 的i=", i)
		fmt.Println("all 的v=", v)
	}

	println("结构体数组部分------------------")

	users := [2]struct {
		name string
		age  int
	}{
		{"hongjun", 20},
	}

	for i, v := range users {
		fmt.Println("users 的i=", i)
		fmt.Println("users 的v=", v)
	}

	// 声明一个名为 array 类型为 int,长度为 10 的数组
	var array [10]int
	for i := 0; i < len(array); i++ {
		fmt.Printf("从%d开始遍历数组 array，并对其赋值\n", i)
		array[i] = i
	}
	fmt.Printf("数组 array 最后一个值 =%v \n", array[len(array)-1])
}

func spliceFunc() {
	// 由于数组是定长的(数组的长度是其类型的一部分，因此数组不能改变大小)
	// 因此在 Go 中更多使用的是切片，(切片是长度可以变化的数组)切片和数组又息息相关
	var arrSlice []string
	if arrSlice == nil {
		fmt.Println("切片 arrSlice 是空的")
		println("切片 arrSlice 的长度为", len(arrSlice))
	}
	// arrSlice[0] = "hello" // 这一行编译通过，但是运行报错，切片只是声明了但没有初始化
	// 本质上切片的底层实现是一个指向数组的指针，再没有存入一个具体的数组之前，就是nil
	// arrSlice = []string{""}  直接是以这种方式简单方式就可以使用
	// 或者通过内置函数 make 进行初始化 make([]T, len, cap)
	// 长度必须定义，容量可以省略并且容量必须大于等于长度

	println("对切片 arrSlice 进行初始化")
	arrSlice = make([]string, 12)
	arrSlice[0] = "hello"
	arrSlice[1] = "world"
	fmt.Printf("切片 arrSlice %v \n", arrSlice)
	fmt.Printf("切片 arrSlice 长度%v \n", len(arrSlice))
	fmt.Printf("切片 arrSlice 容量%v \n", cap(arrSlice))

	// 切片可以截取
	// 切片通过两个下标来截取，一个上界一个下界两者之间以冒号分割
	// 多种写法 splice[begin:end]，半开区间，包含 begin，排除 end
	// 多种写法 splice[:end]  从下标 0 开始到下标 end-1 的元素
	// 多种写法 splice[begin:]  从下标 begin 开始所有的元素
	// 多种写法 splice[:]  从下标 0 开始到下标最大值位置的元素(即所有元素)

	var srcarr = []string{"见到你", "很高兴"}
	fmt.Printf("最开始 srcarr 的值%v \n", srcarr)
	newarr := srcarr[:1]
	fmt.Printf("newarr 引用 srcarr 后的值%v \n", newarr)
	// 这里由于 newarr 是经过引用 srcarr 得到的，那么 newarr 改变了之后 srcarr 也改变了，其本质上切片就是由指针指向数组
	println("将 newarr 的第一个元素改为 空字符串")
	newarr[0] = ""
	fmt.Printf("切片 srcarr 被 newarr 引用后由于 newarr 改变，那么切片 srcarr 也改变为%v \n", srcarr)
	// 切片的复制 使用内置函数 copy
	strArr := make([]string, 2)
	strArr[0] = "go"
	fmt.Printf("拷贝之前strArr=%v \n", strArr)
	// 将 arrSlice 的元素拷贝给到 strArr(会覆盖 strArr 原有的元素),拷贝个数由两者长度最小值确定
	i := copy(strArr, arrSlice)
	fmt.Printf("总共拷贝了%d 个元素 \n", i)
	fmt.Printf("拷贝之后 strArr =%v \n", strArr)
	fmt.Printf("拷贝之后 arrSlice =%v \n", arrSlice)

	// 切片的追加使用内置函数 append
	var arr1 = []string{"hello", ",", "go"}
	arr2 := []string{"language"}
	// 将切片 arr2 的元素追加到 arr1 后面
	strings := append(arr1, arr2...)
	fmt.Printf("arr1 追加一个切片 arr2 之后的值%v \n", strings)
	// arr1 再追加元素
	newArr1 := append(strings, "is", "good")
	fmt.Printf("strings 追加单个字符串元素之后的值%v \n", newArr1)

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	fmt.Println("slice = ", slice)
}
