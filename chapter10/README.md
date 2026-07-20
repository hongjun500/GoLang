# 反射

> 反射是 在 **运行时** 读取并操作接口变量里那对`<动态类型,动态值>`的机制。

> 不仅能"看"（检查类型/取值），还能"改"（修改值）、"造"（创建新值）、"调"（动态调用方法）

> 它是 `encoding/json`、`gorm`、`validator`、RPC 框架、依赖注入容器的底层基石

## 一句话前提：反射操作的是"接口里的东西"

`reflect.TypeOf` / `reflect.ValueOf` 的形参都是 `any`（空接口）。传参时，值会先被
"装箱"进接口，反射再从接口里取出 `<类型, 值>` 这一对信息。**没有接口，就没有反射。**

## 反射三定律

### 第一定律：接口值 + 反射对象

```go
var x float64 = 3.4
t := reflect.TypeOf(x)  // reflect.Type，动态类型 float64
v := reflect.ValueOf(x) // reflect.Value，动态值 3.4
```

### 第二定律：反射对象 -> 接口值

```go
bak := v.Interface().(float64) // Interface() 装箱回 any,再类型断言取回
```

### 第三定律：想修改，`Value` 必须"可设置"（settable）

```go
elem := reflect.ValueOf(&x).Elem() // 传指针 + Elem() 才可寻址
elem.SetFloat(7.1)                 // OK；直接 ValueOf(x).SetFloat 会 panic
```

> 可设置 = 可寻址（addressable）+ 已导出（exported）。二者缺一不可

## `Type` vs `Kind`：最容易混的一对概念

| 概念   | 含义                     | 例子（`type MyInt int`） |
| ------ | ------------------------ | ------------------------ |
|        |                          |                          |
| `Type` | 声明的**具体类型**       | `main.MyInt`             |
| `Kind` | 编译器眼里的**底层类型** | `int`                    |

> 判断"是不是整数一族"看 `Kind`；判断"是不是某个具体类型"看 `Type`。 `switch v.Kind()` 是处理未知类型的标准姿势。

## 结构体字段与 `struct tag`

```go
t := reflect.TypeOf(u)
for i := range t.NumField() {
    f := t.Field(i)                 // StructField：字段元信息
    tag := f.Tag.Get("json")        // 读标签
    exported := f.IsExported()      // Go 1.17+，判断是否导出
}
```

- `t.Field(i)` 拿**元信息**（`StructField`:名字、类型、`tag`）
- `v.Field(i)` 拿**值**（`Value`:字段值）
- **未导出字段不能** `Interface()`，遍历时务必先 `IsExported()` 判断

## 动态调用方法

```go
m := reflect.ValueOf(obj).MethodByName("Add")
out := m.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)})
```

> 只有导出方法才在反射可见的方法集里；参数/返回值都以 `[]reflect.Value` 传递。

## 底层原理解析

### 接口的内存布局：反射的**数据源头**

`Go` 的接口在运行时是 **两个机器字**:

- 空接口 `any` -> `eface{*_type, data unsafe.Pointer}`(类型指针 + 数据指针)
- 非空接口 -> `iface{*itab, data unsafe.Pointer}`(`itab` 含类型与方法表)
  `reflect.TypeOf(x)` 就是把 `_type` 指针包装成 `reflect.Type`; `reflect.ValueOf(x)` 则把`<_type, data, flag>` 包装成 `reflect.Value`。

### `reflect.Value` 的`flag`：为什么 **能/不能** 改

`reflect.Value` 内部约等于:

```go
type Value struct {
    typ *rtype        // 类型信息
    ptr unsafe.Pointer // 指向数据
    flag uintptr       // 位标记：Kind、flagRO、flagAddr、flagIndir...
}
```

- `flagAddr`：是否可寻址。传指针 `Elem()` 后会被置上，才有资格 `Set`
- `flagRO(read-only)`：读未导出字段/方法得到的 `Value` 会被置上，`Set/Interface` 都会 `panic`

## 避坑指南

1. 性能：反射慢，且慢在**运行时 + 分配 + 无法内联**

- `Interface()` 会装箱、`Call()` 会分配 `[]reflect.Value`，值常常逃逸到堆上
- 反射代码编译器无法内联/优化，热路径上可能比直接调用慢一到两个数量级
- 最佳实践：
  - 把 `reflect.Type`、 `[]StructField`、`MethodByName` 的结果缓存（按类型做 `map` 缓存），别在循环里反复求
  - 类型有限且已知时，优先 `type switch` 而非反射
  - 极致性能场景用代码生成（`go generate` / `easyjson`）+ 或 `Go` 1.18+ 泛型替代反射

2. `panic` 雷区（反射最容易在运行时炸）

| 操作            | 触发条件                      | 现象    |
| --------------- | ----------------------------- | ------- |
|                 |                               |         |
| `v.Int()`       | `Value` 的 `Kind` 不是整型    | `panic` |
| `v.SetXxx()`    | 不可寻址 / 未导出             | `panic` |
| `v.Interface()` | 来自未导出字段                | `panic` |
| `v.Elem()`      | `v` 不是指针/接口             | `panic` |
| `v.Field(i)`    | `i` 越界 或 `v` 不是 `struct` | `panic` |

> 操作前先用 `Kind()` 判断种类、用 `CanSet()`/`CanAddr()` 判断可否设置

3. `reflect.ValueOf(nil` 返回零值 `Value`

- `IsValid()==false`, 对它做任何操作都会 `panic`，接口可能为 `nil`时先 `if v.IsValid()` 判断
- 判断内部指针/切片/`map` 是否为 `nil` 用 `v.IsNil()`，且要先确认 `v.Kind()`

4. 值 vs 指针：想改就必须传指针

```go
reflect.ValueOf(x)       // 副本，CanSet()==false
reflect.ValueOf(&x).Elem() // 可寻址，CanSet()==true
```

5. 未导出字段的三层结论:
   1. 能读（`Int`/`String()`）
   2. 不能 `Interface()`
   3. 不能 `Set() (flagRO)`

6. 比较用`reflect.DeepEqual`，别拿它当"万能等号"值
  - 适合比较含`slice`/`map`/指针的复合结构
  - 但它慢、且对“含函数字段”“NaN"”不同底层类型“有坑；能用 `==` 就别用 `DeepEqual`
