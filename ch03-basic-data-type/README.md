# 第三章 基础数据类型


> Go 语言将数据类型分为四类:基础类型、复合类型、引用类型和接口类型。本章介绍基础类 型,包括:数字、字符串和布尔型。复合数据类型——数组(§4.1)和结构体(§4.2)——是 通过组合简单类型,来表达更加复杂的数据结构。引用类型包括指针(§2.3.2)、切片 (§4.2))字典(§4.3)、函数(§5)、通道(§8),虽然数据种类很多,但它们都是对程序 中一个变量或状态的间接引用。


### 3.1 整型

```golang
var u uint8 = 255 fmt.Println(u, u+1, u*u) // "255 0 1"

var i int8 = 127 fmt.Println(i, i+1, i*i) // "127 -128 1"
```

```golang
func compute() (value float64, ok bool) {
    // ...
    if failed {
        return 0, false
    }
    return result, true
}
```