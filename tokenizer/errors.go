package tokenizer

import "strings"

type InvalidTokenError struct {
	Expected TokenType
	Buf      []byte
}

func (e *InvalidTokenError) Error() string {
	var sb strings.Builder
	sb.WriteString("invalid token ")
	sb.Write(e.Buf[:min(20, len(e.Buf))])
	sb.WriteString(", expected ")
	sb.WriteString(e.Expected.String())
	return sb.String()
}

type ReadError struct {
	Err error
}

func (e *ReadError) Error() string {
	return "read error: " + e.Err.Error()
}
