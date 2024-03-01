package parser

import (
    "plhtml/ast"
    "plhtml/token"
)

func (parser Parser) parseProgram() ast.ProgramNode {
    myLogger.Debug("=BEG= Program")
    parser.parseDoctype()
    prg := parser.parseHTML()
    myLogger.Debug("=END= Program")
    return prg
}

func (parser *Parser) parseDoctype() {
    myLogger.Debug("=BEG= Doctype")
    parser.eat(token.LessThan)
    parser.eat(token.Excl)
    parser.eat(token.Doctype)
    parser.eat(token.HTML)
    parser.eat(token.GreaterThan)
    myLogger.Debug("=END= Doctype")
}

func (parser *Parser) parseHTML() ast.ProgramNode {
    myLogger.Debug("=BEG= HTML")
    parser.eat(token.LessThan)
    parser.eat(token.HTML)
    parser.eat(token.Lang)
    parser.eat(token.Equal)
    parser.eat(token.DQuote)
    parser.parseIdentifier()
    parser.eat(token.DQuote)
    parser.eat(token.GreaterThan)
    programTitle := parser.parseProgramHeader()
    programBody := parser.parseProgramBody()
    parser.parseCloseTag(token.HTML)
    myLogger.Debug("=END= HTML")
    return ast.ProgramNode{Title: programTitle, Body: programBody}
}

func (parser *Parser) parseProgramHeader() ast.StringConstNode {
    myLogger.Debug("=BEG= Prg Header")
    parser.parseOpenTag(token.Head)
    programTitle := parser.parseProgramTitle()
    parser.parseCloseTag(token.Head)
    myLogger.Debug("=END= Prg Header")
    return programTitle
}

func (parser *Parser) parseProgramTitle() ast.StringConstNode {
    myLogger.Debug("=BEG= Prg Title")
    parser.parseOpenTag(token.Title)
    programTitle := parser.parseStringConst()
    parser.parseCloseTag(token.Title)
    myLogger.Debug("=END= Prg Title")
    return programTitle
}

func (parser *Parser) parseProgramBody() ast.ProgramBodyNode {
    myLogger.Debug("=BEG= Prg Body")
    parser.parseOpenTag(token.Body)
    mainFunc := parser.parseMainFunc()
    parser.parseCloseTag(token.Body)
    myLogger.Debug("=END= Prg Title")
    return ast.ProgramBodyNode{MainFunc: mainFunc}
}

func (parser *Parser) parseMainFunc() ast.MainFuncNode {
    myLogger.Debug("=BEG= Main")
    parser.parseOpenTag(token.Main)
    statements := parser.parseStatements()
    parser.parseCloseTag(token.Main)
    myLogger.Debug("=END= Main")
    return ast.MainFuncNode{Statements: statements}
}
