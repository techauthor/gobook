package main

import (
	"flag"
	"fmt"
)

var (
	b bool //声明一个参数b用来接收命令行参数
)

func init() {
	//声明一个默认值为false的命令行参数d，并将参数d的值绑定给变量b
	flag.BoolVar(&b, "d", false, "demo BoolVar function")
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
