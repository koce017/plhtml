package main

import (
	"fmt"
	"strings"
	"testing"

	"plhtml/parser"
	"plhtml/scanner"
	"plhtml/util"
)

var tests = [...]string{
    "factorial",
    "fibonacci",
    "hello",
    "leap",
    "prime",
	"reals",
    "scopes",
}

func TestParser(t *testing.T) {
    for _, test := range tests {
		if util.FileExists("./tests/parser/" + test + ".expected.txt") {
			tokens := scan("./tests/" + test + ".html")
			myParser := parser.New()
			prgNode := myParser.Parse(tokens)
			compare(t, "parser/"+test, prgNode.ToString())
		}
    }
}

func scan(file string) []scanner.Token {
    source := util.ReadFile(file)
    myScanner := scanner.New()
    return myScanner.Scan(source)
}

func compare(t *testing.T, testPath string, actual string) {
    expected := util.ReadFile("./tests/" + testPath + ".expected.txt")

    if strings.TrimSpace(expected) != strings.TrimSpace(actual) {
        util.WriteFile("./tests/"+testPath+".actual.txt", actual)
        fmt.Println("FAIL: " + testPath)
        t.Fail()
    } else {
        fmt.Println("PASS: " + testPath)
    }
}
