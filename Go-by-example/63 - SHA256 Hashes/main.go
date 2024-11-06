package main

import (
	"crypto/sha256"
	"fmt"
)

// alterado
func main() {
	s := "sha256 this string"

	salt := []byte("salt")

	fmt.Println("Salt :", salt)

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(salt)

	fmt.Println("Salt after sum:", salt)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
