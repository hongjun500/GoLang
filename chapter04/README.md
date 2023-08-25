# 类型和结构体

## 类型

### 自定义类型

> 使用关键字 `type` 可以自定义类型

```go
 type Name Type
```

```go
type MyInt int

var i int
var j MyInt
```
> `Go` 是静态类型，即只有一种类型在编译时已知并固定： int、float32、*MyType、[]byte等。
变量 `i` 和 `j` 具有不同的静态类型，尽管它们具有相同的基础类型，，但不能直接赋值，需要进行类型转换




- `Type` 代表任意类型，可以是内置类型，也可以是自定义类型
    - `int`, `float64`, `string`, `bool` 等
    - `func` 函数类型
    - `interface` 接口类型
    - `struct` 结构体类型
    - `map` 字典类型
    - `chan` 通道类型
    - `[]Type` 切片类型
    - `*Type` 指针类型
    - ...

### 接口类型的断言和转换

> **断言的类型必须是接口类型**

**断言的语法: `value, ok := x.(T)`,可以只返回一个值，但是如果断言失败，那么会抛出一个 panic 异常**

```go
var str any = "hello go"
s, ok := str.(string)
println(s, ok) // 会输出 hello go true
i := str.(int) // 会抛出一个 panic 异常
```

**接口类型还可以使用 `switch` 语法进行断言操作**

```go
var x any = "hello go" // 这里的 any 是一个空接口，代表任意类型
switch v := x.(type) {
case string:
fmt.Println("string", v)
case int:
fmt.Println("int", v)
case bool:
fmt.Println("bool", v)
default:
fmt.Println("unknown")
}
```

### 类型别名

> 使用关键字 `type` 可以定义类型别名

```go
type Name = Type
```

- 这里的 `Type` 只能是已存在的类型，包括自定义类型
- 需要注意*自定义类型*和*类型别名*的区别看似是有无 `=`，但其实自定义类型只会在存在代码中，编译完成时并不会存在该类型。
  而类型别名是语言本身就包含的类型，通过`fmt.Printf("%T",Name)`的方式打印，如果是自定义类型那么将会打印出`package.Name`
  ，而类型别名则是`package.Type`

## 结构体

> 关键字 `struct` 可以定义结构体

```go
package main

type Name struct {
	field1 string
	field2 bool
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}

var circle Circle
circle.X = 10
circle.Y = 10
circle.Radius = 5
fmt.Println(circle) // {Point:{X:10 Y:10} Radius:5}

```

**可以看成是一种自定义类型，但是它的成员是可以是任意类型的**

- 结构体字段书写格式：`字段名 字段类型`，并且注重大小写，如果字段名首字母大写，则该字段可以被外部包访问，否则只能在当前包内访问
- 结构体字段的访问：`结构体变量.字段名`
- 结构体的零值：结构体的每个字段都有一个零值，零值取决于字段类型，例如 `int` 的零值是 `0`，`string` 的零值是 `""`，`bool`
  的零值是 `false`，`struct` 的零值是每个字段的零值
- 结构体的初始化：`var 变量名 类型 = 结构体字面量`，结构体字面量的书写格式：`{字段1: 值1, 字段2: 值2, ...}`
  ，如果省略了字段名，则表示该字段使用零值
-
结构体的匿名字段：结构体的字段可以是任意类型，如果字段没有字段名，那么该字段就是一个匿名字段，匿名字段的类型必须是命名类型或者指向命名类型的指针，如果是指针类型，那么可以通过 `结构体变量.匿名字段.字段名`
的方式访问匿名字段的字段
-
结构体的嵌套：结构体的字段可以是结构体类型，这样的结构体称为嵌套结构体，嵌套结构体的字段可以通过 `结构体变量.字段名.字段名`
的方式访问
- 结构体的比较：结构体是值类型，如果结构体的每个字段都是可比较的，那么结构体也是可比较的，可以使用 `==` 或 `!=`
  进行比较，如果结构体的字段有不可比较的类型，那么结构体也是不可比较的

## 匿名结构体

> 匿名结构体没有类型名称，只有类型定义，可以直接使用字面量语法创建匿名结构体的值
> 匿名结构体一般用于临时场景

```go
package main

var user struct {
	name string
	age  int
}
person := struct {
name string
age  int
}{
name: "john",
age:  18,
}
user.name = "Tom"
user.age = 18
fmt.Println(user) // {Tom 18}
fmt.Println(person) // {Tom 18}
``` 


