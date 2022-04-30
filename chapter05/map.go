package main

import (
	"fmt"
	"sync"
)

func maps() {
	fmt.Println("map相关内容-----------------------------")
	var m map[string]string
	fmt.Printf("map没有使用make初始化时map==nil %v \n", m == nil)

	// 这种声明方式用的比较多
	m = make(map[string]string)
	fmt.Printf("map使用了make初始化时map==nil %v \n", m == nil)
	m["go"] = "lang"
	// 返回的elem是键对应的值，hasKey是一个布尔值，意为键是否存在，true存在，false不存在
	elem, hasKey := m["go"]
	if hasKey {
		fmt.Printf("映射m[\"go\"]的值=%v \n", elem)
	} else {
		fmt.Println("映射m[\"go\"]不存在")
	}
	// 这种声明则比较少
	maps := map[int]string{1: "壹", 2: "贰", 3: "叁"}

	for key, value := range maps {
		// fmt.Println(maps)
		// map是无序的，每次打印输出的结果可能会不一样
		fmt.Printf("key=%v,value=%v \n", key, value)
	}

	// map的元素删除：使用内置函数delete(m map[Type]Type1, key Type)
	// 将maps中键为1的元素删除
	delete(maps, 1)
	// delete(maps, 1)  即使这个键值对不存在时也不会存在
	fmt.Printf("maps删除键为1的元素后maps=%v", maps)

	fmt.Println("有锁的map部分:sync.Map")
	// sync.Map的底层并不是基于指针，而是结构体
	// 此种map的创建方式不需要使用make函数
	syncMap := sync.Map{}
	fmt.Printf("syncMap=%v \n", syncMap)
	// sync.Map的所有操作都是基于其自带方法的形式实现
	// 赋值
	syncMap.Store("apple", "苹果")
	// 这里使用了不同的键值对赋值是没有问题的，因为本身这种map结构底层就是结构体，结构是包含多种类型的集合
	syncMap.Store(2, "橘子")
	syncMap.Store(true, "是")
	syncMap.Store("origin", "橘子")
	// 取值
	value, hasKey := syncMap.Load(2)
	fmt.Printf("key=2, value=%v, 键是否存在：%v \n", value, hasKey)
	// 删除元素
	syncMap.Delete(3) // 键不存在时也不会出错排除极端情况

	// 遍历
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("syncMap自带的遍历方式获取key和value")

		fmt.Printf("key=%v, value=%v \n", key, value)
		return true
	})

	// 此处使用for循环100次对syncMap进行遍历，并且返回了一个false, 由于map的无序性，100次的打印结果不会一样，因为返回值是false,只会进行一次遍历取出某个key:value
	/*for i := 0; i < 100; i++ {
		// 遍历
		syncMap.Range(func(key, value interface{}) bool {
			fmt.Println("syncMap自带的遍历方式获取key和value")

			fmt.Printf("key=%v, value=%v \n", key, value)
			return false
		})
	}*/

	// 定义键和值都是任意类型的map
	interfaceMap := make(map[interface{}]interface{})
	interfaceMap[1] = 23
	interfaceMap["1"] = "字符串1"
	interfaceMap[true] = 1
	interfaceMap[false] = 0
	fmt.Println("---------interfaceMap---------", interfaceMap)

	map2 := map[interface{}]interface{}{}
	map2["map2"] = "interface"
	map2[1] = true
	map2[0] = false

	fmt.Println("---------interfaceMap【map2】---------", map2)
}
