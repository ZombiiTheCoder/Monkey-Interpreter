package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

func Start(in io.Reader, out io.Writer, user string) {

	var PROMPT = user + " >> "

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
		
		// l2 := lexer.New(line)
		// for tok := l2.NextToken(); tok.Type != token.EOF; tok = l2.NextToken() {
		// 	fmt.Printf("%+v\n", tok)
		// }
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}