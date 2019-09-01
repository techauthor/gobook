package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "你好吗?"
	b := make([]byte, len(s))
	r := strings.NewReader(s)
	n, _ := r.Read(b)
	fmt.Println(n, string(b))

	b2 := make([]byte, len(s))
	r2 := bytes.NewReader([]byte(s))
	n2, _ := r2.Read(b2)
	fmt.Println(n2, string(b2))

	r3 := bufio.NewReader(strings.NewReader(s))
	line, _, _ := r3.ReadLine()
	fmt.Println(string(line))

	//writer
	w := bytes.NewBuffer([]byte("kao"))
	w.WriteString(" ni ma bf")
	fmt.Println(string(w.Bytes()))

	w2 := bufio.NewWriter(os.Stdout)
	w2.WriteString("hello")
	w2.Flush()
	fmt.Println(string(w2.Flush()))
}
