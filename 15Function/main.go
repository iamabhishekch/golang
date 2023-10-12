package main

import "fmt"

func main() {
	fmt.Println("Welcome to the study of Function")
	greeter()
	greeteeTwo()

	result :=  adder(5, 3) // to understand this do it step by step first create function and then call function than assign strign and than print that string
	fmt.Println(result)
}
func greeter() {
	fmt.Println("namste from go lang ")
	// we can't create function inside the function
}

func greeteeTwo() {
	fmt.Println("another namste from go lang")
}

func adder(valOne int, valTwo int) int { // we always have to mention the signature means what type of val is that int or string
	return valOne + valTwo
}
