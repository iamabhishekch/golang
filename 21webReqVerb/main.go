package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome in web request verb code")
	PerformGetRequest()

}

func PerformGetRequest() {

	const myurl = "http://localhost:8000/post"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)
	var resonseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := resonseString.Write(content)

	fmt.Println("Bytecount is:", byteCount) // here watching what's inside the bytecount and it just  lenth
	fmt.Println(resonseString.String()) // by this wil have our raw data with us 


	//fmt.Println(content) // the content is always come in bytes format so we have to print that in string format

	// fmt.Println(string(content))
}
