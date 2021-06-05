package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {
    s := "Mohana Kandula"

	hashedString := hashing(s)

	fmt.Println(s)

    fmt.Printf("%x\n", hashedString)
	
}

func hashing(s string)[]byte{

	h := sha1.New()

	h.Write([]byte(s))

	hashedString := h.Sum(nil)

	return hashedString
}