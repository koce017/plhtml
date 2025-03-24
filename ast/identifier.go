package ast

import (
    "go/constant"
    "plhtml/scope"
)

type IdentifierNode struct {
    Line  int
    Name  string
    Scope *scope.Scope
}

func (node IdentifierNode) GetLine() int {
    return node.Line
}

func (node IdentifierNode) ToString() string {
    return node.Name
}

func (node *IdentifierNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitIdentifier(node)
}

func (node *IdentifierNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitIdentifier(node)
}
