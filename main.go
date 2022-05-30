package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i := 0; i < 10; i++ {
		pass := []byte("pa$$W0rd!")
		cost := 10
		hash, _ := bcrypt.GenerateFromPassword(pass, cost)

		fmt.Printf("%s: %s\n", pass, hash)
	}
}
