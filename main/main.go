package main

import (
	"Nutlang/repl"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	username := user.Username
	// Capitalize the first character
	if len(username) > 0 {
		username = strings.ToUpper(username[0:1]) + username[1:]
	}
	fmt.Printf("Hello %s! This is the Nut programming language!\n",
		username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
