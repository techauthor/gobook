package main

import (
	"fmt"
	"reflect"
)

func hello(t interface{}) {

	if _, ok := t.(*Named); ok {
		fmt.Println("oh.yeah...")
	} else {
		fmt.Println("false...", reflect.TypeOf(t))
	}
}

type Named interface {
	GetName() string
}

type f1 func(string)

type Student struct {
	Name string
}

func (s *Student) GetName() string {
	return s.Name
}

func main() {

	//var s Named

	s := Student{
		Name: "kook",
	}

	fmt.Println(s.GetName())

}
