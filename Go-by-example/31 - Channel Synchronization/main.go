package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	fmt.Println("Starting main execution...")

	done := make(chan bool, 1)
	go worker(done)

	<-done

	fmt.Println("im actually done")
}
