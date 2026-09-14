package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Sany programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	// REPL will be added here later
}
