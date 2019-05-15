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
	*v = ValueDemo(strings.Split("a,b,c", ","))
	return "example `-v a,b,c`"
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
