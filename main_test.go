package main

import (
    "fmt"
    "strconv"
    "strings"
    "testing"

    "plhtml/parser"
    "plhtml/scanner"
    "plhtml/token"
    "plhtml/util"
)

var tests = [...]string{
    "factorial",
    "fibonacci",
    "hello",
    "leap",
    "prime",
    "scopes",
}

func TestScanner(t *testing.T) {
    for _, test := range tests {
        tokens := scan("./tests/" + test + ".html")
        compare(t, "scanner/"+test, tokensToString(tokens))
    }
}

func TestParser(t *testing.T) {
    for _, test := range tests {
        tokens := scan("./tests/" + test + ".html")
        myParser := parser.New()
        prgNode := myParser.Parse(tokens)
        compare(t, "parser/"+test, prgNode.ToString())
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

func tokensToString(tokens []scanner.Token) string {

    result := ""
    for _, tok := range tokens {

        if tok.Type == token.EOF {
            break
        }

        var value string
        switch tok.Type {
        case token.Identifier, token.StringConst:
            value = util.Unescape(tok.StrVal)
        case token.IntConst:
            value = strconv.Itoa(tok.IntVal)
        case token.RealConst:
            value = util.FloatToString(tok.RealVal)
        case token.BoolConst:
            value = strconv.FormatBool(tok.BoolVal)
        }

        if value != "" {
            value = "|" + value + "|"
        }

        result += tok.Type.String() + value + "\n"

    }

    return result
}
