package ast

import "plhtml/token"

type ControlFlowStmtNode struct {
    Line       int
    Type       token.Type
    Condition  ExpressionNode
    Statements []StatementNode
}

func (node ControlFlowStmtNode) GetLine() int {
    return node.Line
}

func (node ControlFlowStmtNode) ToString(lvl int) string {
    str := ident(lvl, node.Type.String()+" "+node.Condition.ToString())
    for _, stmt := range node.Statements {
        str += "\n" + stmt.ToString(lvl+1)
    }
    return str
}

func (node *ControlFlowStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitControlFlowStmt(node)
}

func (node *ControlFlowStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitControlFlowStmt(node)
}
