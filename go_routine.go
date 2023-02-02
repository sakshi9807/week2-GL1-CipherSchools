package main

import (
	"fmt"
	"sync"
)

//var wait sync.WaitGroup

func main() {
	var wait sync.WaitGroup
	counter := 20000
	wait.Add(counter)
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()
	// wait.Add(2)
	// go hello()
	// go hello()
	// wait.Wait()
	//time.Sleep(time.Second * 5)
}
func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter)
	wait.Done()
}

// func hello() {
// 	fmt.Println(1)
// 	go func() {
// 		fmt.Println(19)
// 		fmt.Println(13)
// 		fmt.Println(14)
// 	}()
// }
