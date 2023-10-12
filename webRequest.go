// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const url = "https://lco.dev"

// func main() {

// 	response, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close() // caller's responsilbilty to close the connection
// 	fmt.Printf("Type of response is: %T\n", response)

// }

package main

import (
	"fmt"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // Use defer to ensure the response body is closed when the function exits

	fmt.Printf("Type of response is: %T\n", response)
}
