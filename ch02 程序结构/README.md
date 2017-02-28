# 第二章 程序结构

#### 2.3 变量
``` golang
var s string
fmt.Println(s)
// ""

var i, j, k int
var b, f, s = true, 2, 3, "four"

var f, err = os.Open(name)
// os.Open returns a file and an error
```

#### 2.3.1 简短变量声明
``` golang
// 变量类型根据表达式来自动推导
anim := gif.GIF{LoopCount: nframes}
freq := rand.Flout64() * 3.0
t := 0.0

i := 100
var boiling flout64 = 100
// 应该是字符串列表
var names []string
var err error
var p Point

// 和 var 声明语句一样，简短变量声明语句也可以用来声明和初始化一组变量；
i, j := 0, 1
// 加深记忆：:= 是变量声明语句，而 = 是变量赋值操作

i, j = j, i
// 交换 i 和 j 的值

// 简短变量声明语句也可以用函数的返回值来声明和初始化变量：
f, err := os.Open(name)
if err != nil {
    return err
}
// use file
f.close()
这里和 Python 一样
```

> 这里有一个比较微妙的地方：简短变量声明左边的变量可能并不完全都是刚刚声明的。如果有一些已经在相同的语法域声明过了，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。

``` golang
// 在下面代码中，第一个语句声明了 in 和 err 两个变量。
// 第二个语句只声明了 out 一个变量，然后对已经声明的 err 进行了赋值操作。
in, err := os.Open(infile)
out, err := os.Create(outfile)

// 简短变量声明语句中必须至少要声明一个新的变量
f, err := os.Open(infile)
f, err := os.Create(outfile)
// complie error: no new variables
// 解决方法是第二个简短变量语句改用普通的多重赋值

// 当然，这些是在同级词法域的基础上
```

#### 2.3.2. 指针
> 一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。通过指针，我们可以直接读或者更新对应变量的值，而不需要知道该变量的名字（如果变量有名字的话）。
```golang
var x int
// &x 将产生一个指向该整数变量的指针，指针对应的数据类型是 *int
// 被称为『指向 int 类型的指针』

// 如果指针名字为 p，那么可以说『 p 指针指向变量 x 』，或者
// 『p 指针保存了 x 变量的内存地址』

// *p 表达式对应 p 指针指向的变量的值

// 例子：
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

> 任何指针的零值都是 nil。如果 p 指向某个有效变量，那么 p != nil 测试为真。指针之间也可以进行测试的，只有当它们指向的同一个变量或全部是 nil 时才相等。

```golang
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```

> 在 go 语言中，返回函数中局部变量的地址也是安全的，调用 f 函数时创建局部变量 v，在局部变量地址被返回之后依然有效，因为指针 p 依然引用这个变量。
```golang
var p = f()

func f() *int {
    v := 1
    return &v
}
// 每次调用 f 函数都将返回不同的结果：
fmt.Println(f() == f()) // "false"
```

> 通过指针更新变量值（模拟 ++i )
```golang
func incr(p *int) int {
    *p++
    return *p
}

v := 1
incr(&v)              // side effect: v is now 2
fmt.Println(incr(&v)) // "3" (and v is 3)
```