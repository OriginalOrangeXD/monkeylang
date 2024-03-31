package token

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"
    // identifiers and literals
    IDENT   = "IDENT"
    INT     = "INT"
    // operators
    ASSIGN  = "="
    PLUS    = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"
    NOT_EQ   = "!="
    LT = "<"
    GT = ">"
    EQ = "=="
    // delim
    COMMA     = ","
    SEMICOLON = ";"
    COLON     = ":"
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    LBRACKET   = "["
    RBRACKET   = "]"
    //keywords
    FUNCTION  = "FUNCTION"
    LET       = "LET"
    TRUE      = "TRUE"
    FALSE     = "FALSE"
    IF        = "IF"
    ELSE      = "ELSE"
    RETURN    = "RETURN"
    STRING    = "STRING"
    // Operators
)

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

var keywords = map[string] TokenType {
    "fn":     FUNCTION,
    "let":    LET,
    "true":   TRUE,
    "false":  FALSE,
    "if":     IF,
    "else":   ELSE,
    "return": RETURN,
    "string": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}
