package token

type Type int

const (
    Illegal Type = iota
    EOF

    Doctype
    Lang
    HTML
    Head
    Title
    Body
    Main
    Var
    Class
    Output
    Input
    Name
    Data
    Value
    Div
    If
    While

    Plus
    Minus
    Multiply
    Slash
    Modulo
    LParen
    RParen

    Excl
    AndOp
    OrOp

    LtOp
    GtOp
    LeqOp
    GeqOp
    EqOp
    NeqOp

    Identifier

    IntConst
    RealConst
    BoolConst
    StringConst

    DQuote
    Equal
    LessThan
    GreaterThan
)

func (tokenType Type) String() string {
    return [...]string{
        Illegal:     "Illegal",
        EOF:         "EOF",
        Doctype:     "Doctype",
        Lang:        "Lang",
        HTML:        "Html",
        Head:        "Head",
        Title:       "Title",
        Body:        "Body",
        Main:        "Main",
        Var:         "Var",
        Class:       "Class",
        Output:      "Output",
        Input:       "Input",
        Name:        "Name",
        Data:        "Data",
        Value:       "Value",
        Div:         "Div",
        If:          "If",
        While:       "While",
        Plus:        "+",
        Minus:       "-",
        Multiply:    "*",
        Slash:       "/",
        Modulo:      "%",
        LParen:      "(",
        RParen:      ")",
        Excl:        "!",
        AndOp:       "AndOp",
        OrOp:        "OrOp",
        LtOp:        "LtOp",
        GtOp:        "GtOp",
        LeqOp:       "LeqOp",
        GeqOp:       "GeqOp",
        EqOp:        "EqOp",
        NeqOp:       "NeqOp",
        Identifier:  "Identifier",
        IntConst:    "IntConst",
        RealConst:   "RealConst",
        BoolConst:   "BoolConst",
        StringConst: "StringConst",
        DQuote:      "\"",
        Equal:       "=",
        LessThan:    "<",
        GreaterThan: ">",
    }[tokenType]
}

var KeywordLexemes = map[string]Type{
    "doctype": Doctype,
    "lang":    Lang,
    "html":    HTML,
    "head":    Head,
    "title":   Title,
    "body":    Body,
    "main":    Main,
    "var":     Var,
    "class":   Class,
    "output":  Output,
    "input":   Input,
    "name":    Name,
    "data":    Data,
    "value":   Value,
    "div":     Div,
    "if":      If,
    "while":   While,
}

var BoolOpLexemes = map[string]Type{
    "&lt;":     LtOp,
    "&gt;":     GtOp,
    "&leq;":    LeqOp,
    "&geq;":    GeqOp,
    "&equals;": EqOp,
    "&ne;":     NeqOp,
    "&and;":    AndOp,
    "&or;":     OrOp,
}
