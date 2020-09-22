
## 1.3 程序结构

### 名称

Go中的函数、变量、常量、类型、语句标签和包的名称遵循一个简单的规则：名称的开头是一个字母或下划线，后面可以跟任意数量的字符、数字、和下划线，并区分大小写。

Go有25个像if这样的关键字，只能用在语法允许的地方，它们不能作为名称使用：

break	default	func	interface	select

case	defer	go	map	struct

chan	else	goto	package	switch

const	fallthrough	if	range	type

continue	for	import	return	var

之所以刻意地将 Go 代码中的关键字保持的这么少，是为了简化在编译过程第一步中的代码解析。和其它语言一样，关键字不能够作标识符使用。

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符，其中包含了基本类型的名称和一些基本的内置函数，它们的作用都将在接下来的章节中进行进一步地讲解。

append	bool	byte	cap	close	complex	complex64	complex128	uint16

copy	false	float32	float64	imag	int	int8	int16	uint32

int32	int64	iota	len	make	new	nil	panic	uint64

print	println	real	recover	string	true	uint	uint8	uintptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。

Go语言中，第一个字母的大小写决定了其（变量、常量、类型、函数）可见性，如果第一个字母名称是大写开头，意味着他是可以导出的（export），即对于包外他是可见可访问的，可以被自己包以为的其他程序所使用。

名称本身没有长度限制，但Go的编程风格更倾向于短名称的使用，特别是作用域较小的局部变量，我们经常可以看到单字母如i这样的命名。通常，名称的作用域越大，就使用越长的名称才会更具意义。

Go程序员使用"驼峰式"式的命名风格。

### 声明

Go语言中主要有4个主要的声明：变量（var）、常量（const）、类型（type）、函数（func）。

Go程序存储在一个或多个以.go为后缀的文件里。每一个文件以package声明开头，表明文件属于哪个包。

package声明后面是import声明，然后是包级别的类型、变量、常量、函数声明，不区分顺序。

例如：

```go
//包声明
package main

//import声明
import "fmt"

//以下声明不区分顺序

//类型声明
type A string

//变量声明
var a A = "hello"

//常量声明
const b = 100

func main(){
    fmt.Println(a)
    fmt.Println(b)

    //output:
    //"hello"
    //100
}
```

### 变量

var声明创建一个具体类型的变量，请记住变量声明的通用形式：

`var name type = expression`

**类型和表达式部分可以胜率一个，但是不能都省略。如果类型省略，他的类型将有初始化表达式决定；如果表达式省略，其初始值对应于该类型的零值。**

对于数字来说，零值是0；布尔值的零值是false;字符串是""；接口和引用类型（slice、指针、map、通道、函数）是nil。

对于一个像数组或者结构体这样的复合类型，零值是其所有元素或者成员的零值。

**零值机制保障所有的变量是*定义良好的*，Go里面不存在未初始化变量。**

可以声明一个变量列表，并选择使用对应的表达式列表对其进行初始化。忽略类型允许声明多个不同类型的变量。

```go

var i,j,k int                   //int,int,int
var a,b,c = true,2.3,"four"     //bool,float64,string

```

变量也可以通过调用返回多个值的函数进行初始化：

`var f,e = os.Open(name)        //os.Open 返回一个文件和一个错误`

需要注意的是左边被赋值的变量数量，要和函数的返回值数量一样多。

#### 短变量声明

在函数中，一种称作短变量声明的可选形式可以用来声明和初始化局部变量。格式为：

`name := expression`

短变量声明的*name*类型由*expression*的类型决定。

```go
func main() {
    name := "go"            //name为string类型
    i,j := 0,1              //多个变量也可以以短变量的方式声明和初始化
    f,e := os.Open(name)    //短变量的声明也可以以调用函数返回值的方式声明
}
```

注意，:= 表示声明，= 表示赋值。

### 赋值

赋值语句用来更新变量所指的值，他最简单的形式由赋值符=，以及符号左边的变量和右边的表达式组成。

`name = expression`

另外，每一个算术和二进制操作符有一个对应的赋值操作符，他避免了在表达式中重复变量本身。

```go
func main() {
    m := 1
    v := 2
    m += v      //等同于m = m + v
    v++         //等同于v = v + 1
    v--         //等同于v = v - 1
    
}
```

其他算术和二进制操作符也有对应的赋值操作符，如：m *= v

注意：Go里面没有++v和--v注意的操作。

可以把不需要的值赋值给空标识符_，Go语言不允许出现无用的临时变量。

```go
f,_ := os.Open(name)    //丢弃函数的第二个返回值
```






