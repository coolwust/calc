package main

import (
	"text/scanner"
	"io"
	"strings"
	"fmt"
)

func parse(s string) expr {
	l := new(lex)
	l.s = new(scanner.Scanner)
	l.s.Init(strings.NewReader(s))
	l.s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanInts
	l.r = s.Scan()
	if l.r == scanner.EOF {
		return nil
	}
	return l.additive()
}

type lex struct {
	s *scanner.Scanner
	r rune
}

func (l *lex) additive() expr {
	x := l.multiplicative()
	if l.r == scanner.EOF {
		return &additive{x: x}
	}
	if l.r != '+' && l.r != '-' {
		panic(fmt.Sprintf("parse: %s: expect + or -, get %c", l.s.Pos(), l.r))
	}
	op := string(l.r)
	l.r = l.s.Scan()
	return &additive{op: op, x: x, y: l.additive()}
}

func (l *lex) multiplicative() expr {
	x := l.unary()
	if l.r = l.s.Scan(); l.r == scanner.EOF || (l.r != '*' && l.r != '/' && l.r != '%') {
		return &multiplicative{x: x}
	}
	op := string(l.r)
	l.r = l.s.Scan()
	return &multiplicative{op: op, x: x, y: l.multiplicative()}
}

func (l *lex) unary() expr {
	if l.r < 0 {
		return unary{x: l.fn()}
	}
	if l.r != '+' && l.r != '-' {
		panic(fmt.Sprintf("parse: %s: unexpected %c", l.s.Pos(), l.r))
	}
	op := string(l.r)
	l.r = l.s.Scan()
	return unary{op: op, x: l.unary()}
}

func (l *lex) fn() expr {
	if l.r != scanner.Indent || l.s.Peek != '(' {
		return &
	}
}
