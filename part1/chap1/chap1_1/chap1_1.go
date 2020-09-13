//chap1_1.go
package main

import (
	"fmt"
	"net/http"
)

func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Golang")
}

func main() {
	fmt.Println("hello chap1_1...")
	http.HandleFunc("/gobook/chap1_1", HelloGo)
	http.ListenAndServe(":8888", nil)
}
