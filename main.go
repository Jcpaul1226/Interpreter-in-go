package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming Language! \n", user.Username)
	fmt.Printf("Type in some commands and I will Tokenize them! \n")
	repl.Start(os.Stdin, os.Stdout)
}
