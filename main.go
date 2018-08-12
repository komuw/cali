package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/komuw/khaled/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bless up %s! \n\tMajor Key alert. \n\tThis is the khaled programming language!\n", user.Username)
	fmt.Printf("You can type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
