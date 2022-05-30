package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := []byte("pa$$W0rd!")

	//for i := 0; i < 10; i++ {
	//	steps := 10
	//	hashed, _ := bcrypt.GenerateFromPassword(password, steps)
	//	fmt.Printf("%s: %s\n", password, hashed)
	//}

	hashedPassword := []byte("$2a$10$k.axZGHq8ae/ZhvGUVebyuaXsF9iVlSzoKGayoXb8flol/J1SZ0/K")
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err)

}
