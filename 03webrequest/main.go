package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://hunbalsiddiqui.com"

func main() {
	fmt.Println("MHS web request")

	response, err := http.Get(url)
	if err != nil {
		panic((err))
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic((err))
	}

	fmt.Printf("String is: ", string(databytes))
}
