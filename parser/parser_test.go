package parser


import(
  "testing"
  "rghdrizzle/language/ast"
  "rghdrizzle/language/lexer"
)

func TestLetStatement(t *testing.T){
  input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
l := lexer.New(input)
p := New(l)

program := p.ParseProgram()
if program == nil{
  t.fatalf("ParseProgram method returned nil")
}
if len(program.Statements)!=3{
  t.Fatalf("program.Statements does not contain 3 statements. got=%d",
len(program.Statements))
}

tests:= []struct{}
}
