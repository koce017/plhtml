package main

import (
	"fmt"
	"os"
	"plhtml/interpreter"
	"plhtml/parser"
	"plhtml/scanner"
	"plhtml/semantic"
	"plhtml/util"
)

func main() {
    in := os.Stdin
    var err error

	source := util.ReadFile(os.Args[1])
	convertToJs := len(os.Args) > 2 && os.Args[2] == "--js"

	myScanner := scanner.New()
    tokens := myScanner.Scan(source)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)

	if convertToJs {
		fmt.Print(prgNode.ToString())
		return
	}

    analyzer := semantic.NewAnalyzer()
    prgNode.AcceptAnalyzer(analyzer)

    interp := interpreter.New(in)
    prgNode.AcceptInterpreter(interp)

    err = in.Close()
    util.Check(err)
}
