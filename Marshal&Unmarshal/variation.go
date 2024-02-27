package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Test struct {
	Text string
}

func main() {

	var test Test
	str := []byte(`{"Text":"What is the output of this C code?\n#include <stdio.h>\nint main()"}`)

	reader := bytes.NewReader(str)

	json.NewDecoder(reader).Decode(&test)

	fmt.Println(test)
}
