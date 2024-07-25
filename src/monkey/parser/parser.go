/*
* File: parser/parser.go
*
* Description: Contains the parser for the monkey programming language
*
 */

package parser

import (
	"fmt"

	"github.com/vtallen/go-interpreter/ast"
	"github.com/vtallen/go-interpreter/lexer"
	"github.com/vtallen/go-interpreter/token"
)

/*
* Struct: Parser
*
* Description: The parser struct for the monkey programming language
 */
type Parser struct {
	l *lexer.Lexer // Pointer to the lexer

	curToken  token.Token // Current token under consideration
	peekToken token.Token // Next token in the program, used to figure out what to do

	errors []string // Any arrors that occur during parsing
}

/*
* Function: New
*
* Parameters: l *lexer.Lexer - Pointer to the lexer of the program
*
* Returns: *Parser - Pointer to the parser created
*
* Description: Creates a new parser for the monkey programming
 */
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read to tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

/*
* Function: Parser.Errors
*
* Parameters: none
*
* Returns: []string - Array of strings containing the errors that occured during parsing
*
* Description: Returns the errors that occured during parsing
*
 */
func (p *Parser) Errors() []string {
	return p.errors
}

/*
* Function: Parser.nextToken
*
* Parameters: none
*
* Returns: none
*
* Description: Advances the current token and peek token by one token
 */
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
* Function: Parser.ParseProgram
*
* Parameters: none
*
* Returns: *ast.Program - Pointer to the AST of the program
*
* Description: Parses the program and returns the AST
*
 */
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Skipping expressions for now
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Skipping expressions for now
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

/*
* Function Parser.peekError
*
* Parameters: t token.TokenType - The token type that was expected
*
* Returns: none
*
* Description: Adds an error to the parser's error array when a token of t was expected and something else was found
*
 */
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
