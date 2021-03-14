package main

import (
	"log"
	"sync"
	"time"
)

// Select on Channels and Ticker
// Spawn a goroutine which after 10 seconds, notifies all the 4 executions to stop printin using a boolean done channel as an input.

const spawnNo = 4

func main(){
	TimeTicker := time.NewTicker(1 * time.Second)

	wg := new(sync.WaitGroup)
	done := make(chan bool)

	for i := 0; i < spawnNo; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					log.Println("running for..", i)
				}
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
	TimeTicker.Stop()
	close(done)
	wg.Wait()

	log.Println("Time for running ticker is completed")
}