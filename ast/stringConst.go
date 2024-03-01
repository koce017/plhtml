package ast

import (
    "go/constant"
    "plhtml/util"
)

type StringConstNode struct {
    Line  int
    Value string
}

func (node StringConstNode) GetLine() int {
    return node.Line
}

func (node StringConstNode) ToString() string {
    return "\"" + util.Unescape(node.Value) + "\""
}

func (node *StringConstNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitStringConst(node)
}

func (node *StringConstNode) AcceptInterpreter(v IInterpreter) constant.Value {
    return v.VisitStringConst(node)
}
