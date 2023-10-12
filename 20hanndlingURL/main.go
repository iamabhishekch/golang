package main

import (
	"fmt"
	"net/url"
)

const myURL string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to the study of handling URL ")
	fmt.Println(myURL)

	// Parsing
	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port()) // this is method not properties
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type of the query are: %T\n", qparams)

	fmt.Println(qparams["paymentid"])

	partsofURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=abhishek",
	}
	anotherURL := partsofURL.String()
	fmt.Println(anotherURL)

}
