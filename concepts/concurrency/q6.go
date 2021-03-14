package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Select on multi channels
// On the lines of exercise 2, dispatch 8 executions of the cpuBound function
// which either exits on Ctrl-C or after 10 seconds. And prints a different
// message for each kind of exit. I.e If exited via Ctrl-C should print
// "Caught interrupt" and in case of timeout should print "Goodbye! timeout"

const spawnNo1 = 8

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) Run() {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			fmt.Fprintf(f, ".")
		}
	}
}

func (t *Task) Stop() {
	close(t.closed)
	t.wg.Wait()
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 2),
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	for ix := 0; ix < spawnNo1; ix++ {
		task.wg.Add(1)
		go func() { defer task.wg.Done(); task.Run() }()
	}

	select {
		case <-c:
			log.Printf("Caught interrupt")
			task.Stop()
		case <-time.After(10 * time.Second):
			log.Println("out of time :(")
			task.Stop()
	}
}
