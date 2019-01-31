
## 1.1 一个简单的HttpServer


首先我们通过一段简单的实操，来建立对Golang这门编程语言的一个感性上的认识：


```Golang
//ch1_1.go
package main

import (
	"fmt"
	"net/http"
)

func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Golang")
}

func main() {
	http.HandleFunc("/hello", HelloGo)
	http.ListenAndServe(":8888", nil)
	fmt.Println("hello ch2...")
}
```

**程序说明：**

- 当我们通过/hello url访问时，由HelloGo函数来负责执行，并响应给客户端；
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