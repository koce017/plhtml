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
    return ident(lvl, "Read "+node.Identifier.ToString())
}

func (node *ReadStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitReadStmt(node)
}

func (node *ReadStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitReadStmt(node)
}
