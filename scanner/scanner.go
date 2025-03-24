package scanner

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "unicode"

    "plhtml/logger"
    "plhtml/token"
    "plhtml/util"
)

var myLogger = logger.New("SCANNER")

func (scanner *Scanner) SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Scanner struct {
    line       int
    column     int
    prevColumn int
    tabSize    int
    index      int
    source     []rune
}

type Token struct {
    Type    token.Type
    Line    int
    Column  int
    IntVal  int
    BoolVal bool
    RealVal float64
    StrVal  string
}

type tokenStartPair struct {
    line   int
    column int
}

func New() *Scanner {
    scanner := new(Scanner)
    scanner.tabSize = 4
    return scanner
}

func (scanner *Scanner) SetTabSize(tabSize int) {
    scanner.tabSize = tabSize
}

func (scanner *Scanner) Scan(source string) []Token {
    scanner.line = 1
    scanner.index = 0
    scanner.column = 0
    scanner.source = []rune(source)

    var tokens []Token

    for {
        tok := scanner.nextToken()
        tokens = append(tokens, tok)
        if tok.Type == token.EOF {
            break
        }
    }

    return tokens
}

func (scanner *Scanner) nextToken() Token {

    for {

        ch, ok := scanner.nextChar()

        if !ok {
            break
        }

        if unicode.IsSpace(ch) {
            continue
        }

        tokenStart := tokenStartPair{
            line:   scanner.line,
            column: scanner.column,
        }

        if ch == '`' {
            return scanner.lexString(ch, tokenStart)
        }

        // after lexString because <!-- could be part of string literal
        if ch == '<' && scanner.lexComment(tokenStart) {
            continue
        }

        if unicode.IsDigit(ch) {
            return scanner.lexNumber(ch, tokenStart)
        }

        if ch == '&' || unicode.IsLetter(ch) {
            return scanner.lexWord(ch, tokenStart)
        }

        switch ch {
        case '+':
            return newTok(token.Plus, tokenStart)
        case '-':
            return newTok(token.Minus, tokenStart)
        case '*':
            return newTok(token.Multiply, tokenStart)
        case '/':
            return newTok(token.Slash, tokenStart)
        case '%':
            return newTok(token.Modulo, tokenStart)
        case '(':
            return newTok(token.LParen, tokenStart)
        case ')':
            return newTok(token.RParen, tokenStart)
        case '!':
            return newTok(token.Excl, tokenStart)
        case '"':
            return newTok(token.DQuote, tokenStart)
        case '=':
            return newTok(token.Equal, tokenStart)
        case '<':
            return newTok(token.LessThan, tokenStart)
        case '>':
            return newTok(token.GreaterThan, tokenStart)
        default:
            panic(fmt.Sprintf("Illegal character %c at %d:%d.", ch, tokenStart.line, tokenStart.column))
        }

    }

    return Token{Type: token.EOF}
}

func (scanner *Scanner) lexString(ch rune, tokenStart tokenStartPair) Token {
    var ok bool
    str := string(ch)

    for {
        ch, ok = scanner.nextChar()

        if ok {

            if ch == '`' {
                break
            }

            if ch == '\n' {
                panic(fmt.Sprintf("Error at %d:%d: newline in string", scanner.line, scanner.column))
            }

            if ch == '\\' {
                ech, err := scanner.lexEscape()
                if err != nil {
                    panic(fmt.Sprintf("Error at %d:%d: %s", scanner.line, scanner.column-1, err))
                }
                str += ech
            } else {
                str += string(ch)
            }

        } else {
            panic(fmt.Sprintf("Unterminated string at %d:%d.", tokenStart.line, tokenStart.column))
        }
    }

    tok := newTok(token.StringConst, tokenStart)
    tok.StrVal = str[1:]
    return tok
}

func (scanner *Scanner) lexEscape() (string, error) {
    ch, ok := scanner.nextChar()

    if ok {
        switch ch {
        case '\\':
            return "\\", nil
        case 't':
            return "\t", nil
        case 'n':
            return "\n", nil
        }
    }

    return "", errors.New("invalid escape sequence")
}

func (scanner *Scanner) lexComment(tokenStart tokenStartPair) bool {
    // <!--.*-->

    ch1, ok1 := scanner.lookahead(0)
    ch2, ok2 := scanner.lookahead(1)
    ch3, ok3 := scanner.lookahead(2)

    if ok1 && ok2 && ok3 && ch1 == '!' && ch2 == '-' && ch3 == '-' {
        for {
            ch, ok := scanner.nextChar()
            if !ok {
                panic(fmt.Sprintf("End of file inside comment at %d:%d.", tokenStart.line, tokenStart.column))
            }
            if ch == '-' {
                ch1, ok1 := scanner.lookahead(0)
                ch2, ok2 := scanner.lookahead(1)
                if ok1 && ok2 && ch1 == '-' && ch2 == '>' {
                    scanner.nextChar()
                    scanner.nextChar()
                    return true
                }
            }
        }
    }

    return false
}

func (scanner *Scanner) lexNumber(ch rune, tokenStart tokenStartPair) Token {
    var ok bool
    isReal := false
    number := "" + string(ch)

    for ch, ok = scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

        if ch == '.' && !isReal {
			number += "."
			isReal = true
		} else if unicode.IsNumber(ch) {
            number += string(ch)
        } else {
            scanner.goBack()
            break
        }

    }

    if isReal {
        realVal, err := util.StrToFloat64(number)
        util.Check(err)
        tok := newTok(token.RealConst, tokenStart)
        tok.RealVal = realVal
        return tok
    }

    intVal, err := strconv.Atoi(number)
    util.Check(err)
    tok := newTok(token.IntConst, tokenStart)
    tok.IntVal = intVal
    return tok
}

func (scanner *Scanner) lexWord(ch rune, tokenStart tokenStartPair) Token {
    // &[a-zA-Z];            operators
    // [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords

    word := string(ch)

    for ch, ok := scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

        valid := unicode.IsLetter(ch) ||
            (word[0] == '&' && ch == ';') ||
            (word[0] != '&' && unicode.IsNumber(ch))

        if valid {
            word += string(ch)
        } else {
            scanner.goBack()
            break
        }

    }

    if tok, ok := scanner.lexBoolOperator(word, tokenStart); ok {
        return tok
    }

    word = strings.ToLower(word)

    if word == "true" || word == "false" {
        tok := newTok(token.BoolConst, tokenStart)
        tok.BoolVal = word == "true"
        return tok
    }

    if tok, ok := token.KeywordLexemes[word]; ok {
        return newTok(tok, tokenStart)
    }

    tok := newTok(token.Identifier, tokenStart)
    tok.StrVal = word
    return tok
}

func (scanner *Scanner) lexBoolOperator(word string, tokenStart tokenStartPair) (Token, bool) {

    firstChar := word[0]
    lastChar := word[len(word)-1:]

    if firstChar != '&' {
        return Token{Type: token.Illegal}, false
    }

    if lastChar != ";" {
        panic(fmt.Sprintf("Unterminated operator %s at %d:%d.", word, tokenStart.line, tokenStart.column))
    }

    if tok, ok := token.BoolOpLexemes[word]; ok {
        return newTok(tok, tokenStart), true
    }

    panic(fmt.Sprintf("Operator %s is not valid at %d:%d.", word, tokenStart.line, tokenStart.column))

}

func newTok(tokType token.Type, tokenStart tokenStartPair) Token {
    return Token{Type: tokType, Line: tokenStart.line, Column: tokenStart.column}
}

func (scanner *Scanner) lookahead(i int) (rune, bool) {
    if scanner.index+i < len(scanner.source)-1 {
        return scanner.source[scanner.index+i], true
    }
    return 0, false
}

func (scanner *Scanner) nextChar() (rune, bool) {

    if scanner.index >= len(scanner.source)-1 {
        return 0, false
    }

    ch := scanner.source[scanner.index]
    scanner.index++

    myLogger.Debug("Read char: %s\n", strconv.Quote(string(ch)))

    if ch == '\n' {
        scanner.prevColumn = scanner.column
        scanner.column = 0
        scanner.line++
    } else if ch == '\t' {
        scanner.column += scanner.tabSize
    } else {
        scanner.column++
    }

    return ch, true

}

func (scanner *Scanner) goBack() {
    if scanner.index < len(scanner.source) {

        scanner.index--
        ch := scanner.source[scanner.index]

        myLogger.Debug("Unread char: %s\n", strconv.Quote(string(ch)))

        if ch == '\n' {
            scanner.column = scanner.prevColumn
            scanner.line--
        } else if ch == '\t' {
            scanner.column -= scanner.tabSize
        } else {
            scanner.column--
        }

    }
}
