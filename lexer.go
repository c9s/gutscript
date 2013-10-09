package main

import (
	"fmt"
	"unicode"
)

type CoffeeLex struct {
	// the line
	input string
	start int
	pos   int
	items chan LexItem
}

type TokenType int

const (
	TokenDot = iota
	TokenEOF
	TokenIf
	TokenElse
	TokenElseIf
	TokenDigit
	TokenIdentifier
	TokenForeach
)

type LexItem struct {
	Val string
	Typ TokenType
}

func (self *LexItem) String() string {
	switch self.Typ {
	case TokenEOF:
		return "EOF"
	}
	return fmt.Sprintf("%q", self.Val)
}

/*
func (l *CoffeeLex) run() {
	for state := l.Lex(); state != nil {
		state = state(l)
	}
	close(l.items)
}
*/

// returns a token
func (l *CoffeeLex) Lex(lval *CoffeeSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.input) {
			return 0
		}
		c = rune(l.input[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c - '0')
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c - 'a')
		return LETTER
	}
	return int(c)
}

func (l *CoffeeLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
