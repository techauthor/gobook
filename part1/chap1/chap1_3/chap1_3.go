//chap1_3.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Name(c *gin.Context) {
	path := c.Request.URL.Path
	c.JSON(200, path)
}

func main() {
	fmt.Println("hello chap1_3...")
	router := gin.New()
	routerGroup := router.Group("/gobook")
	//routers.RegisterMiddleWare(g)

	routerGroup.Handle("GET", "/chap1_3/path", Name)
	router.Run(":8888")
}
