/*
* File: token/token.go
*
* Description:
*
 */

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLIGAL = "ILLEGAL" // Represents a token the parser does not know about
	EOF     = "EOF"     // Represents the end of a file and tells the parser when to stop

	// Identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // literals like: 1234

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiter characters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords: reserved words that have meaning that are not variables
	FUNCTION = "FUNCTION" // functions defined as fn()
	LET      = "LET"      // Variable declaration like "let five = 5;"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Contains a map of reserved words for the language and their corresponding TokenType
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

/*
* Function: LookupIdent
*
* Paremters: ident string - The identifier to look up
*
* Returns: TokenType - The type of the identifier. Returns IDENT if ident is not a reserved word
*
* Description: Looks up the identifier in the keywords map and returns the corresponding TokenType. If ident is not a reserved word, we can
*              conclude it is a variable name and IDENT will be returned
 */
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
