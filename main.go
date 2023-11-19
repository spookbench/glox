package main

import (
	"fmt"
	"os"

	"github.com/spookbench/glox/lexer"
)

var args []string

func main() {
	switch i := len(args); i {
	case 0:
		runPrompt()
	case 1:
		runFile(args[0])
	default:
		fmt.Println("Usage glox [script]")
		os.Exit(64)

	}
}

func runFile(path string) {
	file, err := os.OpenFile(path, 1, 1)
	if err != nil {
		panic(fmt.Sprintf("failed to open file %q: %v", path, err))
	}
	defer file.Close()
	var f []byte
	_, err = file.Read(f)
	if err != nil {
		panic(fmt.Sprintf("failed to read file %q: %v", path, err))
	}

	run(string(f))

	if HadError {
		os.Exit(65)
	}
}

func run(s string) {
	l := lexer.New(s)
	tokens := l.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

var HadError = false

func parserError(line int, msg string) {
	fmt.Printf("[line %d] Error %s: %s /n", line, "", msg)
	HadError = true
}

func runPrompt() {
	var line string
	fmt.Println("glox Lox interpreter v0.01")

	for {
		fmt.Print("> ")
		fmt.Scan(&line)
		if line == "" {
			continue
		}
		run(line)
		HadError = false
	}

}
