package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	winner := b[0] >> 7
	fmt.Println(winner)
}

