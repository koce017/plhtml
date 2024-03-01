package ast

import "plhtml/scope"

type VarAssignNode struct {
    Line       int
    Identifier IdentifierNode
    Value      ExpressionNode
    Scope      *scope.Scope
}

func (node VarAssignNode) GetLine() int {
    return node.Line
}

func (node VarAssignNode) ToString(lvl int) string {
    return ident(lvl, node.Identifier.ToString()+" = "+node.Value.ToString())
}

func (node *VarAssignNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitVarAssign(node)
}

func (node *VarAssignNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitVarAssign(node)
}
