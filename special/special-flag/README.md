
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

第一种定义方法：
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



```Golang
//chap1_1.go
package main

import (
	"fmt"
	"net/http"
)

func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Golang")
}

func main() {
	fmt.Println("hello chap1_1...")
	http.HandleFunc("/gobook/chap1_1", HelloGo)
	http.ListenAndServe(":8888", nil)
}
```

**程序说明：**

- 当我们通过/gobook/chap1_1 访问时，由HelloGo函数来负责执行，并响应给客户端；
- HelloGo函数只执行一个非常简单的逻辑，输出字符串“Golang”；
- ListenAndServe函数，监听8888端口，并启动http server服务；

注意：8888前面的`:`不能少

启动命令:

```
$ cd ~/gobook/part1/chap1/
$ go run ch1_1.go
```

浏览器输入[http://localhost:8888/hello](http://localhost:8888/hello)

如果还不尽兴，可将代码编译为可执行文件:

```
$ cd ~/gobook/part1/chap1/
$ go build
```

当前目录会多出一个chap1文件，该文件是一个可执行文件，直接运行该文件:

```
$ cd ~/gobook/part1/chap1/
$ ./chap1
```

浏览器输入[http://localhost:8888/hello](http://localhost:8888/hello)

Bingo~!

---

### 知识点

#### 1、Go文件结构

package *pkgname*

import （

  *包名……*

）

例如：

```
import (
	"fmt"
	"net/http"
)
```

包名声明

#### 2、main函数

和大多数编程语言一样，main函数是go程序的入口函数,且不能有参数和返回值。