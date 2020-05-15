package string

import (
	"fmt"
	"strings"
)

func ExampleReplaceAll() {
	s := "hel/lo wo/rl/d"
	key := strings.ReplaceAll(s, "/", "")
	fmt.Println(key)

	// Output:
	// hello world
}
