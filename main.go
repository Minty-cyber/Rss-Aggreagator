package main

import (
	"fmt"
	"os"
	"log"

)

func main() {
	fmt.Println("Hello world")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not available")
	}

	fmt.Println("Port:", portString)
}