package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	couter int
	wg     sync.WaitGroup

	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", couter)

}

func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {

		mutex.Lock()
		{
			value := couter

			runtime.Gosched()

			value++

			couter = value
		}
		mutex.Unlock()

	}
}
