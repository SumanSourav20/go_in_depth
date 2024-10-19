package main

import (
	"fmt"
	"time"
)

func printNum(num int) {
	fmt.Println(num)

}
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Working")
		}
	}
}
func main() {
	myChannel := make(chan int)
	anotherChannel := make(chan string)
	go func() {
		//printNum(5)
		myChannel <- 1
	}()

	go func() {
		printNum(10)
		anotherChannel <- "Another Channel"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

	// ch := <-myChannel

	// fmt.Println(ch)
	done := make(chan bool)

	go doWork(done)
	time.Sleep(time.Second * 2)
	close(done)
}
