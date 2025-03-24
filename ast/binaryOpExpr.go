package ast

import (
    "go/constant"
    "plhtml/token"
)

type BinaryOpExprNode struct {
    Line      int
    LeftExpr  ExpressionNode
    Operator  token.Type
    RightExpr ExpressionNode
}

func (node BinaryOpExprNode) GetLine() int {
    return node.Line
}

func (node BinaryOpExprNode) ToString() string {
    return node.LeftExpr.ToString() + " " + node.Operator.String() + " " + node.RightExpr.ToString()
}

func (node *BinaryOpExprNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitBinaryOpExpr(node)
}

func (node *BinaryOpExprNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitBinaryOpExpr(node)
}
