package main

import (
	"fmt"
	"time"
)

func ExampleFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", i)
	}
	fmt.Println()
	// Output:
	// 0,1,2,3,4,5,6,7,8,9,
	//
}

func ExampleForRangeMap() {
	m := map[int]string{
		0: "0",
		1: "1",
	}
	for key, value := range m {
		fmt.Printf("key is:[%d] and the value is:[%s]\n", key, value)
	}

	// Output:
	// key is:0 and the value is:0
	// key is:1 and the value is:1
}

func ExampleForRangeSlice() {
	s := []int{0, 1, 2}
	for index, value := range s { // the variable `index` can be ignored like this `for _ , value := range s`
		fmt.Printf("index is:[%d] and the value is:[%d]\n", index, value)
	}
	// Output:
	// index is:0 and the value is:0
	// index is:1 and the value is:1
	// index is:2 and the value is:2
}

func ExampleForChannel() {
	fmt.Println("print 3 timestamps like this:")
	tick := time.Tick(1 * time.Second)
	i := 0
	for now := range tick { //the variable `now` can be omitted if it's unused,like this:`for range tick`
		if i > 2 {
			break
		}
		fmt.Printf("now time is:%s\n", now)
		i++
	}
}

func ExampleForRangeIgnoreVar() {
	s := []int{0, 1, 2}
	for range s { //loop len(s) times.
		fmt.Println("foo...")
	}
	// Output:
	// foo...
	// foo...
	// foo...
}

func ExampleForDeadLoop() {
	for {
		fmt.Println("dead loop....")
	}
	// Output:
	// dead loop....
	// ...
}

func main() {
	//ExampleForChannel()
}
