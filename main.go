package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var cpuNumber = runtime.NumCPU()

//Expected result: executing HTTP calls (IO operation) with goroutines is faster than without them as each one
// can be executed in a separate goroutine and we don't have to wait a previous call to finish to schedule a new
// one.
func main() {
	println("Number of CPUs: ", cpuNumber)
	httpCallsWithGoroutines()
	httpCallsWithoutGoroutines()
}

func httpCallsWithGoroutines() {
	fmt.Println("Starting calls with goroutines")
	start := time.Now()
	executeHttpCallsWithGoroutines()
	finish := time.Now()
	fmt.Println("Finished calls with goroutines", finish.Sub(start))
}

func executeHttpCallsWithGoroutines() {
	var waitForAll sync.WaitGroup
	waitForAll.Add(cpuNumber)
	for j := 0; j < cpuNumber; j++ {
		go func(wait *sync.WaitGroup) {
			doHttpCall()
			wait.Done()
		}(&waitForAll)
	}
	waitForAll.Wait()
}

func httpCallsWithoutGoroutines() {
	fmt.Println("Starting calls without goroutines")
	start := time.Now()
	executeHttpCallsWithoutGoroutines()
	finish := time.Now()
	fmt.Println("Finished calls without goroutines", finish.Sub(start))
}

func executeHttpCallsWithoutGoroutines() {
	for j := 0; j < cpuNumber; j++ {
		doHttpCall()
	}
}

func doHttpCall() {
	_, _ = http.Get("http://google.com/")
}
