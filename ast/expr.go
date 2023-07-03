package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Ident struct {
	Token token.Token
	Value string
}

func (i *Ident) exprNode()            {}
func (i *Ident) TokenLiteral() string { return i.Token.Literal }
func (i *Ident) String() string       { return i.Value }

type IntLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntLiteral) exprNode()            {}
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string       { return il.Token.Literal }

type PrefixExpr struct {
	Token token.Token
	Op    string
	Right Expr
}

func (pe *PrefixExpr) exprNode()            {}
func (pe *PrefixExpr) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpr) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Op)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpr struct {
	Token token.Token
	Left  Expr
	Op    string
	Right Expr
}

func (oe *InfixExpr) exprNode()            {}
func (oe *InfixExpr) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpr) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Op + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) exprNode()            {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type IfExpr struct {
	Token       token.Token
	Condition   Expr
	Consequence *BlockStmt
	Alternative *BlockStmt
}

func (ie *IfExpr) exprNode()            {}
func (ie *IfExpr) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpr) String() string {
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

type FuncLiteral struct {
	Token token.Token
	Params []*Ident
	Body *BlockStmt
}

func (fl *FuncLiteral) exprNode() {}
func (fl *FuncLiteral) TokenLiteral() string { return fl.Token.Literal}
func (fl *FuncLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Params {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

type CallExpr struct {
	Token token.Token
	Func Expr
	Args []Expr
}

func (ce *CallExpr) exprNode() {}
func (ce *CallExpr) TokenLiteral() string { return ce.Token.Literal}
func (ce *CallExpr) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Args {
		args = append(args, a.String())
	}

	out.WriteString(ce.Func.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
