package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(phrase)
	// pass the data to channel
	doneChan <- true
	// close the channel in the slow operations
	close(doneChan)
}

//func main() {
//	done := make(chan bool)
//	go greet("Nice to meet you", done)
//	go greet("How's things going on", done)
//	go slowGreet("Shall we have a cofee?", done)
//	go slowGreet("mmm..., Ok :)!!!", done)
//	go greet(" let's go.. :)!", done)
//
//	//// data come out from channel
//	<-done
//	<-done
//	<-done
//	<-done
//	<-done
//}

//func main() {
//	dones := make([]chan bool, 5)
//	dones[0] = make(chan bool)
//	go greet("Nice to meet you", dones[0])
//	dones[1] = make(chan bool)
//	go greet("How's things going on", dones[1])
//	dones[2] = make(chan bool)
//	go slowGreet("Shall we have a cofee?", dones[2])
//	dones[3] = make(chan bool)
//	go slowGreet("mmm..., Ok :)!!!", dones[3])
//	dones[4] = make(chan bool)
//	go greet(" let's go.. :)!", dones[4])
//
//	for _, done := range dones {
//		<-done
//	}
//
//}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you", done)
	go greet("How's things going on", done)
	go slowGreet("Shall we have a cofee?", done)
	go slowGreet("mmm..., Ok :)!!!", done)
	go greet(" let's go.. :)!", done)

	// special feature by go
	//for donesChan := range done {
	//	fmt.Println(donesChan)
	//}

	for range done {
		<-done
	}
}
