package main

import "fmt"

func main() {
	fmt.Println("welcome to the study of loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])

	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and day is %v", index, day)
	// }
	rougueValue := 1

	
		
	

	for rougueValue < 10 {

		if rougueValue == 5 +1 /* 5 will not count it will print only 4 so i did +1*/  {
			break
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

}
