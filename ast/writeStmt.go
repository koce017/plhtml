package ast

type WriteStmtNode struct {
    Line  int
    Value ExpressionNode
}

func (node WriteStmtNode) GetLine() int {
    return node.Line
}

func (node WriteStmtNode) ToString(lvl int) string {
    return ident(lvl, "Write: "+node.Value.ToString())
}

func (node *WriteStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitWriteStmt(node)
}

func (node *WriteStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitWriteStmt(node)
}
