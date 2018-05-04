package main

/* producer-consumer problem in Go */

import (
	"fmt"
	"time"
)

var producerCount = 2
var consumerCount = 1
var pDone = make(chan string, producerCount)
var cDone = make(chan string, consumerCount)
var msgs = make(chan int, 20)

func produce1() {
	for i := 10; i < 20; i++ {
		msgs <- i
		time.Sleep(10 * time.Microsecond)
	}
	pDone <- "Producer 1 is done"
}

func produce2() {
	for i := 20; i < 30; i++ {
		msgs <- i
		time.Sleep(10 * time.Microsecond)
	}
	pDone <- "Producer 2 is done"
}

func consume() {
	for {
		msg, ok := <-msgs
		if !ok {
			break
		}
		fmt.Println(msg)
		//simulate long processing
		time.Sleep(1 * time.Second)
	}
	cDone <- "Consumer is done"
}

func main() {
	go produce1()
	go produce2()
	go consume()
	for i:=0;i<producerCount;i++ {
		fmt.Println(<-pDone)
	}
	close(msgs)
	fmt.Println("msgs channel now closed")
	fmt.Println(<-cDone)
	close(cDone)
}
