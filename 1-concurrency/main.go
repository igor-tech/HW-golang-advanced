package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chanRandom := make(chan int, 10)
	chanSquares := make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			chanRandom <- rand.Intn(101)
		}

		close(chanRandom)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range chanRandom {
			squared := v * v
			chanSquares <- squared
		}

		close(chanSquares)
	}()

	wg.Wait()

	for v := range chanSquares {
		fmt.Printf("%d ", v)
	}
}
