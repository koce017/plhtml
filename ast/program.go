package ast

type ProgramNode struct {
    Line  int
    Title StringConstNode
    Body  ProgramBodyNode
}

func (node ProgramNode) GetLine() int {
    return node.Line
}

func (node ProgramNode) ToString() string {
    return node.Title.ToString() + node.Body.ToString() + "\n"
}

func (node *ProgramNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitProgram(node)
}

func (node *ProgramNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitProgram(node)
}
