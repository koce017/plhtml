package ast

import (
    "fmt"
    "go/constant"
)

type RealConstNode struct {
    Line  int
    Value float64
}

func (node RealConstNode) GetLine() int {
    return node.Line
}

func (node RealConstNode) ToString() string {
    return fmt.Sprintf("%f", node.Value)
}

func (node *RealConstNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitRealConst(node)
}

func (node *RealConstNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitRealConst(node)
}
