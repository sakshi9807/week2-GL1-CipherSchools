package main

import (
	"fmt"
	"time"
)

func sendInt(ch chan int) {
	ch <- 1
}
func sendBool(ch chan bool) {
	ch <- true
}
func main321() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go sendBool(ch2)
	go sendInt(ch1)
	select {
	case getInt := <-ch1:
		fmt.Println(getInt)
	case getBool := <-ch2:
		fmt.Println(getBool)
		// default:
		// 	fmt.Println("bye")
	}
}
// func main(){
// 	channel := make(chan int)
// 	go func(){
// 		channel<-1
// 		time.Sleep(time.Second)
// 		channel <- 2
// 		close(channel)
// 	}()
// 	for ch:= range channel{
// 		fmt.Println(ch)
// 	}
// }

// func main() {
// 	helloCh := make(chan string, 1)
// 	goodByeCh := make(chan string, 1)
// 	quitCh := make(chan bool)
// 	go receiveMessage(helloCh, goodByeCh, quitCh)
// 	go sendMessage(helloCh, "hello world")
// 	time.Sleep(time.Second)
// 	go sendMessage(helloCh, "Good bye world")
// 	result := <-quitCh
// 	fmt.Println("message from quiteCh", result)

// }
func sendMessage(ch chan<- string, message string) {
	ch <- message
}
func receiveMessage(helloCh, goodByeCh <-chan string, quiteCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh: ", message)
		case message := <-goodByeCh:
			fmt.Println("Message from helloCh: ", message)
		case <-time.After(time.Second * 2):
			fmt.Println("Nothing receiving in 2 second: Exiting thereceiveMessage goroutine")
			quiteCh <- true
			break
		}
	}
}
