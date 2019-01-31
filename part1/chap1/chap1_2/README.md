
## 1.2 下载安装Go 1.11

前面我们通过一段简单的实操，认识到Golang这门编程语言的优雅和便捷，接下来我们介绍一
下Go的安装。

本教程使用Go 1.11版本作为该语言的学习版本，因为1.11版本之后，Go语言引入了Module
概念，支持对依赖的管理。可以从[这里](https://studygolang.com/dl)下载您需要的版本。

关于IDE的选择，网络上有很多文章推荐，大家可以根据各自的使用习惯选用，在此仅推荐一款
非常好用的，Jetbrains的产品GoLand。大家可以从[这里](https://www.jetbrains.com/go/download/#section=mac)
了解并下载需要的版本。Goland是付费软件，金主或者情怀党可以自行付费。也可以去万能的
某宝上，买一个EDU邮箱(**很便宜**)，申请学生一年用的激活码，国外软件在这块做得都是比较好的。

GO安装完成后，我们可以通过以下命令查看我们的安装版本：

```go version ```

通过以下命令查看我们的安装情况：

```go env ```

正确安装会有类似如下显示：

```
GOARCH="amd64"
GOBIN=""
GOCACHE="~/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="~/wks"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="~/gobook/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="……"
```
通过```go help```了解go命令的使用方法，具体命令后续也会慢慢接触和介绍。

### go mod 介绍

go mod 命令是1.11版本之后，支持Module之后新增的特性之一。go mod 可以检查代码依赖，
自动生成项目的mod文件。任何一个源文件目录或者是空目录都可以作为一个module，只要包含
有mod文件。

#### 初始化Module

通过`go mod init [module name]`来初始化一个Module，完成后在该目录下生成一个
mod文件，里面只有一行`module [module name]`

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