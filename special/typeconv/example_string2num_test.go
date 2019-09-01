package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	b := strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(b))
}
