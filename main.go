package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/themarkfullton/go-brainer-interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! I am the Go-Brainer interpreter!\n\n", user.Username)

	fmt.Printf("Please enter some code for me to interpret!\n")

	repl.Start(os.Stdin, os.Stdout)
}
