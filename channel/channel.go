package main

import (
	"fmt"
)

func receive(ch <-chan int) {
	for i := range ch {
		fmt.Println("接收", i)
	}
}

func send(ch chan<- int, data int) {
	ch <- data
	fmt.Println("发送", data)

}

func main() {
	ch := make(chan int)

	go receive(ch)

	func() {
		for i := 0; i <= 1000; i++ {
			send(ch, i)
		}
	}()

}
