package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Study of If else")
	loginCount := 13

	var result string

	if loginCount < 10 {
		result = "be more active"

	} else if loginCount <= 15 {
		result = "average user"
	} else {
		result = "regular coustmer"
		
	}
	fmt.Println(result)  // did a mistake here by writing print statement in the "else" curly brackets

	if num := 3; num < 10 {  // here we intialize num first then give condition 
		fmt.Println("number is less tha 10")

	} else {
		fmt.Println("numbr is greater than 10")
	}

}


