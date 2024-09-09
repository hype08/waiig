package parser

import (
	"github.com/hype08/waiig/01/src/monkey/lexer"
	"github.com/hype08/waiig/01/src/monkey/token"
	"github.com/hype08/waiig/02/src/monkey/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
