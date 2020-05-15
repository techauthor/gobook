package json

import (
	jiter "github.com/json-iterator/go"
)

var jsoniter = jiter.ConfigCompatibleWithStandardLibrary

func ExampleJsoniter() {
	intput := struct {
		Name string
		Age  int
	}{
		Name: "txvier", Age: 50,
	}
	jsoniter.Marshal(intput)
	// Output:
}
