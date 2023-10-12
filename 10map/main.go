package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Study of Map")

	languages := make(map[string]string) // we use this as map goes like make then kya bana hai map [key] values on our case both are string

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("JS stand for: ", languages["JS"])

	// to delete a line

	delete(languages, "RB")
	fmt.Println(languages)

}
