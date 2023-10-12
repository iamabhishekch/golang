package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Study of Struct")

	Bittu := User{"Abhishek", 23, true, "chaurasiyaa750@go.dev"}
	fmt.Println(Bittu)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}
