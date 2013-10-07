package main

import "fmt"
import "strconv"
import "time"

func main() {

	jobs := make(chan string, 200)

	// consumer
	go func() {
		for {
			job := <- jobs
			fmt.Println("Got job:",job)
		}
	}()

	// producer
	i := int64(0)
	for{
		jobs <- strconv.FormatInt(i,10)
		time.Sleep(time.Second * 5)
		i++
	}

}
