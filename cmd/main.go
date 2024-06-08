package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MyPrintln(id int, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Xin chào, tôi là goroutine: ", id)
}

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func producer(factor int, c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * factor
	}
}

func consumer(wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("Received: %d\n", i)
	}
}

func worker(queue <-chan int, workerID int, done chan bool, killSignal chan bool) {
	for true {
		select {
		case k := <-queue:
			fmt.Println("Worker", workerID, "processing job", k)
			done <- true
		case <-killSignal:
			fmt.Println("Worker", workerID, "exiting")
			return
		}
	}
}

func searchingBy(server string) string {
	return fmt.Sprintf("Searching by %s\n", server)
}

func Worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("Hello")
		case <-ctx.Done():
			fmt.Println("Worker done")
			return ctx.Err()
		}
	}
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	recover()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Worker(ctx, &wg)
	}
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
