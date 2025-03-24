package parser

import (
    "plhtml/ast"
    "plhtml/token"
)

func (parser *Parser) parseStatements() []ast.StatementNode {
    myLogger.Debug("=BEG= Statements")

    var stmts []ast.StatementNode
    for stmt := parser.parseStatement(); stmt != nil; stmt = parser.parseStatement() {
        stmts = append(stmts, stmt)
    }

    myLogger.Debug("=END= Statements")
    return stmts
}

func (parser *Parser) parseStatement() ast.StatementNode {
    parser.eat(token.LessThan)

    switch parser.peek().Type {
    case token.Var:
        node := parser.parseVarDecl()
        return &node
    case token.Data:
        node := parser.parseVarAssign()
        return &node
    case token.Output:
        node := parser.parseWriteStmt()
        return &node
    case token.Input:
        node := parser.parseReadStmt()
        return &node
    case token.Div:
        node := parser.parseControlFlowStmt()
        return &node
    }

    myLogger.Debug("'%s' not a statement", string(parser.peek().Type))
    parser.goBack() // '<'
    return nil
}

func (parser *Parser) parseVarDecl() ast.VarDeclNode {
    line := parser.current().Line
    myLogger.Debug("=BEG= Var Declaration %d", line)
    parser.eat(token.Var)
    parser.eat(token.Class)
    parser.eat(token.Equal)
    parser.eat(token.DQuote)
    varType := parser.parseIdentifier()
    parser.eat(token.DQuote)
    parser.eat(token.GreaterThan)
    varName := parser.parseIdentifier()
    parser.parseCloseTag(token.Var)
    myLogger.Debug("=END= Var Declaration")
    return ast.VarDeclNode{Line: line, Type: varType, Identifier: varName}
}

func (parser *Parser) parseVarAssign() ast.VarAssignNode {
    line := parser.current().Line
    myLogger.Debug("=BEG= Var Assignment %d", line)
    parser.eat(token.Data)
    parser.eat(token.Value)
    parser.eat(token.Equal)
    parser.eat(token.DQuote)
    value := parser.parseExpr()
    parser.eat(token.DQuote)
    parser.eat(token.GreaterThan)
    identifier := parser.parseIdentifier()
    parser.parseCloseTag(token.Data)
    myLogger.Debug("=END= Var Assignment")
    return ast.VarAssignNode{Line: line, Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() ast.WriteStmtNode {
    line := parser.current().Line
    myLogger.Debug("=BEG= Write %d", line)
    parser.eat(token.Output)
    parser.eat(token.GreaterThan)
    value := parser.parseExpr()
    parser.parseCloseTag(token.Output)
    myLogger.Debug("=END= Write")
    return ast.WriteStmtNode{Line: line, Value: value}
}

func (parser *Parser) parseReadStmt() ast.ReadStmtNode {
    line := parser.current().Line
    myLogger.Debug("=BEG= Read %d", line)
    parser.eat(token.Input)
    parser.eat(token.Name)
    parser.eat(token.Equal)
    parser.eat(token.DQuote)
    identifier := parser.parseIdentifier()
    parser.eat(token.DQuote)
    parser.eat(token.GreaterThan)
    myLogger.Debug("=END= Read")
    return ast.ReadStmtNode{Line: line, Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() ast.ControlFlowStmtNode {
    line := parser.current().Line
    myLogger.Debug("=BEG= Control Flow %d", line)
    parser.eat(token.Div)
    parser.eat(token.Data)
    parser.eat(token.Minus)
    stmtType := parser.eat(token.If, token.While)
    parser.eat(token.Equal)
    parser.eat(token.DQuote)
    condition := parser.parseExpr()
    parser.eat(token.DQuote)
    parser.eat(token.GreaterThan)
    stmts := parser.parseStatements()
    parser.parseCloseTag(token.Div)
    myLogger.Debug("=END= Control Flow")
    return ast.ControlFlowStmtNode{Line: line, Type: stmtType, Condition: condition, Statements: stmts}
}
