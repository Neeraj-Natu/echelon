package main

import (
	"fmt"
	"github.com/Neeraj-Natu/shifu/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Shifu programming language!\n",
		user.Username)
	fmt.Printf("Attain inner peace by typing in the commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
