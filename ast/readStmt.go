package ast

import "plhtml/scope"

type ReadStmtNode struct {
    Line       int
    Identifier IdentifierNode
    Scope      *scope.Scope
}

func (node ReadStmtNode) GetLine() int {
    return node.Line
}

func (node ReadStmtNode) ToString(lvl int) string {
	return ident(lvl, node.Identifier.ToString()+" = prompt(\"Enter the value for "+node.Identifier.ToString()+":\")")
}

func (node *ReadStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitReadStmt(node)
}

func (node *ReadStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitReadStmt(node)
}
