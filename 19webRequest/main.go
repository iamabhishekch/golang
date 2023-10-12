package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("welcome to the study of Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T/n", response)

	defer response.Body.Close() // caller responsible for closing
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
