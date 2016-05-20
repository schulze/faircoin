package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b := make([]byte, 1)
	n, err := rand.Read(b)
	if err != nil || n != 1 {
		panic(err)
	}
	winner := b[0] >> 7
	fmt.Println(winner)
}

