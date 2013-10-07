package main

import (
	"fmt"
	"strconv" // added for giving proper name to job
	"time"    // added for delaying consecutive inserts into job queue
)

func main() {

	jobQueue := make(chan string, 200)
	done := make(chan bool,1)
	// consumer
	go func() {
		for {
			job := <-jobQueue
			fmt.Println("Got job:", job)
		}
		done <- true
	}()
	// producer
	go func() {
		i := int64(0)
		for {
			jobQueue <- strconv.FormatInt(i, 10)
			time.Sleep(time.Second * 5) // to wait for 5 seconds before adding one more job
			i++
		}
		done <- true
	}()
	<-done
}
