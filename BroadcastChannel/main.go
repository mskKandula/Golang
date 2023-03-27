package main

import (
	"context"
	"fmt"
	"sync"
)

type BroadcastServer struct {
	Source         <-chan int
	Listeners      []chan int
	AddListener    chan chan int
	RemoveListener chan (<-chan int)
}

func (s *BroadcastServer) Subscribe() <-chan int {
	newListener := make(chan int)
	s.AddListener <- newListener
	return newListener
}

func (s *BroadcastServer) CancelSubscription(channel <-chan int) {
	s.RemoveListener <- channel
}

func NewBroadcastServer(ctx context.Context, source <-chan int) *BroadcastServer {
	service := &BroadcastServer{
		Source:         source,
		Listeners:      make([]chan int, 0),
		AddListener:    make(chan chan int),
		RemoveListener: make(chan (<-chan int)),
	}
	go service.Serve(ctx)
	return service
}

func (s *BroadcastServer) Serve(ctx context.Context) {
	defer func() {
		for _, listener := range s.Listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newListener := <-s.AddListener:
			s.Listeners = append(s.Listeners, newListener)
		case listenerToRemove := <-s.RemoveListener:
			for i, ch := range s.Listeners {
				if ch == listenerToRemove {
					s.Listeners[i] = s.Listeners[len(s.Listeners)-1]
					s.Listeners = s.Listeners[:len(s.Listeners)-1]
					close(ch)
					break
				}
			}
		case val, ok := <-s.Source:
			if !ok {
				return
			}
			for _, listener := range s.Listeners {
				if listener != nil {
					select {
					case listener <- val:
					case <-ctx.Done():
						return
					}

				}
			}
		}
	}

}

func rangeChannel(ctx context.Context, n int) <-chan int {
	valueStream := make(chan int)
	go func() {
		defer close(valueStream)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case valueStream <- i:
			}
		}
	}()
	return valueStream
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Generates a channel sending integers
	// From 0 to 9
	range10 := rangeChannel(ctx, 10)

	broadcaster := NewBroadcastServer(ctx, range10)
	listener1 := broadcaster.Subscribe()
	listener2 := broadcaster.Subscribe()
	listener3 := broadcaster.Subscribe()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := range listener1 {
			fmt.Printf("Listener 1: %v/10 \n", i+1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range listener2 {
			fmt.Printf("Listener 2: %v/10 \n", i+1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range listener3 {
			fmt.Printf("Listener 3: %v/10 \n", i+1)
		}
	}()

	wg.Wait()
}
