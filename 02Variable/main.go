package main

import (
	"fmt"
)

func main() {
	var username string = "abhishek"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVall uint16 = 255 // here if we write 256 than will have error becaue uint is not allowed value more than 255. for more go through the doc.
	fmt.Println(smallVall)
	fmt.Printf("Variable is of type: %T\n", smallVall)

	var smallFLoat float32 = 255.4567891234 // float 32 will get only 5 digit after decimal for more  float64 are req.
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T\n", smallFLoat)

	// default values and some aliases
	var anotherVarriable int // of we don't initialize any values in int then the default value will be zerl (0).
	fmt.Println(anotherVarriable)
	fmt.Printf("varriable is of type: %T\n", anotherVarriable)

	// implicit type // in this compiler automatically assign the type of variable to the var.
	var website = "here we are not writing the type of variable"
	fmt.Println(website)

	// no var style

	numberOfUser :=  /*see the operator*/ 30000.0 // like this we are allowed to do but only in methods like func main () not outside the method.
	fmt.Println(numberOfUser)
}
