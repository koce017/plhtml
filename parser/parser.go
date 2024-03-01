package parser

import (
    "fmt"

    "plhtml/ast"
    "plhtml/logger"
    "plhtml/scanner"
    "plhtml/token"
)

type Token = scanner.Token
type TokenType = token.Type

var myLogger = logger.New("PARSER")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Parser struct {
    index  int
    tokens []Token
}

func New() *Parser {
    parser := new(Parser)
    return parser
}

func (parser *Parser) Parse(tokens []Token) ast.ProgramNode {
    parser.index = -1
    parser.tokens = tokens
    return parser.parseProgram()
}

func (parser Parser) current() Token {
    return parser.tokens[parser.index]
}

func (parser *Parser) next() Token {
    if parser.index < 0 || parser.current().Type != token.EOF {
        parser.index++
    }
    return parser.current()
}

func (parser *Parser) goBack() {
    parser.index--
}

func (parser Parser) peek() Token {
    next := parser.next()
    parser.goBack()
    return next
}

func (parser *Parser) eatOpt(expected ...TokenType) bool {
    actual := parser.peek().Type
    myLogger.Debug("'%s'", string(actual))

    for _, exp := range expected {
        if actual == exp {
            return true
        }
    }

    return false
}

func (parser *Parser) eat(expected ...TokenType) TokenType {
    ok := parser.eatOpt(expected...)
    if !ok {
        nextToken := parser.peek()
        panic(fmt.Sprintf("Unexpected token '%s' at %d:%d.", nextToken.Type.String(), nextToken.Line, nextToken.Column))
    }
    return parser.next().Type
}

func (parser *Parser) parseOpenTag(expected TokenType) {
    myLogger.Debug("<%s> expected", string(expected))
    parser.eat(token.LessThan)
    parser.eat(expected)
    parser.eat(token.GreaterThan)
}

func (parser *Parser) parseCloseTag(expected TokenType) {
    myLogger.Debug("</%s> expected", string(expected))
    parser.eat(token.LessThan)
    parser.eat(token.Slash)
    parser.eat(expected)
    parser.eat(token.GreaterThan)
}
