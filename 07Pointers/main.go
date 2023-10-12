package main

import "fmt"

func main() {
	fmt.Println("Study of Pointers")

	// var ptr *int // this how we decare a pointer
	// fmt.Println("the value of pointer is", ptr)

	myNumber := 23 // from the output we got "0xc00000a0b8" this the addres of memory 23 where the value of myNumber is stored in.

	var ptr = &myNumber // for the direct pointer we use "*" and for the reffrence pointer we use "&"

	fmt.Println("The value of pointer is", ptr)
	fmt.Println("The value of pointer is", *ptr) // for this * will get the value of myNumber which is 23.

	*ptr = *ptr + 2
	fmt.Println("the value of new pinter is", *ptr) 

}
