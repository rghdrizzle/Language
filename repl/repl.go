package repl

import(
	"bufio"
	"fmt"
	"io"
	"rghdrizzle/language/lexer"
	//"rghdrizzle/language/tokens"
	"rghdrizzle/language/parser"
	"rghdrizzle/language/evaluator"
	"rghdrizzle/language/objects"
)

const promt ="%>>"

func StartRepl(in io.Reader,out io.Writer){
	scanner := bufio.NewScanner(in)
	env  := objects.NewEnvironment()

	for {
		fmt.Print(promt)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line:= scanner.Text()
		l:= lexer.New(line)
		p:= parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors())!=0{
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program,env)
		if evaluated!=nil{
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out,"\n")
		}
	}

}

func printParserErrors(out io.Writer,errors []string){
	for _, msg:= range errors{
		io.WriteString(out,"\t"+msg+"\n")
	}
}