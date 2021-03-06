package main

import (
	"fmt"
	"os"
)

const RandomText = "Ola! Hablas espanol?"

func main() {

	s := new(getPort())

	fmt.Printf("Server is running on port %v\n", getPort())
	s.run()
}

func getPort() string {
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	return port
}
