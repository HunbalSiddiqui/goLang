package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hunbalsiddiqui/mongoapi/router"
)

func main() {
	fmt.Println("MongoAPIs")
	r := router.Router()
	fmt.Println("Server is started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000...")
}
