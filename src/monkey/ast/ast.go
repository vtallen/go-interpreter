/*
* File: ast/as.go
*
* Description: This file defines the abstract syntax tree (AST) for the Monkey programming language.
 */

package ast

import (
	"bytes"
	"strings"

	"github.com/vtallen/go-interpreter/token"
)

/*
* Interface: Node
*
* Description: This struct defines the interface for all nodes in the abstract syntax tree (AST).
 */
type Node interface {
	TokenLiteral() string // All nodes should be able to return the literal value of the token that they are associated with.
	String() string       // All nodes should be able to return a string representation of the node.
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
* Struct: Statement
*
* Implements: Statement
*
* Description: This struct is a dummy struct that will be embedded in all other statement structs. It is used to help the go compiler throw errors if we use a statement where an expression should go.
 */
type ExpressionStatement struct {
	Statement
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
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
* Function: Program.String
*
* Paramters: none
*
* Return: string - A string representation of the program's AST nodes.
*
* Description: This function returns a string representation of the program's AST nodes. For each type of statement,
*           the String() method of that statement is called and the result is concatenated to the output string.
 */
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

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

func (i *Identifier) String() string { return i.Value }

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

/*
* Function: ReturnStatement.String
*
* Paramters: none
*
* Return: string - A string representation of the return statement.
*
* Description: This function returns a string representation of the return statement. If there is a return value, it is included in the output string.
*
 */
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

/*
* Struct: IntegerLiteral
*
* Implements: Expression, Statement
*
* Description: This struct represents an integer literal in the Monkey programming language.
 */
type IntegerLiteral struct {
	Expression
	Statement
	Token token.Token // The token.INT token
	Value int64       // The underlying value of the integer
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
