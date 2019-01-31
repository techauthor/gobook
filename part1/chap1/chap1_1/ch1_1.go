//ch1_1.go
package chap1_1

import (
	"fmt"
	"net/http"
)

func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Golang")
}

func main() {
	fmt.Println("hello ch1_1...")
	http.HandleFunc("/hello", HelloGo)
	http.ListenAndServe(":8888", nil)
}