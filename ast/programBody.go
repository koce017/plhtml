package ast

type ProgramBodyNode struct {
    Line     int
    MainFunc MainFuncNode
}

func (node ProgramBodyNode) GetLine() int {
    return node.Line
}

func (node ProgramBodyNode) ToString() string {
    return node.MainFunc.ToString()
}

func (node *ProgramBodyNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitProgramBody(node)
}

func (node *ProgramBodyNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitProgramBody(node)
}
