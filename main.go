package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/andreeamnl/GPL/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! slay\n",
		user.Username)
	fmt.Printf("Let's start\n")
	repl.Start(os.Stdin, os.Stdout)
}
