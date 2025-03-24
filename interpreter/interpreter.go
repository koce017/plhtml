package interpreter

import (
    "bufio"
    "errors"
    "fmt"
    "go/constant"
    "os"
    "plhtml/scope"
    "strconv"

    "plhtml/ast"
    "plhtml/logger"
    "plhtml/token"
    "plhtml/util"
)

type TokenType = token.Type

var myLogger = logger.New("INTER")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Interpreter struct {
    in *bufio.Scanner
}

func New(file *os.File) *Interpreter {
    interp := new(Interpreter)
    interp.in = bufio.NewScanner(file)
    return interp
}

func (interp *Interpreter) VisitBinaryOpExpr(node *ast.BinaryOpExprNode) constant.Value {
    leftValue := node.LeftExpr.AcceptInterpreter(interp).(constant.Value)
    rightValue := node.RightExpr.AcceptInterpreter(interp).(constant.Value)
    myLogger.Debug("%s %s %s", leftValue.String(), node.Operator.String(), rightValue.String())

    var err error
    var value constant.Value

    if isStr(leftValue) || isStr(rightValue) {
        value, err = opsWithStrings(leftValue, node.Operator, rightValue)
    } else if leftValue.Kind() == constant.Int || rightValue.Kind() == constant.Int {
        value, err = opsWithIntegers(leftValue, node.Operator, rightValue)
    } else if isNum(leftValue) && isNum(rightValue) {
        value, err = opsWithFloats(leftValue, node.Operator, rightValue)
    } else if isBool(leftValue) && isBool(rightValue) {
        value, err = opsWithBooleans(leftValue, node.Operator, rightValue)
    } else {
        panic(fmt.Sprintf("Error on line %d: operator %s is not supported for operands of given types", node.GetLine(), node.Operator.String()))
    }

    if err != nil {
        panic(fmt.Sprintf("Error on line %d: %s", node.GetLine(), err))
    }

    return value
}

func (interp *Interpreter) VisitBoolConst(node *ast.BoolConstNode) constant.Value {
    myLogger.Debug(fmt.Sprintf("Bool %t", node.Value))
    return constant.MakeBool(node.Value)
}

func (interp *Interpreter) VisitControlFlowStmt(node *ast.ControlFlowStmtNode) {
    myLogger.Debug(node.Type.String())
    switch node.Type {
    case token.If:
        if constant.BoolVal(node.Condition.AcceptInterpreter(interp).(constant.Value)) {
            for _, stmt := range node.Statements {
                stmt.AcceptInterpreter(interp)
            }
        }
        break
    case token.While:
        for constant.BoolVal(node.Condition.AcceptInterpreter(interp).(constant.Value)) {
            for _, stmt := range node.Statements {
                stmt.AcceptInterpreter(interp)
            }
        }
    }
}

func (interp *Interpreter) VisitIdentifier(node *ast.IdentifierNode) constant.Value {
    value := node.Scope.GetValue(node.Name)
    myLogger.Debug("%s: %s", node.Name, value.String())
    return value
}

func (interp *Interpreter) VisitIntConst(node *ast.IntConstNode) constant.Value {
    return constant.MakeInt64(int64(node.Value))
}

func (interp *Interpreter) VisitMainFunc(node *ast.MainFuncNode) {
    for _, stmt := range node.Statements {
        stmt.AcceptInterpreter(interp)
    }
}

func (interp *Interpreter) VisitProgram(node *ast.ProgramNode) {
    fmt.Println(node.Title.Value)
    node.Body.AcceptInterpreter(interp)
}

func (interp *Interpreter) VisitProgramBody(node *ast.ProgramBodyNode) {
    node.MainFunc.AcceptInterpreter(interp)
}

func (interp *Interpreter) VisitReadStmt(node *ast.ReadStmtNode) {
    interp.in.Scan()

    var value constant.Value
    sym, _, _ := node.Scope.Lookup(node.Identifier.Name)

    switch sym.Type {
    case scope.TypeInteger:
        if val, err := util.StrToInt64(interp.in.Text()); err == nil {
            value = constant.MakeInt64(val)
        } else {
            panic("Integer expected.")
        }
    case scope.TypeReal:
        if val, err := util.StrToFloat64(interp.in.Text()); err == nil {
            value = constant.MakeFloat64(val)
        } else {
            panic("Real number expected.")
        }
    case scope.TypeBoolean:
        if val, err := util.StrToBool(interp.in.Text()); err == nil {
            value = constant.MakeBool(val)
        } else {
            panic("true or false expected.")
        }
    case scope.TypeString:
        value = constant.MakeString(interp.in.Text())
    default:
        panic(fmt.Sprintf("Error at line %d: You can only read input for primitive types.", node.Line))
    }

    node.Scope.SetValue(node.Identifier.Name, value)
}

func (interp *Interpreter) VisitRealConst(node *ast.RealConstNode) constant.Value {
    return constant.MakeFloat64(node.Value)
}

func (interp *Interpreter) VisitStringConst(node *ast.StringConstNode) constant.Value {
    return constant.MakeString(node.Value)
}

func (interp *Interpreter) VisitUnaryExpr(node *ast.UnaryExprNode) constant.Value {
    exprValue := node.Expr.AcceptInterpreter(interp).(constant.Value)

    switch node.Operator {
    case token.Minus:
        if exprValue.Kind() == constant.Int {
            exprVal, _ := constant.Int64Val(exprValue)
            return constant.MakeInt64(-exprVal)
        } else if exprValue.Kind() == constant.Float {
            exprVal, _ := constant.Float64Val(exprValue)
            return constant.MakeFloat64(-exprVal)
        }
        break
    case token.Excl:
        if exprValue.Kind() == constant.Bool {
            return constant.MakeBool(!constant.BoolVal(exprValue))
        }
    }

    panic(fmt.Sprintf("Error on line %d: operator %s is not supported for expression of given type", node.GetLine(), node.Operator.String()))
}

func (interp *Interpreter) VisitVarAssign(node *ast.VarAssignNode) {
    value := node.Value.AcceptInterpreter(interp).(constant.Value)
    myLogger.Debug("%s = %s", node.Identifier.Name, value.String())
    node.Scope.SetValue(node.Identifier.Name, value)
}

func (interp *Interpreter) VisitVarDecl(node *ast.VarDeclNode) {
    myLogger.Debug("%s %s", node.Type.Name, node.Identifier.Name)
    symbol, _, _ := node.Scope.Lookup(node.Identifier.Name)
    switch symbol.Type {
    case scope.TypeInteger:
        symbol.SetValue(constant.MakeInt64(0))
    case scope.TypeReal:
        symbol.SetValue(constant.MakeFloat64(0))
    case scope.TypeBoolean:
        symbol.SetValue(constant.MakeBool(false))
    case scope.TypeString:
        symbol.SetValue(constant.MakeString(""))
    default:
        panic(fmt.Sprintf("Error at line %d: Assignment is only supported for primitive types.", node.Line))
    }
}

func (interp *Interpreter) VisitWriteStmt(node *ast.WriteStmtNode) {
    exprValue := node.Value.AcceptInterpreter(interp).(constant.Value)
    fmt.Print(constant.StringVal(exprValue))
}

func isNum(val constant.Value) bool {
    return val.Kind() == constant.Int || val.Kind() == constant.Float
}

func isStr(val constant.Value) bool {
    return val.Kind() == constant.String
}

func isBool(val constant.Value) bool {
    return val.Kind() == constant.Bool
}

func opsWithStrings(left constant.Value, operator TokenType, right constant.Value) (constant.Value, error) {
    if operator == token.Plus {
        return strcat(left, right)
    }

    if isStr(left) && isStr(right) {
        leftVal := constant.StringVal(left)
        rightVal := constant.StringVal(right)

        switch operator {
        case token.EqOp:
            return constant.MakeBool(leftVal == rightVal), nil
        case token.NeqOp:
            return constant.MakeBool(leftVal != rightVal), nil
        }
    }

    return nil, errors.New("operator " + operator.String() + " cannot be applied to two strings")
}

func opsWithIntegers(left constant.Value, operator TokenType, right constant.Value) (constant.Value, error) {
    leftVal, _ := constant.Int64Val(left)
    rightVal, _ := constant.Int64Val(right)
    switch operator {
    case token.Plus:
        return constant.MakeInt64(leftVal + rightVal), nil
    case token.Minus:
        return constant.MakeInt64(leftVal - rightVal), nil
    case token.Multiply:
        return constant.MakeInt64(leftVal * rightVal), nil
    case token.Slash:
        return constant.MakeInt64(leftVal / rightVal), nil
    case token.Modulo:
        return constant.MakeInt64(leftVal % rightVal), nil
    case token.LtOp:
        return constant.MakeBool(leftVal < rightVal), nil
    case token.GtOp:
        return constant.MakeBool(leftVal > rightVal), nil
    case token.LeqOp:
        return constant.MakeBool(leftVal <= rightVal), nil
    case token.GeqOp:
        return constant.MakeBool(leftVal >= rightVal), nil
    case token.EqOp:
        return constant.MakeBool(leftVal == rightVal), nil
    case token.NeqOp:
        return constant.MakeBool(leftVal != rightVal), nil
    default:
        return nil, errors.New("operator " + operator.String() + " cannot be applied to two integers")
    }
}

func opsWithFloats(left constant.Value, operator TokenType, right constant.Value) (constant.Value, error) {
    leftVal, _ := constant.Float64Val(left)
    rightVal, _ := constant.Float64Val(right)
    switch operator {
    case token.Plus:
        return constant.MakeFloat64(leftVal + rightVal), nil
    case token.Minus:
        return constant.MakeFloat64(leftVal - rightVal), nil
    case token.Multiply:
        return constant.MakeFloat64(leftVal * rightVal), nil
    case token.Slash:
        return constant.MakeFloat64(leftVal / rightVal), nil
    case token.LtOp:
        return constant.MakeBool(leftVal < rightVal), nil
    case token.GtOp:
        return constant.MakeBool(leftVal > rightVal), nil
    case token.LeqOp:
        return constant.MakeBool(leftVal <= rightVal), nil
    case token.GeqOp:
        return constant.MakeBool(leftVal >= rightVal), nil
    case token.EqOp:
        return constant.MakeBool(leftVal == rightVal), nil
    case token.NeqOp:
        return constant.MakeBool(leftVal != rightVal), nil
    default:
        return nil, errors.New("operator " + operator.String() + " cannot be applied to real numbers")
    }
}

func opsWithBooleans(left constant.Value, operator TokenType, right constant.Value) (constant.Value, error) {
    leftVal := constant.BoolVal(left)
    rightVal := constant.BoolVal(right)
    switch operator {
    case token.EqOp:
        return constant.MakeBool(leftVal == rightVal), nil
    case token.NeqOp:
        return constant.MakeBool(leftVal != rightVal), nil
    case token.AndOp:
        return constant.MakeBool(leftVal && rightVal), nil
    case token.OrOp:
        return constant.MakeBool(leftVal || rightVal), nil
    default:
        return nil, errors.New("operator " + operator.String() + " cannot be applied to booleans")
    }
}

func strcat(leftValue constant.Value, rightValue constant.Value) (constant.Value, error) {
    if leftValue.Kind() == constant.String && rightValue.Kind() == constant.String {
        return constant.MakeString(constant.StringVal(leftValue) + constant.StringVal(rightValue)), nil
    }
    if leftValue.Kind() == constant.String {
        return strcatImplicitConversion(leftValue, rightValue, true)
    }
    if rightValue.Kind() == constant.String {
        return strcatImplicitConversion(rightValue, leftValue, false)
    }
    return nil, errors.New("could not perform string concatenation")
}

func strcatImplicitConversion(stringValue constant.Value, otherValue constant.Value, ordered bool) (constant.Value, error) {
    stringVal := constant.StringVal(stringValue)
    switch otherValue.Kind() {
    case constant.Int:
        intVal, _ := constant.Int64Val(otherValue)
        intValStr := strconv.FormatInt(intVal, 10)
        if ordered {
            return constant.MakeString(stringVal + intValStr), nil
        } else {
            return constant.MakeString(intValStr + stringVal), nil
        }
    case constant.Float:
        floatVal, _ := constant.Float64Val(otherValue)
        floatValStr := util.FloatToString(floatVal)
        if ordered {
            return constant.MakeString(stringVal + floatValStr), nil
        } else {
            return constant.MakeString(floatValStr + stringVal), nil
        }
    case constant.Bool:
        boolVal := constant.BoolVal(otherValue)
        boolValStr := strconv.FormatBool(boolVal)
        if ordered {
            return constant.MakeString(stringVal + boolValStr), nil
        } else {
            return constant.MakeString(boolValStr + stringVal), nil
        }
    }
    return nil, errors.New("you can concatenate string with other strings, integers, floats and booleans")
}
