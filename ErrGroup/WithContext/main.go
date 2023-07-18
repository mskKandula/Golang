package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		select {
		case <-time.After(3 * time.Second):
			fmt.Println("In 3")

		case <-ctx.Done():
			return ctx.Err()
		}

		return nil

	})

	g.Go(func() error {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("In 1")
			return errors.New("first error")

		case <-ctx.Done():
			return ctx.Err()
		}
	})

	g.Go(func() error {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("In 2")

		case <-ctx.Done():
			return ctx.Err()
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}

}
