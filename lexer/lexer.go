package lexer

import (
	"Nutlang/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) makeTwoCharToken(peekLiteral byte, oneCharTokenType token.TokenType, twoCharTokenType token.TokenType) token.Token {
	if l.peekChar() == peekLiteral {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return token.Token{Type: twoCharTokenType, Literal: literal}
	} else {
		return newToken(oneCharTokenType, l.ch)
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = l.makeTwoCharToken('=', token.BANG, token.NOT_EQ)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken('=', token.LT, token.LTE)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken('=', token.GT, token.GTE)
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '&':
		tok = l.makeTwoCharToken('&', token.BITWISEAND, token.AND)
	case '|':
		tok = l.makeTwoCharToken('|', token.BITWISEOR, token.OR)
	case '=':
		tok = l.makeTwoCharToken('=', token.ASSIGN, token.EQ)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			// Check we stopped on "." -> FLOAT
			if l.ch == '.' {
				// Advance
				l.readChar()
				if isDigit(l.ch) {
					// Concat the rest of the floating point number
					tok.Literal = tok.Literal + "." + l.readNumber()
					tok.Type = token.FLOAT
				} else {
					return newToken(token.ILLEGAL, l.ch)
				}
			} else {
				tok.Type = token.INT
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readString() string {
	/*
		position := l.position + 1
		for {
			l.readChar()
			if l.ch == '"' || l.ch == 0 {
				break
			}
		}
		return l.input[position:l.position]
	*/
	b := strings.Builder{}
	for {
		l.readChar()

		// Support some basic escapes like \"
		if l.ch == '\\' {
			switch l.peekChar() {
			case '"':
				b.WriteByte('"')
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case 't':
				b.WriteByte('\t')
			case '\\':
				b.WriteByte('\\')
			}

			// Skip over the '\\' and the matched single escape char
			l.readChar()
			continue
		} else {
			if l.ch == '"' || l.ch == 0 {
				break
			}
		}

		b.WriteByte(l.ch)
	}

	return b.String()
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
