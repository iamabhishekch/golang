package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in golang.")

	content := "This is the file thata needs to go at - xanderbilla.com"

	file, err := os.Create("./mygolang.txt")

	if err != nil {
		panic(err)

	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)

	}

	fmt.Println("lenght is:", length)
	defer file.Close() // we should always use defer in this because if anything left then it will be executed at last

}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)
	fmt.Println("Text inside files is:", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
