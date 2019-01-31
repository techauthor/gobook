
## 1.2 下载安装Go 1.11

前面我们通过一段简单的实操，认识到Golang这门编程语言的优雅和便捷，接下来我们介绍一
下Go的安装。

本教程使用Go 1.11版本作为该语言的学习版本，因为1.11版本之后，Go语言引入了Module
概念，支持对依赖的管理。可以从[这里](https://studygolang.com/dl)下载您需要的版本。

关于IDE的选择，网络上有很多文章推荐，大家可以根据各自的使用习惯选用，在此仅推荐一款
非常好用的，Jetbrains的产品GoLand。大家可以从[这里](https://www.jetbrains.com/go/download/#section=mac)
了解并下载需要的版本。Goland是付费软件，金主或者情怀党可以自行付费。也可以去万能的
某宝上，买一个EDU邮箱(**很便宜**)，申请学生一年用的激活码，国外软件在这块做得都是比较好的。

Go安装完成后，我们可以通过以下命令查看我们的安装版本：

```go version ```

通过以下命令查看我们的安装情况：

```go env ```

正确安装会有类似以下显示：

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
通过`go help`可以了解go命令的使用方法，具体命令后续也会慢慢接触和介绍。

### go mod 介绍

go mod 命令是1.11版本之后，支持Module之后新增的特性之一。Go Module可以检查代码依赖，
自动更新项目的mod文件。任何一个源文件目录或者是空目录都可以作为一个module，只要包含
有mod文件。

#### 1、初始化Module

通过

`go mod init [module name]`

先来初始化一个Module，完成后在该目录下生成一个
mod文件，里面只有一行

`module [module name]`

本项目则初始化为：

`go mod init gobook`

接下来，我们通过以下命令，手动触发依赖检查，更新mod文件。(后续随时都可以使用该命令，
新增缺失的或删除未用的依赖，更新到mod文件中)

```
go mod tidy
```

后续当我们使用go build，go test以及go list命令时，go也会自动检查项目依赖，并更新
go.mod文件。

更新完成后，可以通过命令：

```
go list -m all
```
查看项目的依赖信息，或者用`go mod graph`也能达到同样的效果。大家也可以直接打开
gobook项目中的go.mod文件查看。

更多go mod使用信息，请参考`go mod help`

#### 2、管理依赖版本

通过上述操作，默认情况下，Module会依赖最新的版本，包括最新的tag版本。但通常最新的
版本不一定是稳定版或者Release版本，这时候我们需要指定版本该如何操作呢？
在此之前，我们需要先了解版本号的构成。

以下是项目中对`github.com/gin-contrib/sse`
包的依赖示例，之后接**空格**，**空格**之后是一串由字母和数字组成的字符，通过短线分隔开。

```
github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
```

短线分隔开的三段组成分别表示：

**版本号-时间戳-Hash值**

通常多数情况下，在我们非常明确依赖的版本时，可以在空格后直接写上依赖的版本号即可，如下：

```
github.com/gin-gonic/gin v1.3.0
```

如果是一些非Release版本或者还没有正式版本时，我们如何操作呢？

通常我们使用的版本管理工具，都会给每一次提交，赋予一个唯一的标识符。以Github为例，
对应gin v1.3.0 这一次提交，可以直接在Github上找到如下标识信息：

```
commit b869fe1415e4b9eb52f247441830d502aece2d4d
```

我们可以直接把`b869fe1415e4b9eb52f247441830d502aece2d4d`复制到版本号的位置，如下：

```
github.com/gin-gonic/gin b869fe1415e4b9eb52f247441830d502aece2d4d
```

保存后使用以下命令

```
go list -m -json all
```

观察下面变化

```
{
	"Path": "github.com/gin-gonic/gin",
	"Version": "v1.3.0", //注意这里
	"Time": "2018-08-14T08:58:52Z",
	"Dir": "~/wks/pkg/mod/github.com/gin-gonic/gin@v1.3.0",
	"GoMod": "~/wks/pkg/mod/cache/download/github.com/gin-gonic/gin/@v/v1.3.0.mod"
}
```
再次打开mod文件，已经非常直观的显示版本了。如果使用GoLand IDE，IDE会自动帮你完成go list 动作。

当然，我们也可以通过`go mod edit`命令完成上述动作：

```
go mod edit -require="github.com/gin-gonic/gin@b869fe1415e4b9eb52f247441830d502aece2d4d"
```

---

### 小历史

在go module之前，go项目是需要依赖GOPATH这个环境变量的。一个标准的go项目，需要
有bin、src，pkg等目录构成。其中：

- src 用来存放源码文件
- pkg 用来存放编译时生成的中间文件
- bin 则用来存放编译后生成的可执行文件