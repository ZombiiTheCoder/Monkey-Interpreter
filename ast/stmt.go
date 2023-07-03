package ast

import (
	"bytes"
	"monkey/token"
)

type LetStmt struct {
	Token token.Token
	Name  *Ident
	Value Expr
}

func (ls *LetStmt) stmtNode()            {}
func (ls *LetStmt) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStmt) String() string {
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

type ReturnStmt struct {
	Token       token.Token
	ReturnValue Expr
}

func (rs *ReturnStmt) stmtNode()            {}
func (rs *ReturnStmt) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStmt) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExprStmt struct {
	Token token.Token
	Expr  Expr
}

func (es *ExprStmt) stmtNode()            {}
func (es *ExprStmt) TokenLiteral() string { return es.Token.Literal }

func (es *ExprStmt) String() string {
	if es.Expr != nil {
		return es.Expr.String()
	}
	return ""
}

type BlockStmt struct {
	Token token.Token
	Stmts []Stmt
}

func (bs *BlockStmt) stmtNode() {}
func (bs *BlockStmt) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStmt) String() string {
	var out bytes.Buffer

	for _, s := range bs.Stmts {
		out.WriteString(s.String())
	}

	return out.String()
}