package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	chanRandom := make(chan int, 10)
	chanSquares := make(chan int, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			chanRandom <- rnd.Intn(101)
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
