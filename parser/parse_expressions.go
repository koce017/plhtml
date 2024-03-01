package parser

import (
    "fmt"

    "plhtml/ast"
    "plhtml/token"
)

func (parser *Parser) parseExpr() ast.ExpressionNode {
    return parser.parseExprL(1)
}

func (parser *Parser) parseExprL(l int) ast.ExpressionNode {
    var operators []TokenType

    switch l {
    case 1:
        operators = []TokenType{token.AndOp, token.OrOp}
    case 2:
        operators = []TokenType{token.EqOp, token.NeqOp}
    case 3:
        operators = []TokenType{token.LtOp, token.GtOp, token.LeqOp, token.GeqOp}
    case 4:
        operators = []TokenType{token.Plus, token.Minus}
    case 5:
        operators = []TokenType{token.Multiply, token.Slash, token.Modulo}
    case 6:
        return parser.parseFactor()
    }

    expr := parser.parseExprL(l + 1)
    for parser.eatOpt(operators...) {
        expr = &ast.BinaryOpExprNode{Line: expr.GetLine(), LeftExpr: expr, Operator: parser.next().Type, RightExpr: parser.parseExprL(l + 1)}
    }
    return expr
}

func (parser *Parser) parseFactor() ast.ExpressionNode {
    nextToken := parser.peek()
    switch nextToken.Type {
    case token.IntConst:
        node := parser.parseIntConst()
        return &node
    case token.RealConst:
        node := parser.parseRealConst()
        return &node
    case token.BoolConst:
        node :=  parser.parseBoolConst()
        return &node
    case token.StringConst:
        node :=  parser.parseStringConst()
        return &node
    case token.Identifier:
        node :=  parser.parseIdentifier()
        return &node
    case token.LParen:
        return parser.parseParenExpr()
    case token.Plus, token.Minus, token.Excl:
        node :=  parser.parseUnaryExpr()
        return &node
    default:
        panic(fmt.Sprintf("Invalid factor '%s' at %d:%d.", string(nextToken.Type), nextToken.Line, nextToken.Column))
    }
}

func (parser *Parser) parseIdentifier() ast.IdentifierNode {
    myLogger.Debug("=BEG= Identifier")
    parser.eat(token.Identifier)
    myLogger.Debug("=END= Identifier")
    return ast.IdentifierNode{Line: parser.current().Line, Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() ast.IntConstNode {
    myLogger.Debug("=BEG= Integer Constant")
    parser.eat(token.IntConst)
    myLogger.Debug("=END= Integer Constant")
    return ast.IntConstNode{Line: parser.current().Line, Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() ast.RealConstNode {
    myLogger.Debug("=BEG= Real Constant")
    parser.eat(token.RealConst)
    myLogger.Debug("=END= Real Constant")
    return ast.RealConstNode{Line: parser.current().Line, Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() ast.BoolConstNode {
    myLogger.Debug("=BEG= Boolean Constant")
    parser.eat(token.BoolConst)
    myLogger.Debug("=END= Boolean Constant")
    return ast.BoolConstNode{Line: parser.current().Line, Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() ast.StringConstNode {
    myLogger.Debug("=BEG= String Constant")
    parser.eat(token.StringConst)
    myLogger.Debug("=END= String Constant")
    return ast.StringConstNode{Line: parser.current().Line, Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ast.ExpressionNode {
    myLogger.Debug("=BEG= ()")
    parser.eat(token.LParen)
    expr := parser.parseExpr()
    parser.eat(token.RParen)
    myLogger.Debug("=END= ()")
    return expr
}

func (parser *Parser) parseUnaryExpr() ast.UnaryExprNode {
    myLogger.Debug("=BEG= Unary")
    op := parser.eat(token.Plus, token.Minus, token.Excl)
    expr := parser.parseFactor()
    myLogger.Debug("=END= Unary")
    return ast.UnaryExprNode{Line: expr.GetLine(), Operator: op, Expr: expr}
}
