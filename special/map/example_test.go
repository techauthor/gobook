package examplemap

import "fmt"

type Student struct {
	Name string
	Age  int
}

//map的基本使用
func Example1() {
	//make初始化
	m := make(map[string]int)
	//由于map是无序的，先针对key创建一个有序的slice
	k := []string{"1", "2", "3", "4", "5"}
	m[k[0]] = 1
	m[k[1]] = 2
	m[k[2]] = 3
	m[k[3]] = 4
	m[k[4]] = 5

	for _, v := range k {
		fmt.Println(v, m[v])
	}

	// output:
	// 1 1
	// 2 2
	// 3 3
	// 4 4
	// 5 5
}

//map的删除
func Example2() {
	//字面量初始化
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	//删除后对应元素为储存类型的零值
	delete(m, "3")

	fmt.Println(m["1"])
	fmt.Println(m["2"])
	fmt.Println(m["3"])
	fmt.Println(m["4"])
	fmt.Println(m["5"])

	// output:
	// 1
	// 2
	// 0
	// 4
	// 5

}

//map的key是否存在
func Example3() {
	//字面量初始化
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	delete(m, "3")

	if v, ok := m["3"]; !ok {
		//ok为false，代表的是对应的key不存在，和存在但是值为0是两回事。
		//也是判断对应key是否存在的唯一方法。
		fmt.Println(ok, v)
	}

	//output:
	//false 0
}
