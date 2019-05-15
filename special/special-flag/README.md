
## Special-Flag


Golang对于命令行参数有很好的支持，通过flag包，可以快速简易地完成命令行参数的定义和解析。

Golang命令行参数的支持分为两步：定义与解析。在开始这两部分的学习之前，我们需要了解Golang支持怎样的参数格式。

Golang支持以下三种命令行参数格式：

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