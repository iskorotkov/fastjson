package tokenizer

import (
	"strings"
)

type Token struct {
	Type    TokenType
	Literal []byte
}

func (t Token) String() string {
	switch t.Type {
	case TokenTypeLiteral, TokenTypeQuotedLiteral:
		var sb strings.Builder
		sb.WriteString(t.Type.String())
		sb.WriteByte('(')
		sb.Write(t.Literal)
		sb.WriteByte(')')
		return sb.String()
	default:
		return t.Type.String()
	}
}

func (t Token) Unquote() []byte {
	if t.Type != TokenTypeQuotedLiteral {
		return t.Literal
	}
	return t.Literal[1 : len(t.Literal)-1]
}
