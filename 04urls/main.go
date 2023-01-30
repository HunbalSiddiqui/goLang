package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dfgsdfsd"

func main() {
	fmt.Println("Welcome to URLs")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}
