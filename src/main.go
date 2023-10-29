package main

import (
	"fmt"
	"os"
	"source/repl"
)

func main() {
	fmt.Println("Welcome! Feel free to type commands : ")
	repl.Start(os.Stdin, os.Stdout)
}