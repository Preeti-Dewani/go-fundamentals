package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Context
// Repeat Exercise 3, but instead of using done channel, use a context.

func main(){

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					log.Println("running for..", i)
				}
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()

	log.Println("Time for running ticker is completed")
}