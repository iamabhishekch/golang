package main

import "fmt"

func main() {
	fmt.Println("welcome to the stufy of array")

	var fruitList [4]string // [4] is the lenght of the array

	fruitList[0] = "apple"
	fruitList[1] = "Orange"
	fruitList[3] = "peach"

	fmt.Println("this is the list of fruit", fruitList)
	/* "this is the list of fruit [apple Orange  peach]" this is the output that will get  here after the orange
	we can see there is two space becasue we gave element is allowed for 4 and we don't enter the 4th element we directly
	jump to last element via entering [3]. */

	fmt.Println("this is the length of the array", len(fruitList)) // byu this syntax we can see the length of array
	// the length will always get from the aassigned length at the beginning. like we assigned 4 and gave element only 3 but the len syntax wil show 4 is the lenght

	// Assigning value directly

	var vegyList = [3]string{"spinach", "coriander leaf", "brinjal"} // this is how we direclty assign the values
	fmt.Println("vegyList itemis are", vegyList)

}
