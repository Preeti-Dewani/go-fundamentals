package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// CPU Bound parallelism and WaitGroups
// Extend the Exercise 1 to initialze `num := 16`. Loop over an array of `num`
// random numbers calling `cpuBound(ix)`, and on Ctrl-C each method prints `fmt.Sprintf("Goodbye %d", n)`

const num = 16
var wg1 sync.WaitGroup
var arr  = [num]int{1, 2, 3 ,4, 5, 6, 7, 8, 9, 10}

func cpuBound1(in int) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer wg1.Done()

	<-c
	fmt.Println("Goodbye..", arr[in])
}

func main(){

	for ix, _ := range arr {
		wg1.Add(1)
		go cpuBound1(ix)
	}

	wg1.Wait()
}