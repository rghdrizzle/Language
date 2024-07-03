package main

import(
	"rghdrizzle/language/repl"
	"fmt"
	"os"
	"os/user"
)

func main(){
	user , err:= user.Current()
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Welcome to the Drizzle Language %s\n",user.Username)
	repl.StartRepl(os.Stdin,os.Stdout)
}