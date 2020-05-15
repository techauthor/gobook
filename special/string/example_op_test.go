package string

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

//https://hermanschaaf.com/efficient-string-concatenation-in-go/

//主要结论
//
//在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
//在一些性能要求较高的场合，尽量使用builder以获得更好的性能
//较少字符串连接的场景下性能最好，而且代码更简短清晰，可读性更好
//如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf

//golang 里面的字符串都是不可变的，每次运算都会产生一个新的字符串，
// 所以会产生很多临时的无用的字符串，不仅没有用，还会给 gc 带来额外的负担，所以性能比较差
func ExampleAdd() {
	hello := "hello"
	world := "world"
	for i := 0; i < 10; i++ {
		var str string
		str += hello + "," + world
		fmt.Println(str)
	}
	//output:
}

func ExampleAppend() {
	hello := "hello"
	world := "world"
	s := []string{hello}
	s = append(s, world)
	fmt.Println(s)
	//output:
}

//内部使用 []byte 实现，不像直接运算符这种会产生很多临时的字符串，
// 但是内部的逻辑比较复杂，有很多额外的判断，还用到了 interface，所以性能也不是很好
func ExampleSprintf() {
	hello := "hello"
	world := "world"
	for i := 0; i < 10; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
	//output:
}

//join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入，
// 在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
func ExampleJoin() {
	hello := "hello"
	world := "world"
	for i := 0; i < 10; i++ {
		n := strings.Join([]string{hello, world}, "")
		fmt.Println(n)
	}
	//output:
}

//这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
func ExampleBufferWrite() {
	hello := "hello"
	world := "world"
	for i := 0; i < 10; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
	//output:
}

// WriteString有对b.buf进行append操作，那对于长的字符串就会触发扩容操作影响性能
// 可先调用Grow方法，事先分配好所需的容量
func ExampleBuilderWrite() {
	var s strings.Builder
	s.Grow(100) // 具体情况计算好，这里随便写的
	s.WriteString("社会主义核心价值观的基本内容:")
	s.WriteString("富强、民主、文明、和谐，是我国社会主义现代化国家的建设目标;")
	s.WriteString("自由、平等、公正、法治，是对美好社会的生动表述;")
	s.WriteString("爱国、敬业、诚信、友善”，是公民基本道德规范。")

	fmt.Println(s.String())
	//output:
}

func ExampleCopy() {
	str := "chinese"
	city := "beijing"
	s := time.Now()
	zstr := []byte(str)
	for i := 0; i < 100000; i++ {
		copy(zstr, city)
	}
	e := time.Since(s)
	fmt.Println("time cost 5:", e, len(zstr))
	//output:
}
