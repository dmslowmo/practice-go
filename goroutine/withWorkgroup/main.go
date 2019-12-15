package main

/* producer-consumer problem in Go */

import (
	"fmt"
	"sync"
	"time"
)

var producerCount = 2
var done = make(chan string, producerCount)
var cDone = make(chan string)
var msgs = make(chan int, 20)

func produce1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i < 20; i++ {
		msgs <- i
		time.Sleep(10 * time.Microsecond)
	}
	done <- "Producer 1 is done"
}

func produce2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 20; i < 30; i++ {
		msgs <- i
		time.Sleep(10 * time.Microsecond)
	}
	done <- "Producer 2 is done"
}

func consume() {
	for {
		msg, ok := <-msgs
		if !ok {
			break
		}
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
	cDone <- "Consumer is done"
}

func workerMonitor(wg *sync.WaitGroup) {
	wg.Wait()
	close(msgs)
	close(done)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go produce1(wg)
	wg.Add(1)
	go produce2(wg)
	go workerMonitor(wg)
	go consume()
	for i:=0;i<producerCount;i++ {
		fmt.Println(<-done)
	}
	fmt.Println(<-cDone)
}
