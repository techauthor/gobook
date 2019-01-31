
## 1.2 下载安装Go 1.11



前面我们通过一段简单的实操，认识到Golang这门编程语言的优雅和便捷，接下来我们介绍一下Go的安装。

本教程使用Go 1.11版本作为该语言的学习版本，因为1.11版本之后，Go语言引入了Module概念，支持对依赖的管理。
可以从[这里](https://studygolang.com/dl)下载您需要的版本。

关于IDE的选择，网络上有很多文章推荐，大家可以根据各自的使用习惯选用，在此仅推荐一款非常好用的，Jetbrains的产品GoLand。
大家可以从[这里](https://www.jetbrains.com/go/download/#section=mac)了解并下载需要的版本。Goland是付费软件，金主或者情怀党可以自行付费。
也可以去万能的某宝上，买一个EDU邮箱，申请学生一年用的激活码，国外软件在这块做得都是比较好的。

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