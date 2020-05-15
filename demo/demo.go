package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c)

	for cs := range c {
		switch cs {
		case os.Interrupt:
			fmt.Println("enen")
			close(c)
		default:
			fmt.Println("default")
		}
	}
}
