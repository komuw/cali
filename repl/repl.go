package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/komuw/khaled/lexer"
	"github.com/komuw/khaled/token"
)

/*
A REPL(“Read Eval Print Loop) reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again.

We don’t know how to fully “Eval” khaled source code yet. We only have one part of the
process that hides behind “Eval”: we can tokenize khaled source code.
*/

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	/*
		This tokenizes khaled source code and prints the tokens.
		Later on, we will expand on this and add parsing and evaluation to it.
	*/
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
