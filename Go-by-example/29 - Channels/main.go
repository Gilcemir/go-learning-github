package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping10" }()
	go func() { messages <- "ping" }()

	msg := <-messages
	ms2 := <-messages
	fmt.Println(msg)
	fmt.Println(ms2)
}
