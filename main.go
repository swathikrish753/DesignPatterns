package main

import (
	"fmt"
	"singleton-demo/singleton"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			db := singleton.GetInstance()

			fmt.Printf("Goroutine %d -> %p\n", id, db)
		}(i)
	}

	wg.Wait()
}
