package main

import (
	"vue-admin-element/initialize"
	// "log"
)

func main() {
	initialize.Run()
}

// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func main() {
// 	password := "admin"
// 	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

// 	fmt.Println("Password:", password)
// 	fmt.Println("Hash:    ", hash)
// 	hash = "$2a$14$0GhZwfu42R3whHYBv5mmZen7Yk1bty85UOgCPNG.BKMxy85DZWV82"

// 	match := CheckPasswordHash(password, hash)
// 	fmt.Println("Match:   ", match)
// }
