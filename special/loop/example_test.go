package special_loop

import (
	"fmt"
	"sort"
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
	var s []int
	for key := range m {
		s = append(s, key)
	}
	sort.Ints(s)
	for _, v := range s {
		fmt.Printf("key is:[%d] and the value is:[%s]\n", v, m[v])
	}

	// Output:
	//key is:[0] and the value is:[0]
	//key is:[1] and the value is:[1]
}

func ExampleForRangeSlice() {
	s := []int{0, 1, 2}
	for index, value := range s { // the variable `index` can be ignored like this `for _ , value := range s`
		fmt.Printf("index is:[%d] and the value is:[%d]\n", index, value)
	}
	// Output:
	// index is:[0] and the value is:[0]
	// index is:[1] and the value is:[1]
	// index is:[2] and the value is:[2]
}

func ExampleForChannel() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	for index := range ch { //the variable `now` can be omitted if it's unused,like this:`for range tick`
		fmt.Println(index)
	}
	// Output:
	//0
	//1
	//2
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
	ch := make(chan string)
	go func() {
		ch <- "Done"
	}()
	for {
		// your loop code
		fmt.Println(<-ch)
		break
	}
	// Output:
	// Done
}
