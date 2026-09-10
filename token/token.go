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
		Illegal:     "illegal",
		EOF:         "eof",
		Doctype:     "doctype",
		Lang:        "lang",
		HTML:        "html",
		Head:        "head",
		Title:       "title",
		Body:        "body",
		Main:        "main",
		Var:         "var",
		Class:       "class",
		Output:      "output",
		Input:       "input",
		Name:        "name",
		Data:        "data",
		Value:       "value",
		Div:         "div",
		If:          "if",
		While:       "while",
		Plus:        "+",
		Minus:       "-",
		Multiply:    "*",
		Slash:       "/",
		Modulo:      "%",
		LParen:      "(",
		RParen:      ")",
		Excl:        "!",
		AndOp:       "&&",
		OrOp:        "||",
		LtOp:        "<",
		GtOp:        ">",
		LeqOp:       "<=",
		GeqOp:       ">=",
		EqOp:        "===",
		NeqOp:       "!==",
		Identifier:  "identifier",
		IntConst:    "intConst",
		RealConst:   "realConst",
		BoolConst:   "boolConst",
		StringConst: "stringConst",
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
