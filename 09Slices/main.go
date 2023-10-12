package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"oranges", "apple", "banana"} // by not assigning lenght in this way the array automaticlly gets req. lenght
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Peach") // here we append or add in the array
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) /* here we want to show the array from 1 to 3 which starts from 1 and ends at 2 because last disit
	is always non-inclusive means  if we write 1:3 it show values form 1 to 2. */

	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 245
	highscore[1] = 800
	highscore[2] = 345
	highscore[3] = 445

	fmt.Println(highscore) // if we give more values then it gives error b/c length exceeds the limit

	highscore = append(highscore, 555, 675, 785) // we can add more then alloted lenght via this method.
	fmt.Println(highscore)

	// to sort the arrays

	sort.Ints(highscore)
	fmt.Println(highscore)
	// check if arrays are sorted or not

	fmt.Println(sort.IntsAreSorted(highscore)) // it will give value in true or false.

	// remove values from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	// see again

}
