package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/Neeraj-Natu/shifu/repl"
)

func main() {

	outputPtr := flag.String("output", "lang", "output from lexer, parser or langauge itself")
	flag.Parse()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(SHIFU)
	fmt.Printf("Hello %s! \n",
		user.Username)
	fmt.Printf("Attain inner peace by exploring the language!! \n")

	if *outputPtr == "lexer" {
		repl.StartLexer(os.Stdin, os.Stdout)
	}
	if *outputPtr == "parser" {
		repl.StartParser(os.Stdin, os.Stdout)
	}
	if *outputPtr == "lang" {
		repl.StartLang(os.Stdin, os.Stdout)
	}

}

const SHIFU = `
 _ _ _ _ _ _ _ _ __ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
|      ______    __      __    __________    __________    ___      ___      |
|    /   ____|  |  |    |  |  |___    ___|  |   _______|  |   |    |   |     |
|   |   |       |  |    |  |      |  |      |   |         |   |    |   |     |
|   |   |___    |  |____|  |      |  |      |   |____     |   |    |   |     |
|    \____  \   |   ____   |      |  |      |    ____|    |   |    |   |     |
|         |  |  |  |    |  |      |  |      |   |         |   |    |   |     |
|    _____|  |  |  |    |  |   ___|  |___   |   |         |   |____|   |     |
|   |______ /   |__|    |__|  |__________|  |___|          \__________/      |
|                                                                            |
|   Programming Language                                                     |
|_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ __|

`
