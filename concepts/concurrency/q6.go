package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Select on multi channels
// On the lines of exercise 2, dispatch 8 executions of the cpuBound function
// which either exits on Ctrl-C or after 10 seconds. And prints a different
// message for each kind of exit. I.e If exited via Ctrl-C should print
// "Caught interrupt" and in case of timeout should print "Goodbye! timeout"

var wgQ6 sync.WaitGroup

func cpuBoundV2(n int, ticker *time.Ticker) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}

	defer wgQ6.Done()
	defer f.Close()
	for {
		select {
		case <- ticker.C:
			fmt.Println("Goodbye! timeout")
			os.Exit(n)
		default:
			fmt.Fprintf(f, ".")
		}
	}
}

func main(){

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	timeTicker := time.NewTicker(5 * time.Second)

	wgQ6.Add(1)
	go func(){
		for {
			select {
			case <-c:
				fmt.Println("Caught interrupt")
				timeTicker.Stop()
				return
			}
		}
	}()

	for ix := 0; ix < 8; ix++ {
		wgQ6.Add(1)
		go cpuBoundV2(ix, timeTicker)
	}

	time.Sleep(10 * time.Second)
	timeTicker.Stop()
	wgQ6.Wait()
}