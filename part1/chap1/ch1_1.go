//ch1_1.go
package main

import (
	"fmt"
	"net/http"
)

func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Golang")
}

func main() {
	http.HandleFunc("/hello", HelloGo)
	http.ListenAndServe(":8888", nil)
	fmt.Println("hello ch2...")
}
