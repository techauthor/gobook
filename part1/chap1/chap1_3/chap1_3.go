//chap1_3.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SayHello(cxt *gin.Context) {
	name := cxt.Query("name")
	cxt.JSON(200, "hello:"+name)
}

func main() {
	fmt.Println("hello chap1_3...")
	router := gin.New()
	routerGroup := router.Group("/gobook")
	routerGroup.Handle("GET", "/chap1_3/sayHello", SayHello)
	fmt.Println("open url `http://localhost:8888/gobook/chap1_3/sayHello?name=yourName` in your browser...")
	router.Run(":8888")
}
