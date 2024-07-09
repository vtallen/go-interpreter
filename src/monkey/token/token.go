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
	ASSIGN = "="
	PLUS   = "+"

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
)
