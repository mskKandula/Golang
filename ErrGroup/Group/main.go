package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	eg := new(errgroup.Group)

	eg.Go(func() error {

		time.Sleep(3 * time.Second)

		fmt.Println("In 3")

		return nil
	})

	eg.Go(func() error {

		time.Sleep(1 * time.Second)

		fmt.Println("In 1")

		return errors.New("first error")
	})

	eg.Go(func() error {

		time.Sleep(2 * time.Second)

		fmt.Println("In 2")

		return errors.New("second error")

	})

	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

}
