package main

import (
	"context"
	"github.com/zipkero/gomaphore"
	"log"
	"sync"
	"time"
)

func main() {
	sem := gomaphore.New(2)

	go func() {
		sem.Wait()
		defer sem.Release()

		time.Sleep(2 * time.Second)
	}()

	go func() {
		sem.Wait()
		defer sem.Release()
		time.Sleep(1 * time.Second)
	}()

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			if err := sem.WaitWithContext(ctx); err != nil {
				log.Printf("worker %d: %v\n", i, err)
				return
			}
			defer sem.Release()

			time.Sleep(500 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	sem.Close()
}
