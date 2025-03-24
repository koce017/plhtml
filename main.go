package main

import (
    "os"
    "plhtml/interpreter"
    "plhtml/parser"
    "plhtml/scanner"
    "plhtml/semantic"
    "plhtml/util"
)

func main() {
    const argSource = 1
    const argInput = 2

    in := os.Stdin
    var err error
    var source string

    if len(os.Args) > argSource {
        source = util.ReadFile(os.Args[argSource])
    } else {
        panic("Usage: plhtml <source_file> [<input_file>]")
    }

    if len(os.Args) > argInput {
        in, err = os.Open(os.Args[argInput])
        util.Check(err)
    }

    myScanner := scanner.New()
    tokens := myScanner.Scan(source)

    myParser := parser.New()
    prgNode := myParser.Parse(tokens)

    analyzer := semantic.NewAnalyzer()
    prgNode.AcceptAnalyzer(analyzer)

    interp := interpreter.New(in)
    prgNode.AcceptInterpreter(interp)

    err = in.Close()
    util.Check(err)
}
