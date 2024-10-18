package main

import (
	"fmt"
	"monkeylang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s and welcome to the MonkeyLang REPL!\n", user.Username)
	fmt.Printf("Feel free to start typeing in commands\n\n")
	repl.Init(os.Stdin, os.Stdout)
}
