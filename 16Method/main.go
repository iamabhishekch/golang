package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to the Study of Struct")

	Abhishek := User{"Abhishek", 23, true, "chaurasiyaa750@go.dev"}
	fmt.Println(Abhishek)

	Abhishek.GetStatus()
	Abhishek.GetMail()
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

/* what is method 
Methods are associated with a specific data type or struct 
Methods are called on an instance of a struct using the dot notation

*/

func (u User) GetStatus() {
	fmt.Println("is user is online: ", u.Status)

}

func (u User) GetMail() { 
	u.Email = "bittu.go.dev" // by this way we are manupulating the email but not the original email it's a copy
	fmt.Println("The email of user is: ", u.Email)

}
