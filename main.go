package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of strings that holds all the command-line arguments.
	// The first element (os.Args[0]) is always the name of the program itself.
	fmt.Println("The program name is:", os.Args[0])

	// The rest of the arguments are in os.Args[1:]
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Hello, world! Try providing some arguments.")
	} else {
		fmt.Println("Hello,", args)
	}
}
