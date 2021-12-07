package main

import (
	"fmt"
)


const port = ":8080"
const RandomText = "Hello World 11"


func main() {
	s := new(port)

	fmt.Printf("Server is running on port %v\n", port)
	s.run()
}
