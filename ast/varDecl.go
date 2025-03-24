package ast

import "plhtml/scope"

type VarDeclNode struct {
    Line       int
    Type       IdentifierNode
    Identifier IdentifierNode
    Scope      *scope.Scope
}

func (node VarDeclNode) GetLine() int {
    return node.Line
}

func (node VarDeclNode) ToString(lvl int) string {
    return ident(lvl, node.Type.ToString()+" "+node.Identifier.ToString())
}

func (node *VarDeclNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitVarDecl(node)
}

func (node *VarDeclNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitVarDecl(node)
}
