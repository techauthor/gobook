package json

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func Example() {
	s := []student{
		{
			Name: "txvier",
			Age:  38,
		},
		{
			Name: "kook",
			Age:  22,
		},
	}

	r, _ := json.MarshalIndent(s, "", "\t")

	fmt.Println(string(r))

	// Output:
	//[
	//	{
	//		"Name": "txvier",
	//		"Age": 38
	//	},
	//	{
	//		"Name": "kook",
	//		"Age": 22
	//	}
	//]

}
