package main

import (
	"fmt"
)

func sayHello(name string, ch chan string) {
	ch <- "hello " + name
}

func main() {
	ch := make(chan string)

	go sayHello("brian", ch)
	go sayHello("kyaw", ch)
	go sayHello("zin", ch)
	go sayHello("thant", ch)
	go sayHello("min", ch)

	for range 5 {
		msg := <-ch
		fmt.Println(msg)
	}
}
