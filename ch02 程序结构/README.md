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