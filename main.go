package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/komuw/cali/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unajua niaje %s! \n\tThis is the cali programming language!\n", user.Username)
	fmt.Printf("You can type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
