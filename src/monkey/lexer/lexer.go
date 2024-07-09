/*
* File: lexer/lexer.go
*
* Description:
*
 */
package lexer

import "github.com/vtallen/go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to the current char)
	readPosition int  // current reading position in input (after current char, the next char to be read)
	ch           byte // The current character under examination (char at position in input)
}

/*
* Function: New
*
* Parameters: input string - The input source code to be lexed
*
* Returns: *Lexer - A pointer to a new Lexer object
*
* Description: Creates a new Lexer object with the given input
 */
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Put the lexer into a usable state before NextToken can be called
	return l
}

/*
* Function: Lexer.readChar
*
* Parameters: None
*
* Returns: None
*
* Description: Reads the next character in the input string and advances the lexer's position in the input string.
*
 */
func (l *Lexer) readChar() {
	// This if statement checks if the readPosition is greater than or equal to the length of the input string.
	// If it is, then the lexer has reached the end of the input and sets the current character to 0,
	// which is the ASCII code for the "NUL" character and has no meaning in Monkey.
	// Otherwise, it sets the current character to the character at the current readPosition in the input string.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition // Move the lexer to the next character

	l.readPosition += 1 // Increment the "pointer" to the next character
}

/*
* Function: Lexer.NextToken
*
* Parameters: None
*
* Returns: token.Token - The next token in the input string
*
* Description: Returns the next token in the input string and advances the Lexer to the next character in input
*
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// Move the lexer to the next character
	l.readChar()

	return tok
}

/*
* Function: newToken
*
* Parameters: tokenType token.TokenType - The type of token to create
*             ch        byte            - The character to create the token from
*
* Returns: token.Token - A new token of the given type and character
*
* Description: Helper function that creates a new token of the given type and character
*
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
