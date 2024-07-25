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
* Function: Lexer.readIdentifier
*
* Parameters: None
*
* Returns: string - The identifier that was read from the input string (it can contain a-z, A-Z and _)
*
* Description: Reads an identifier from the input string and returns it. An identifier is defined as
*              a sequence of characters that can contain a-z, A-Z and _ without any whitespace
*
 */
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
* Function: Lexer.readNumber
*
* Parameters: None
*
* Returns: string - The number that was read from the input string (contains characters from 0-9)
*
* Description: Reads a number of an arbitrary length from the input string and returns it.
*
 */

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
* Function: Lexer.skipWhitespace
*
* Parameters: None
*
* Returns: None
*
* Description: Advances the lexer past any whitespace characters in the input string including \t, \n, \r escape codes
 */
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
* Function: Lexer.peekChar
*
* Parameters: None
*
* Retrurns: byte - the next char in the input if it exists, 0 otherwise
*
* Description: Looks at the next char in the input and returns it
*
 */
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

/*
* Function isLetter
*
* Parameters: ch byte - The character to check if it is a letter
*
* Returns: bool - True if the character is in a-z, A-Z, or is _, false otherwise
*
* Description: Checks if the given character is a letter (a-z, A-Z) or an underscore
 */
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
* Function: isDigit
*
* Parameters: ch byte - The character to check if it is a digit
*
* Returns: bool - True if the character is in 0-9, false otherwise
*
* Description: Checks if the given character is a number (0-9)
 */
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// An early return because readIdentifier advaces the readPostition and position fields of
			// the lexer past the last character of the identifier/reserved word so we do not need to call readChar again
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			// This early return is done for the same reason as the previous early return
			return tok
		} else {
			tok = newToken(token.ILLIGAL, l.ch)
		}
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
