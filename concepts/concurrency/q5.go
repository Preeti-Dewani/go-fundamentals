package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Pool
// Using learning of Exercise 2 and 3, Dispatch 10 parallel executions of the
// timeBound(10) method. The program should exit in ~ 10 seconds.
// Alter the code to allow maximum 2 parallel executions. The program should now take ~ 50 seconds

const size = 2
var wait sync.WaitGroup

func initPool(size int) chan bool{
	channels := make(chan bool, size)
	for i :=0; i < size; i++ {
		channels <- true
	}

	return channels
}

func timeBound(n int, channels chan bool) {
	<- channels
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}

	defer func(){
		channels <- true
	}()

	defer wait.Done()
	defer f.Close()

	for x := 0; x < n; x++ {
		fmt.Fprintf(f, ".")
		time.Sleep(1 * time.Second)
	}
}

func main(){
	channels := initPool(size)

	for ix := 0; ix < 10; ix++ {
		wait.Add(1)
		go timeBound(10, channels)
	}

	wait.Wait()
}