package main

import (
	"log"
	"sync"
	"time"
)

// Fan-in
// For an array of 5 random numbers between 0 and 10; dispatch mathBound(n)
// and append the output of each execution to a map of {"<n>": <n's_output>}
// The program must be executed with -race The program must exit within max(n) + 1 seconds.
// Not using a Mutex

const arrSize73 = 5
var w73 sync.WaitGroup

func mathBound73(n int) int {
	time.Sleep(2 * time.Second)
	return n * 8
}

func main(){
	arr := [arrSize73]int{8, 4, 3, 6, 1}
	dict := map[int]int{}

	done := make(chan bool, 1)
	chanMap := make(chan map[int]int, arrSize73)

	go func(){
		defer func(){
			done <- true
		}()

		for i := range chanMap {
			for ix, val := range i{
				dict[ix] = val
			}
		}
	}()

	for ix, n := range arr{
		w73.Add(1)
		go func(ix, n int){
			defer w73.Done()
			val := mathBound73(n)
			chanMap <- map[int]int{ix: val}
		}(ix, n)
	}

	w73.Wait()
	close(chanMap)

	<- done

	for i, n := range dict {
		log.Println("dict..", i, n)
	}
}