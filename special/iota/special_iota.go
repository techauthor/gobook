package main

import "fmt"

//case 1
type Weekday int

const (
	Sunday    Weekday = iota //0
	Monday                   //1
	Tuesday                  //2
	Wednesday                //3
	Thursday                 //4
	Friday                   //5
	Saturday                 //6
)

func ExampleCase1() {
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d\r\n", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	//Output:
	//0,1,2,3,4,5,6
}

//case2
const (
	a = iota //0
	b        //1
	c        //2
)

func ExampleCase2() {
	fmt.Printf("%d,%d,%d\r\n", a, b, c)
	//Output:
	//0,1,2
}

//case3
const (
	d = iota //0
	e = iota //1
	f = iota //2
)

func ExampleCase3() {
	fmt.Printf("%d,%d,%d\r\n", d, e, f)
	//Output:
	//0,1,2
}

//case 4
const (
	g = iota //0
	h = 3.14 //3.14
	i = iota //2
	j = iota //3
)

func ExampleCase4() {
	fmt.Printf("%d,%v,%d,%d\r\n", g, h, i, j)
	//Output:
	//0,3.14,2,3
}

//case5
const (
	k = iota //0
	_        //ignore
	l = 3.14 //3.14
	m = iota //3
	n        //4
)

func ExampleCase5() {
	fmt.Printf("%d,%v,%d,%d\r\n", k, l, m, n)
	//Output:
	//0,3.14,3,4
}

//case6
const (
	o, p = iota, iota //0,0
	q    = iota       //1
)

func ExampleCase6() {
	fmt.Printf("%d,%d,%d\r\n", o, p, q)
	//Output:
	//0,0,1
}

//case7
const r = iota //0
const s = iota //0
const t = iota //0

func ExampleCase7() {
	fmt.Printf("%d,%d,%d\r\n", r, s, t)
	//Output:
	//0,0,0
}

//case8
const (
	u = -1   //-1
	v = iota //1
	w        //2
)

func ExampleCase8() {
	fmt.Printf("%d,%d,%d\r\n", u, v, w)
	//Output:
	//-1,1,2
}

//case9
const (
	a1, a2, a3 = iota + 1, iota + 5, iota //1,2,0
	a4, a5, a6 = iota, iota, iota         //2,3,1
	a7, a8, a9                            //3,4,2
)

func ExampleCase9() {
	fmt.Printf("%d,%d,%d\r\n", a1, a2, a3)
	fmt.Printf("%d,%d,%d\r\n", a4, a5, a6)
	fmt.Printf("%d,%d,%d\r\n", a7, a8, a9)
	//Output:
	//1,2,2,3,3,4
}

func main() {
	ExampleCase9()
}
