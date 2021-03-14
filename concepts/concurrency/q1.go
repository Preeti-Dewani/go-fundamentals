package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Goroutine and signals
// Write a Golang binary that calls the method cpuBound with argument 0, and exits on Ctrl-C with a message "`Goodbye`"

func cpuBound(n int) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {

		fmt.Fprintf(f, ".")
		<- c
		fmt.Println("GoodBye")
		os.Exit(n)

	}

}

func mainv1() {
	cpuBound(0)
}