package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Neeraj-Natu/shifu/evaluator"
	"github.com/Neeraj-Natu/shifu/lexer"
	"github.com/Neeraj-Natu/shifu/object"
	"github.com/Neeraj-Natu/shifu/parser"
)

/*
read from the input source until encountering a newline,
take the just read line and pass it to an instance of our
lexer and that to our Parser once the parser is done with
it's job. we pass the AST into the evaluator which evaluates
the whole program represented by the AST. After all this
we print out the parsing errors if any or the evaluation
result calling Inspect() method on Program that recursively
calls the Inspect() method on all of the statements belonging
to that program.
*/

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		//fmt.Printf("Lexer output is :  \n")
		//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		//	fmt.Printf("%+v\n", tok)
		//}
		//fmt.Printf("----------------------------------------------------------- \n")
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		//io.WriteString(out, "--------- Parser Output ---------------------------")
		//io.WriteString(out, program.String())
		//io.WriteString(out, "\n")
		//io.WriteString(out, "------------------------------------")
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ACCIDENTS)
	io.WriteString(out, "Learning code is an art that takes years to master. Do not be disappointed if you have failed !! \n")
	io.WriteString(out, "parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const ACCIDENTS = `	
 _____________________________
|                             |
|  There are no accidents !!  |                      
|_____________________________|

`
