package main

import (
	"flag"
	"fmt"
	"time"
)

//声明命令行参数绑定变量
var d time.Duration

func init() {
	//忽略error
	defalutD, _ := time.ParseDuration("1h20m30s")
	//定义命令行参数-d，将参数值绑定到变量d，默认值为“1h20m30s”
	flag.DurationVar(&d, "d", defalutD, "-d 1h20m30s")
	flag.Parse()
}

func main() {
	fmt.Println(d)

	//Input:
	//go run special_flag_3.go
	//Output:
	//1h20m30s

	//Input:
	//go run special_flag_3.go -d 3h30m30s
	//Output:
	//3h30m30s
}
