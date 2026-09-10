package ast

type MainFuncNode struct {
	Line       int
	Statements []StatementNode
}

func (node MainFuncNode) GetLine() int {
	return node.Line
}

func (node MainFuncNode) ToString() string {
	str := ""
	for _, stmt := range node.Statements {
		str += stmt.ToString(0) + "\n"
	}
	return str
}

func (node *MainFuncNode) AcceptAnalyzer(analyzer IAnalyzer) {
	analyzer.VisitMainFunc(node)
}

func (node *MainFuncNode) AcceptInterpreter(interp IInterpreter) {
	interp.VisitMainFunc(node)
}
