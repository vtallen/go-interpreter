/*
* File: ast/as.go
*
* Description: This file defines the abstract syntax tree (AST) for the Monkey programming language.
 */

package ast

import "github.com/vtallen/go-interpreter/token"

/*
* Interface: Node
*
* Description: This struct defines the interface for all nodes in the abstract syntax tree (AST).
 */
type Node interface {
	TokenLiteral() string // All nodes should be able to return the literal value of the token that they are associated with.
}

/*
* Interface: Statement
*
* Description: Any node that is a statment will implement the statementNode interface.
 */
type Statement interface {
	Node            // All statements are nodes.
	statementNode() // All statements will have a statementNode method, which is a dummy method to help the
	// go compiler throw errors if we use a statement where an expression should go
}

type Expression interface {
	Node             // All expressions are nodes.
	expressionNode() // All expressions will have an expressionNode method, which is a dummy method to help the
	// go compiler throw errors if we use an expression where a statement should go
}

/*
* Struct Type: Program
*
* Description: This struct represents the entire program. It is the root node of the AST.
 */
type Program struct {
	Statements []Statement // All statements in the program
}

/*
* Function: Program.TokenLiteral
*
* Paramters: none
*
* Return: string - The literal value of the token associated with the first statement in the program.
*
* Description: This function returns the literal value of the token associated with the first statement in the program.
*
 */
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/*
* Struct: LetStatement
*
* Implements: Statement
*
* Description: This struct represents a let statement in the Monkey programming language.
*
 */
type LetStatement struct {
	Statement
	Token token.Token
	Name  *Identifier
	// TODO: should this be a pointer to an expression or a value?
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

/*
* Struct: Identifier
*
* Implements: Expression
*
* Description: This struct represents an identifier in the Monkey programming language.
*
 */
type Identifier struct {
	Expression
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

/*
* Struct: ReturnStatement
*
* Implements: Statement
*
* Description: This struct represents a return statement in the Monkey programming language.
*
 */
type ReturnStatement struct {
	Statement
	Token       token.Token // The 'return' token
	ReturnValue Expression  // The expression to return
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
