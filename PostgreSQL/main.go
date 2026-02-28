package main

import (
	"log"
	"fmt"
	"net/http"
	"postgresql/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}