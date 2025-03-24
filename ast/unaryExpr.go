package ast

import (
    "go/constant"
    "plhtml/token"
)

type UnaryExprNode struct {
    Line     int
    Operator token.Type
    Expr     ExpressionNode
}

func (node UnaryExprNode) GetLine() int {
    return node.Line
}

func (node UnaryExprNode) ToString() string {
    return node.Operator.String() + node.Expr.ToString()
}

func (node *UnaryExprNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitUnaryExpr(node)
}

func (node *UnaryExprNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitUnaryExpr(node)
}
