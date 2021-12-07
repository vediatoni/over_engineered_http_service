package main

import (
	"fmt"
	"os"
)

const RandomText = "Hello World 11"

func main() {
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	s := new(port)

	fmt.Printf("Server is running on port %v\n", port)
	s.run()
}
