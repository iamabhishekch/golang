package main

import "fmt"

func main() {
	// when we use defer keyword before any line it will be executed at the last line of that function
	// for this we can imagine - line number 8 will be executed at  last line (just before baract of the function)
	defer fmt.Println("World ")
	defer fmt.Println("one ")
	defer fmt.Println("two")
	// in case of many defers they excution will revers like output that will be :- hello - two - one - world (it gets reversed)
	fmt.Println("Hello")

	 myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)  
    
		/* Notice one thing here there are five iterations in the loop, means we will have five deferred calls to fmt.Println(i)
		means LIFO(Last In First Out) formula will come so it doesn't print anything immediately; instead, it schedules the fmt.Println(i)
		and it will print in reverse order (4,3,2,1). */
		

	}

}
