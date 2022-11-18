package main

import (
	"fmt"
)

type S struct {
	a, b int
}

func (s *S) String() string {
	return fmt.Sprintf("%s", s)
}

func main() {
	s := &S{3, 4}
	fmt.Println(s)
}
