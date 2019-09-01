package main

import (
	"flag"
	"fmt"
)

var (
	b *bool //声明一个参数b用来接收命令行参数
	s string
)

func init() {
	//声明一个默认值为false的命令行参数d，并将参数d的值赋给变量b
	b = flag.Bool("d", false, "demo Bool function")
	flag.StringVar(&s, "s", "hello_flag", "demo StringVar function")
	//解析命令行参数
	flag.Parse()
}

func main() {
	//通过访问变量b，使用参数d解析后的值
	fmt.Println(*b, s)

	// Input:
	// go run special_flag_1.go
	// Output：
	// false hello_flag

	// Input:
	// go run special_flag_1.go -d -s="thank flag package..."
	// Output：
	// true thank flag package...
}
