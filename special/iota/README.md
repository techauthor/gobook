
## Special-Iota


iota 是Golang中的常量生成器，在常量声明中，iota的取值从0开始，逐项加1。

```
-flag
-flag=x
-flag x
```

假设我们有一个参数d，对应的三种格式如下（其中命令名称假设为cmd）：
```
cmd -d
cmd -d="abc"
cmd -d "abc"
```

三种命令行参数的格式，对于参数值的数据类型也有相应的要求。

* `-flag`格式，参数flag的值，只支持bool类型，也就是true和false；
* `-flag=x`格式，参数flag的值，可以是Golang的基本数据类型；
* `-flag x`格式，参数flag的值，只能是非bool类型。

了解了命令行格式和使用方法，接下来就是如何定义我们要使用的参数。

### 命令行参数的定义

命令参数的定义也有三种方式，我们需要借助flag包来完成命令行参数的定义，对应了三类函数。

```
1. 命名如TypeVal类的函数，如flag.StringVar(...);flag.IntVar(...);flag.BoolVar(...)等等；
2. 命名如Type类的函数，如flag.String(...);flag.Int(...);flag.Bool(...)等等；
3. flag.Var(...)函数
```

其中第三种方式用来自定义复杂的参数类型，会在下一节进行说明，下面我们先来看前面两种定义方式。

我们以Bool类型为例（其他类型使用上相同）。

#### 第一种定义方法：
```
// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.Var(newBoolValue(value, p), name, usage)
}
```

上面是flag包中对于BoolVar的函数签名。该方法通过绑定一个bool类型的变量p，用来接收命令行参数的值。

```
- p *bool 绑定的bool类型变量，用来接收命令行参数的值；
- name string 定义的命令行参数名，如`cmd -d`，我们可以通过赋值`name="d"`，来定义一个名为d的命令行参数；
- usage string 参数的使用描述说明。当传入的参数值有误，或者漏传时，或者通过`-h`获取帮助信息时，会显示给客户端。
```

示例代码：

```
package main

import (
	"flag"
	"fmt"
)

var(
	b bool	//声明一个参数b用来接收命令行参数
)

func init() {
	//声明一个默认值为false的命令行参数d，并将参数d的值绑定给变量b
	flag.BoolVar(&b,"d",false,"demo BoolVar function")
	//解析命令行参数
	flag.Parse()
}

func main() {
	//通过访问绑定变量b，使用参数d解析后的值
	fmt.Println(b)

	// Input:
	// go run special_flag_1.go
	// Output：
	// false

	// Input:
	// go run special_flag_1.go -d
	// Output：
	// true
}

```

#### 第二种定义方法：
```
// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *bool {
	return CommandLine.Bool(name, value, usage)
}
```
该函数参数和第一种方式完全相同，不同的是此时的参数值是通过返回值的方式返回。

示例代码如下：
```
package main

import (
	"flag"
	"fmt"
)

var (
	b *bool //声明一个参数b用来接收命令行参数
)

func init() {
	//声明一个默认值为false的命令行参数d，并将参数d的值赋给变量b
	b = flag.Bool("d", false, "demo Bool function")
	//解析命令行参数
	flag.Parse()
}

func main() {
	//通过访问变量b，使用参数d解析后的值
	fmt.Println(*b)

	// Input:
	// go run special_flag_1.go
	// Output：
	// false

	// Input:
	// go run special_flag_1.go -d
	// Output：
	// true
}

```

special_flag_1.go 中还提供了StringVar函数的示例，供参考。

#### 第三种定义方法：
```
// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func Var(value Value, name string, usage string) {
	CommandLine.Var(value, name, usage)
}
```

使用flag.Var函数定义命令行参数，我们可以使用自定义类型并实现flag.Value接口来接收命令行参数。以下是flag.Value接口的声明:

```
// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
// If a Value has an IsBoolFlag() bool method returning true,
// the command-line parser makes -name equivalent to -name=true
// rather than using the next command-line argument.
//
// Set is called once, in command line order, for each flag present.
// The flag package may call the String method with a zero-valued receiver,
// such as a nil pointer.
type Value interface {
	String() string
	Set(string) error
}
```

其中我们需要在Set方法中完成参数值的解析，在String方法中设置参数的默认值。

假设我们有一个参数v被调用
```
cmd -v a,b,c
```

下面我们使用自定义的string slice来接收"a,b,c"的值，首先定义类型:
```
//自定义一个string slice来接收命令行参数
type ValueDemo []string
```

实现flag.Value接口：
```
//实现flag.Vaule接口Set方法
func (v *ValueDemo) Set(s string) (e error) {
	//解析命令行参数值
	*v = ValueDemo(strings.Split(s, ","))
	return
}

//实现flag.Vaule接口String方法
func (v *ValueDemo) String() string {
	//初始化默认值
	*v = []string{"a","b","c"}
    	return fmt.Sprint(*v)
}
```

上面代码比较简单，在Set函数中，切分","分隔符并赋值给ValueDemo实例，String函数中设置默认值为“a,b,c”。接下来调用flag.Var函数：

```
//声明类型为ValueDemo的变量v，用来接收参数值
var v ValueDemo

func init() {
	//定义名为“v”的命令行参数，并使用自定义类型为ValueDemo的变量v绑定参数值
	flag.Var(&v, "v", "demo flag.Var function")
	flag.Parse()
}
```

完整代码如下：
```
package main

import (
	"flag"
	"fmt"
	"strings"
)

//自定义一个string slice来接收命令行参数
type ValueDemo []string

//实现flag.Vaule接口Set方法
func (v *ValueDemo) Set(s string) (e error) {
	//解析命令行参数值
	*v = ValueDemo(strings.Split(s, ","))
	return
}

//实现flag.Vaule接口String方法
func (v *ValueDemo) String() string {
	//初始化默认值
	*v = []string{"a","b","c"}
    	return fmt.Sprint(*v)
}

//声明类型为ValueDemo的变量v，用来接收参数值
var v ValueDemo

func init() {
	//定义名为“v”的命令行参数，并使用自定义类型为ValueDemo的变量v绑定参数值
	flag.Var(&v, "v", "demo flag.Var function")
	flag.Parse()
}

func main() {
	fmt.Println(v)

	// Input:
	// go run special_flag_2.go
	// Output：
	// false [a b c]

	// Input:
	// go run special_flag_2.go -v java,go,c++
	// Output：
	// [java go c++]
}

```

### 命令行参数的解析

通过调用flag.Parse()方法完成命令行参数的解析，上述实例的Init函数中已经使用到。需要注意的是，参数的解析与支持的命令行格式有关，无效的命令行格式无法被正确解析。

### 一些例子
下面介绍一些复杂的例子

#### flag内置Duration类型参数使用

除了基本的数据类型，flag包也提供了Duration类型的参数定义函数:flag.Duration(...)和flag.DurationVar(...)，下面就以DurationVar函数为例，介绍Duration类型参数的使用方法。

参考代码：

```
package main

import (
	"flag"
	"fmt"
	"time"
)
//声明命令行参数绑定变量
var d time.Duration

func init() {
	//忽略error
	defalutD, _ := time.ParseDuration("1h20m30s")
	//定义命令行参数-d，将参数值绑定到变量d，默认值为“1h20m30s”
	flag.DurationVar(&d, "d", defalutD, "-d 1h20m30s")
	flag.Parse()
}

func main() {
	fmt.Println(d)

	//Input:
	//go run special_flag_3.go
	//Output:
	//1h20m30s

	//Input:
	//go run special_flag_3.go -d 3h30m30s
	//Output:
	//3h30m30s
}

```

### 重点强调

上述介绍了三种方法来定义命令行参数，不论是哪一种方法，归根结底最终都是调用Var(...)函数完成，包括基本的数据类型。

flag包中定义了所有基本数据类型对应于flag.Value接口的实现，如：boolValue、intValue、stringValue等等，他们都实现了flag.Value接口。

每一个flag都对应于一个Flag类型的实例。下面是Flag类型的定义：
```
// A Flag represents the state of a flag.
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```
针对flag.Value接口的实例化

`flag := &Flag{name, usage, value, value.String()}`

### 有用的函数

#### 1. flag.PrintDefaults()、flag.Usage()

输出所有flag的默认值和使用说明

#### 2. flag.Lookup(name string)

根据所定义的命令行参数命，返回对应的Flag类型实例的指针。

[返回目录](https://github.com/techauthor/gobook/blob/master/README.md)


