package main

import (
	"fmt"
	"sync"
)

func main() {
	var c = make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
	// var chanels sync.WaitGroup
	// chanels.Add(10)
	// for i := 0; i < 10; i++ {
	// 	// go Run(i, &chanels)
	// }
	// chanels.Wait()
}

func Run(i int, group *sync.WaitGroup) {
	fmt.Printf("这个值是%v\n", i)
	group.Done()
}
