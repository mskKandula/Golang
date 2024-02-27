package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Text string
}

func main() {

	var test Test
	str := `{"Text":"What is the output of this C code? #include <stdio.h> int main()"}`

	json.MarshalIndent(str, "")

	json.UnmarshalIndent([]byte(str), &test)

	fmt.Println(test)
}
