package main

import (
	"fmt"
	"sync"
)

func maps() {
	fmt.Println("map相关内容-----------------------------")
	var m map[string]string
	fmt.Printf("map 没有使用内置函数 make 初始化时 map == nil %v \n", m == nil)

	// 对于 map 类型的变量，必须使用 make 函数进行初始化
	m = make(map[string]string)
	fmt.Printf("map 使用了 make 初始化后 map==nil %v \n", m == nil)
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
		// map 是无序的，每次打印输出的结果可能会不一样
		fmt.Printf("key=%v,value=%v \n", key, value)
	}

	// 定义一个 map userMap
	var userMap map[any]any

	// 定义一个切片 users
	var users []map[any]any

	// 初始化 userMap 并赋值
	userMap = make(map[any]any)
	// 一般会直接使用 userMap := make(map[any]any) 这种方式，更推荐
	userMap["name"] = "hongjun500"
	userMap["age"] = 20

	// 初始化 users 并赋值
	users = make([]map[any]any, 2)
	users[0] = userMap
	// 使用 for 循环遍历 users
	for index, user := range users {
		fmt.Println("开始遍历 users")
		fmt.Println("users->index:", index)
		fmt.Println("users->user:", user)
	}

	// map 的元素删除：使用内置函数 delete(m map[Type]Type1, key Type)
	// 将 maps 中键为1的元素删除
	delete(maps, 1)
	// delete(maps, 1)  即使这个键值对不存在时也不会报错
	fmt.Printf("maps 删除键为 1 的元素后 maps =%v", maps)

	// 此处使用for循环100次对syncMap进行遍历，并且返回了一个false, 由于map的无序性，100次的打印结果不会一样，因为返回值是false,只会进行一次遍历取出某个key:value
	/*for i := 0; i < 100; i++ {
		// 遍历
		syncMap.Range(func(key, value any) bool {
			fmt.Println("syncMap自带的遍历方式获取key和value")

			fmt.Printf("key=%v, value=%v \n", key, value)
			return false
		})
	}*/

	// 定义键和值都是任意类型的map
	interfaceMap := make(map[any]any)
	interfaceMap[1] = 23
	interfaceMap["1"] = "字符串1"
	interfaceMap[true] = 1
	interfaceMap[false] = 0
	fmt.Println("---------interfaceMap---------", interfaceMap)

	map2 := map[any]any{}
	map2["map2"] = "interface"
	map2[1] = true
	map2[0] = false

	fmt.Println("---------interfaceMap【map2】---------", map2)

	for k, v := range interfaceMap {
		fmt.Println(k, v)
	}

}

// syncMapFunc 有锁的 map
func syncMapFunc() {
	fmt.Println("有锁的 sync.Map 部分-----------------------------")
	// sync.Map的底层并不是基于指针，而是结构体
	// 此种map的创建方式不需要使用make函数
	syncMap := sync.Map{}
	fmt.Printf("syncMap=%v \n", syncMap)
	// sync.Map的所有操作都是基于其自带方法的形式实现
	// 赋值
	syncMap.Store("apple", "苹果")
	// 这里使用了不同的键值对赋值是没有问题的，因为本身这种 map 结构底层就是结构体，结构是包含多种类型的集合
	syncMap.Store(2, "橘子")
	syncMap.Store(true, "是")
	syncMap.Store("origin", "橘子")
	// 取值
	value, hasKey := syncMap.Load(2)
	fmt.Printf("key=2, value=%v, 键是否存在：%v \n", value, hasKey)
	// 删除元素
	syncMap.Delete(3) // 键不存在时也不会出错排除极端情况

	// 遍历
	syncMap.Range(func(key, value any) bool {
		fmt.Println("syncMap 自带的遍历方式获取 key 和 value ")
		fmt.Printf("key=%v, value=%v \n", key, value)
		return true
	})
}

func mapSliceFunc() {
	// 元素类型为 map 的切片(长度为2)
	// make(t Type, size ...IntegerType) Type
	// 创建了一个值为切片，并且切片的类型是一个 map,而 map 的键值对都是 string
	var mapSlice = make([]map[string]string, 2)
	// fmt.Println("初始化创建的mapSlice", mapSlice)
	for key, value := range mapSlice {
		fmt.Printf("切片 mapSlice 未赋值时的 key=%v   ,value = %v \n", key, value)
	}

	// 对切片 mapSlice 进行初始化赋值
	mapSlice[0] = make(map[string]string)
	mapSlice[0]["name"] = "hongjun500"
	mapSlice[0]["age"] = "23"
	// mapSlice[1]["gener"] = "2"  这一行编译通过，但是运行报错因为这里代表取切片下标为 1 的元素 map,且 map 的键为 gender 的值
	for key, value := range mapSlice {
		fmt.Printf("切片 mapSlice 通过赋值后的 key =%v   ,value = %v \n", key, value)

		// 如果此轮循环体中切片的元素不为空则对元素即map再次进行一个遍历
		if value != nil {

			fmt.Printf("切片位置 mapSlice[%v],此轮循环体中切片不为空 \n", key)
			/*for key, value := range value {
				fmt.Printf("切片mapSlice[%v] = %v   , \n", key, value)
			}*/
		} else {
			fmt.Printf("切片位置 mapSlice[%v],此轮循环体中切片是空的 \n", key)
		}

		if len(value) != 0 {

			for key, value := range value {
				fmt.Printf("切片 mapSlice[%v] = %v  \n", key, value)
			}
		}
	}

	// 值为切片的map示例
	// make(t Type, size ...IntegerType) Type 定义
	// 这一行的 code 示例代表创建了一个 map,并且 map 的键是 int 类型吗，map 的值为 string 类型的切片
	var sliceMap = make(map[int][]string)

	value1, key1 := sliceMap[1]
	value2, key2 := sliceMap[2]
	if key1 && key2 {
		fmt.Printf("sliceMap[\"%v\"]=%v, sliceMap[\"%v\"]=%v \n", key1, value1, key2, value2)
	} else {
		value1 = []string{"见到你"}
		value2 = []string{"很高兴"}
		msg := append(value1, value2...)
		sliceMap[0] = msg
	}
	fmt.Println(sliceMap[0])
}
