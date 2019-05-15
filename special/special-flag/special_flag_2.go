package main

import (
	"flag"
	"fmt"
	"strings"
)

type ValueDemo []string

/*
flag.Value 接口声明
type Value interface {
	String() string
	Set(string) error
}
*/

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
	return "example `a,b,c`"
}

var v ValueDemo

func init() {
	flag.Var(&v, "v", "demo v")
	flag.Parse()
}

func main() {
	fmt.Println(v)
}
