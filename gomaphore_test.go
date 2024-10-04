package gomaphore_test

import (
	"context"
	"github.com/zipkero/gomaphore"
	"testing"
	"time"
)

func TestNewGomapahore(t *testing.T) {
	g := gomaphore.New(2)
	if g == nil {
		t.Error("nil Gomaphore")
	}
}

func TestWaitAndRelease(t *testing.T) {
	g := gomaphore.New(2)

	g.Wait()
	g.Wait()

	go func() {
		time.Sleep(1 * time.Second)
		g.Release()
	}()

	g.Release()
}

func TestWaitWithTimeout(t *testing.T) {
	sem := gomaphore.New(1)
	sem.Wait()

	err := sem.WaitWithTimeout(100)
	if err == nil {
		t.Fatalf("expected timeout error")
	}
}

func TestWaitWithContext(t *testing.T) {
	sem := gomaphore.New(1)
	sem.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := sem.WaitWithContext(ctx)
	if err == nil {
		t.Fatalf("expected context error")
	}
}
