## Semaphore in Go

This is a simple example of how to use semaphore in Go.

- Sample code
```go
package main

import (
	"github.com/zipkero/gomaphore"
	"sync"
)

func main() {
	sem := gomaphore.New(2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
			defer wg.Done()
						
            sem.Wait()
            defer sem.Release()
        }(i)
    }
	
	wg.Wait()
	sem.Close()
}
```

- Timeout Sample code
```go
package main

import (
	"github.com/zipkero/gomaphore"
	"log"
	"sync"
	"time"
)

func main() {
	sem := gomaphore.New(1)

	go func() {
		sem.Wait()
		defer sem.Release()

		time.Sleep(1000 * time.Millisecond)
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := sem.WaitWithTimeout(100); err != nil {
				log.Println(err)
				return
			}
			defer sem.Release()
		}(i)
	}

	wg.Wait()
	sem.Close()
}
```

- Context Sample code
```go
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
```
