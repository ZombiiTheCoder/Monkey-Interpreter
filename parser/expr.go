package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/token"
	"strconv"
)

func (p *Parser) parseExpr(precedence int) ast.Expr {

	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp

}

func (p *Parser) parseIdent() ast.Expr {
	return &ast.Ident{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntLiteral() ast.Expr {
	lit := &ast.IntLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0 ,64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpr() ast.Expr {
	expr := &ast.PrefixExpr{
		Token: p.curToken,
		Op: p.curToken.Literal,
	}

	p.nextToken()

	expr.Right = p.parseExpr(PREFIX)

	return expr
}

func (p *Parser) parseInfixExpr(left ast.Expr) ast.Expr {
	expr := &ast.InfixExpr{
		Token: p.curToken,
		Op: p.curToken.Literal,
		Left: left,
	}

	precidence := p.curPrecedence()
	p.nextToken()
	expr.Right = p.parseExpr(precidence)

	return expr
}

func (p *Parser) parseBoolean() ast.Expr {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseGroupExpr() ast.Expr {
	p.nextToken()
	exp := p.parseExpr(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpr() ast.Expr {
	expr := &ast.IfExpr{Token: p.curToken}
	
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expr.Condition = p.parseExpr(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expr.Consequence = p.parseBlockStmt()

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expr.Alternative = p.parseBlockStmt()
	}

	return expr
}

func (p *Parser) parseFuncLiteral() ast.Expr {
	lit := &ast.FuncLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Params = p.parseFuncParams()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStmt()

	return lit
}

func (p *Parser) parseFuncParams() []*ast.Ident {
	idents := []*ast.Ident{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return idents
	}

	p.nextToken()

	ident := &ast.Ident{Token: p.curToken, Value: p.curToken.Literal}
	idents = append(idents, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
	
		ident := &ast.Ident{Token: p.curToken, Value: p.curToken.Literal}
		idents = append(idents, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return idents
}

func (p *Parser) parseCallExpr(funcc ast.Expr) ast.Expr {
	expr := &ast.CallExpr{Token: p.curToken, Func: funcc}
	expr.Args = p.parseCallArgs()
	return expr
}

func (p *Parser) parseCallArgs() []ast.Expr {
	args := []ast.Expr{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()

	args = append(args, p.parseExpr(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
	
		args = append(args, p.parseExpr(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}