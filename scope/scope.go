package scope

import (
    "fmt"
    "go/constant"
    "plhtml/logger"
    "strings"
)

var myLogger = logger.New("SCOPE")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Scope struct {
    Id      int
    Parent  *Scope
    symbols map[string]*Symbol
}

type Symbol struct {
    Line  int
    Name  string
    Type  string
    Value constant.Value
}

const (
    TypeInteger = "integer"
    TypeReal    = "real"
    TypeBoolean = "boolean"
    TypeString  = "string"
)

func New(id int, parent *Scope) *Scope {
    scope := &Scope{
        Id:     id,
        Parent: parent,
    }
    if parent == nil {
        myLogger.Debug("Creating root scope %d.", id)
        scope.symbols = map[string]*Symbol{
            TypeInteger: {Name: TypeInteger},
            TypeReal:    {Name: TypeReal},
            TypeBoolean: {Name: TypeBoolean},
            TypeString:  {Name: TypeString},
        }
    } else {
        myLogger.Debug("Creating new scope %d as child of %d.", id, parent.Id)
        scope.symbols = make(map[string]*Symbol)
    }
    return scope
}

func (scope *Scope) Insert(sym *Symbol) {
    varId := fmt.Sprintf("%s$%d", sym.Name, sym.Line)
    myLogger.Debug("Inserting %s into scope %d.", varId, scope.Id)
    scope.symbols[varId] = sym
}

func (scope *Scope) DeclaredLocally(name string) bool {
    myLogger.Debug("Checking if %s is declared in scope %d.", name, scope.Id)
    for _, sym := range scope.symbols {
        if name == sym.Name {
            return true
        }
    }
    return false
}

func (scope *Scope) Lookup(name string) (*Symbol, string, bool) {
    for currentScope := scope; currentScope != nil; currentScope = currentScope.Parent {
        myLogger.Debug("Looking for %s in scope %d.", name, currentScope.Id)

        if strings.Contains(name, "$") {
            if sym, ok := currentScope.symbols[name]; ok {
                myLogger.Debug("Found %s in scope %d.", name, currentScope.Id)
                return sym, name, true
            }
        } else {
            for _, sym := range currentScope.symbols {
                if name == sym.Name {
                    if currentScope.Parent != nil {
                        name = fmt.Sprintf("%s$%d", sym.Name, sym.Line)
                    }
                    myLogger.Debug("Found %s in scope %d.", name, currentScope.Id)
                    return sym, name, true
                }
            }
        }

    }
    myLogger.Debug("Symbol %s not declared.", name)
    return nil, "", false
}

func (scope *Scope) GetValue(name string) constant.Value {
    myLogger.Debug("Getting value of %s", name)
    sym, _, _ := scope.Lookup(name)
    return sym.Value
}

func (scope *Scope) SetValue(name string, value constant.Value) {
    myLogger.Debug("Setting value for %s", name)
    sym, _, _ := scope.Lookup(name)
    sym.Value = value
}

func (symbol *Symbol) SetValue(value constant.Value) {
    myLogger.Debug("Setting value for %s", symbol.Name)
    symbol.Value = value
}