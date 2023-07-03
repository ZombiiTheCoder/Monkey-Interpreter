package main

import (
	"monkey/repl"
	"os"
	"os/user"
	"strings"
)

// Page 80

func main() {

	user, _ := user.Current()
	u := strings.Split(user.Username, "\\")
	
	repl.Start(os.Stdin, os.Stdout, u[len(u)-1])

}