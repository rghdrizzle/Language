package repl

import(
	"bufio"
	"fmt"
	"io"
	"rghdrizzle/language/lexer"
	"rghdrizzle/language/tokens"
)

const promt ="%>>"

func StartRepl(in io.Reader,out io.Writer){
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(promt)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line:= scanner.Text()
		l:= lexer.New(line)

		for tok:=l.NextToken();tok.Type!=token.EOF;tok = l.NextToken(){
			fmt.Printf("%+v\n",tok)
		}
	}

}