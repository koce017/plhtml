package semantic

import (
    "fmt"
    "go/constant"
    "plhtml/ast"
    "plhtml/logger"
    "plhtml/scope"
    "plhtml/token"
)

var myLogger = logger.New("ANALYZER")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Analyzer struct {
    currentScope *scope.Scope
}

func NewAnalyzer() *Analyzer {
    analyzer := new(Analyzer)
    analyzer.currentScope = scope.New(1, scope.New(0, nil))
    return analyzer
}

func (analyzer *Analyzer) VisitBinaryOpExpr(node *ast.BinaryOpExprNode) constant.Kind {
    leftType := node.LeftExpr.AcceptAnalyzer(analyzer)
    rightType := node.RightExpr.AcceptAnalyzer(analyzer)

    switch node.Operator {
    case token.AndOp, token.OrOp:
        if leftType == constant.Bool && rightType == constant.Bool {
            return constant.Bool
        }
    case token.EqOp, token.NeqOp:
        if leftType == rightType {
            return constant.Bool
        }
    case token.LtOp, token.GtOp, token.LeqOp, token.GeqOp:
        if isNumType(leftType) && isNumType(rightType) {
            return constant.Bool
        }
    case token.Plus:
        if leftType == constant.String || rightType == constant.String {
            return constant.String
        }
        fallthrough
    case token.Minus, token.Multiply, token.Slash, token.Modulo:
        if leftType == constant.Int && rightType == constant.Int {
            return constant.Int
        }
        if isNumType(leftType) && isNumType(rightType) {
            return constant.Float
        }
    }

    panic(fmt.Sprintf("Error on line %d: operator %s is not supported for operands of given types", node.GetLine(), node.Operator.String()))
}

func (analyzer *Analyzer) VisitBoolConst(node *ast.BoolConstNode) constant.Kind {
    return constant.Bool
}

func (analyzer *Analyzer) VisitControlFlowStmt(node *ast.ControlFlowStmtNode) {

    condType := node.Condition.AcceptAnalyzer(analyzer)
    if condType != constant.Bool {
        panic(fmt.Sprintf("Error on line %d: non-bool used as condition", node.GetLine()))
    }

    analyzer.currentScope = scope.New(analyzer.currentScope.Id+1, analyzer.currentScope)

    for i := range node.Statements {
        node.Statements[i].AcceptAnalyzer(analyzer)
    }

    analyzer.currentScope = analyzer.currentScope.Parent

}

func (analyzer *Analyzer) VisitIdentifier(node *ast.IdentifierNode) constant.Kind {
    node.Scope = analyzer.currentScope
    sym, varId, ok := analyzer.currentScope.Lookup(node.Name)
    if !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Name))
    }

    node.Name = varId
    kind := kindOfPrimitiveType(sym.Type)

    if kind == constant.Unknown {
        panic(fmt.Sprintf("Error on line %d: only primitive types are supported", node.GetLine()))
    }

    return kind
}

func (analyzer *Analyzer) VisitIntConst(node *ast.IntConstNode) constant.Kind {
    return constant.Int
}

func (analyzer *Analyzer) VisitMainFunc(node *ast.MainFuncNode) {
    analyzer.currentScope = scope.New(analyzer.currentScope.Id+1, analyzer.currentScope)
    for i := range node.Statements {
        node.Statements[i].AcceptAnalyzer(analyzer)
    }
    analyzer.currentScope = analyzer.currentScope.Parent
}

func (analyzer *Analyzer) VisitProgram(node *ast.ProgramNode) {
    node.Body.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitProgramBody(node *ast.ProgramBodyNode) {
    node.MainFunc.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitReadStmt(node *ast.ReadStmtNode) {
    node.Scope = analyzer.currentScope
    _, varId, ok := analyzer.currentScope.Lookup(node.Identifier.Name)
    if !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }
    node.Identifier.Name = varId
}

func (analyzer *Analyzer) VisitRealConst(node *ast.RealConstNode) constant.Kind {
    return constant.Float
}

func (analyzer *Analyzer) VisitStringConst(node *ast.StringConstNode) constant.Kind {
    return constant.String
}

func (analyzer *Analyzer) VisitUnaryExpr(node *ast.UnaryExprNode) constant.Kind {
    exprType := node.Expr.AcceptAnalyzer(analyzer)

    switch node.Operator {
    case token.Minus:
        if isNumType(exprType) {
            return exprType
        }
    case token.Excl:
        if exprType == constant.Bool {
            return exprType
        }
    }

    panic(fmt.Sprintf("Error on line %d: operator %s is not supported for expression of given type", node.GetLine(), node.Operator.String()))
}

func (analyzer *Analyzer) VisitVarAssign(node *ast.VarAssignNode) {
    node.Scope = analyzer.currentScope
    sym, varId, ok := analyzer.currentScope.Lookup(node.Identifier.Name)
    if !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }
    node.Identifier.Name = varId

    symType := kindOfPrimitiveType(sym.Type)

    if symType == constant.Unknown {
        panic(fmt.Sprintf("Error on line %d: only primitive types are supported", node.GetLine()))
    }

    valType := node.Value.AcceptAnalyzer(analyzer)
    ok = (symType == valType) || (isNumType(symType) && isNumType(valType))

    if !ok {
        panic(fmt.Sprintf("Error on line %d: incompatible types for assigment", node.GetLine()))
    }
}

func (analyzer *Analyzer) VisitVarDecl(node *ast.VarDeclNode) {
    node.Scope = analyzer.currentScope
    if _, _, ok := analyzer.currentScope.Lookup(node.Type.Name); !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }
    if analyzer.currentScope.DeclaredLocally(node.Identifier.Name) {
        panic(fmt.Sprintf("Error on line: %d: variable %s is already declared in this scope", node.GetLine(), node.Identifier.Name))
    }
    analyzer.currentScope.Insert(&scope.Symbol{Name: node.Identifier.Name, Type: node.Type.Name, Line: node.GetLine()})
}

func (analyzer *Analyzer) VisitWriteStmt(node *ast.WriteStmtNode) {
    valueType := node.Value.AcceptAnalyzer(analyzer)
    if valueType != constant.String {
        panic(fmt.Sprintf("Error on line %d: cannot output non-string", node.GetLine()))
    }
}

func isNumType(kind constant.Kind) bool {
    return kind == constant.Int || kind == constant.Float
}

func kindOfPrimitiveType(typeName string) constant.Kind {
    switch typeName {
    case scope.TypeInteger:
        return constant.Int
    case scope.TypeReal:
        return constant.Float
    case scope.TypeBoolean:
        return constant.Bool
    case scope.TypeString:
        return constant.String
    }
    return constant.Unknown
}
