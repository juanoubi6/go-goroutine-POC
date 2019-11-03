package main

import (
	"fmt"
	"time"
  "net/http"
  "sync"
)

//Expected result: executing HTTP calls with goroutines is faster than withouth them
func main() {
  httpCallsWithGoroutines()
  httpCallsWithoutGoroutines()
}

func httpCallsWithGoroutines(){
  fmt.Println("Starting calls with goroutines")
  start := time.Now()
  executeHttpCallsWithGoroutines()
  finish := time.Now()
  fmt.Println("Finished calls with goroutines", finish.Sub(start))
}

func executeHttpCallsWithGoroutines(){
    var waitForAll sync.WaitGroup
    waitForAll.Add(3)
    for j := 0; j <= 2; j++ {
        go func(wait *sync.WaitGroup){
          _, _ = http.Get("http://google.com/")
          wait.Done()
        }(&waitForAll)
    }
    waitForAll.Wait() 
}

func httpCallsWithoutGoroutines(){
  fmt.Println("Starting calls without goroutines")
  start := time.Now()
  executeHttpCallsWithoutGoroutines()
  finish := time.Now()
  fmt.Println("Finished calls without goroutines", finish.Sub(start))
}

func executeHttpCallsWithoutGoroutines(){
  _, _ = http.Get("http://google.com/")
  _, _ = http.Get("http://google.com/")
  _, _ = http.Get("http://google.com/")
}
