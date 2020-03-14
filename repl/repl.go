package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/Neeraj-Natu/shifu/lexer"
	"github.com/Neeraj-Natu/shifu/token"
)

/*
read from the input source until encountering a newline, 
take the just read line and pass it to an instance of our 
lexer and finally print all the tokens the lexer gives us 
until we encounter EOF.
*/

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
