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
	fmt.Println("hello ch1_2...")
	router := gin.New()
	routerGroup := router.Group("/gobook")
	//routers.RegisterMiddleWare(g)

	routerGroup.Handle("GET", "/ch1_2/path", Name)
	router.Run(":8888")
}
