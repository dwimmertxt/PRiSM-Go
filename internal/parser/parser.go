package parser

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type lexer struct {
	tokens []string
	index  int
}

func (l *lexer) init(s string) {
	textScanner := &scanner.Scanner{}
	textScanner.Init(strings.NewReader(s))
	for token := textScanner.Scan(); token != scanner.EOF; token = textScanner.Scan() {
		tokenText := textScanner.TokenText()
		l.tokens = append(l.tokens, tokenText)
	}
}

func (l *lexer) next() string {
	if l.index == len(l.tokens) {
		return ""
	}
	num := l.tokens[l.index]
	l.index++
	return num
}

func (l *lexer) peek() string {
	if l.index == len(l.tokens) {
		return ""
	}
	return l.tokens[l.index]
}

var lex *lexer

type rule struct {
	// prefix operator
	nud func(token string) int
	// infix operator
	led func(left int, token string) int
	// operator precedence (or binding power in Pratt's terminology)
	prec int
}

func binaryLed(left int, token string) int {
	// left-associative
	right := parse(getRule(token).prec + 1)
	switch string(token) {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return left
}

func unary(token string) int {
	// right-associative
	right := parse(30)
	switch token {
	case "-":
		return -right
	}
	return 0
}

func grouping(token string) int {
	// right-associative
	val := parse(0)
	lex.next() // consume ")"
	return val
}

func number(token string) int {
	num, _ := strconv.ParseInt(token, 10, 0)
	return int(num)
}

func getRule(token string) rule {
	if _, err := strconv.ParseInt(token, 10, 64); err == nil {
		return rule{number, nil, -1}
	}
	// EOF
	if string(token) == "" {
		return rule{nil, nil, -1}
	}
	// precedence table
	table := map[string]rule{
		"+": {nil, binaryLed, 10},
		"-": {unary, binaryLed, 10},
		"*": {nil, binaryLed, 20},
		"/": {nil, binaryLed, 20},
		"(": {grouping, nil, 40},
		")": {nil, nil, -1},
	}
	return table[token]
}

func Calculate(s string) int {
	lex = &lexer{}
	lex.init(s)
	return parse(0)
}

/*
	pratt-parser core algorithm:

	1): You can assign each operator token a precedence,
		or binding power in Pratt's terminology.
	2): You have a recursive function that parses expressions, consuming tokens to the right,
		until it reaches an operator of precedence less than or equal to the previous operator -- or just less than if it's a right-associative operator.
	3): In Pratt parsing, tokens can be used in the null and/or left position,
		based on whether they take an expression on the left or not (nud or led in Pratt's terminology).
		Examples of left operators are infix +, postfix ++, or the pseudo-infix a[0] (with operator [).
		Examples of null operators are unary minus -, and grouping parentheses (.
*/
func parse(prec int) int {
	token := lex.next()
	if getRule(token).nud == nil {
		fmt.Println("error with non-correct token")
		return 0
	}

	left := getRule(token).nud(string(token))
	for prec <= getRule(lex.peek()).prec {
		token = lex.next()

		if getRule(token).led != nil {
			left = getRule(token).led(left, token)
		} else {
			fmt.Println("LED not defined for token type", token)
		}
	}

	return left
}