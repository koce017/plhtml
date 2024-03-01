package ast

type MainFuncNode struct {
    Line       int
    Statements []StatementNode
}

func (node MainFuncNode) GetLine() int {
    return node.Line
}

func (node MainFuncNode) ToString() string {
    str := "\nMain:"
    for _, stmt := range node.Statements {
        str += "\n" + stmt.ToString(1)
    }
    return str
}

func (node *MainFuncNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitMainFunc(node)
}

func (node *MainFuncNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitMainFunc(node)
}
