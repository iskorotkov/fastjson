package decoder

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/iskorotkov/fastjson/tokenizer"
)

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return "unsupported type " + e.Type.String()
}

type UnexpectedTokenError struct {
	Expected []tokenizer.TokenType
	Actual   tokenizer.Token
	Value    reflect.Value
}

func (e *UnexpectedTokenError) Error() string {
	var sb strings.Builder
	sb.WriteString("unexpected token ")
	sb.WriteString(e.Actual.String())
	sb.WriteString(", expected one of ")
	for i, expected := range e.Expected {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(expected.String())
	}
	return sb.String()
}

type LiteralParseError struct {
	Err   error
	Token tokenizer.Token
	Value reflect.Value
}

func (e *LiteralParseError) Error() string {
	var sb strings.Builder
	sb.WriteString("can't parse literal ")
	sb.WriteString(e.Token.String())
	sb.WriteString(" for value ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(": ")
	sb.WriteString(e.Err.Error())
	return sb.String()
}

type UnmarshalerError struct {
	Err   error
	Value reflect.Value
}

func (e *UnmarshalerError) Error() string {
	var sb strings.Builder
	sb.WriteString("can't unmarshal ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(": ")
	sb.WriteString(e.Err.Error())
	return sb.String()
}

type ArrayLengthError struct {
	Expected int
	Value    reflect.Value
}

func (e *ArrayLengthError) Error() string {
	var sb strings.Builder
	sb.WriteString("expected array length ")
	sb.WriteString(strconv.Itoa(e.Expected))
	sb.WriteString(" for value ")
	sb.WriteString(e.Value.Type().String())
	return sb.String()
}

type Base64DecodeError struct {
	Err   error
	Value reflect.Value
}

func (e *Base64DecodeError) Error() string {
	var sb strings.Builder
	sb.WriteString("can't decode base64 for value ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(": ")
	sb.WriteString(e.Err.Error())
	return sb.String()
}

type StringEscapeError struct {
	Err   error
	Token tokenizer.Token
	Value reflect.Value
}

func (e *StringEscapeError) Error() string {
	var sb strings.Builder
	sb.WriteString("can't unescape string ")
	sb.WriteString(e.Token.String())
	sb.WriteString(" for value ")
	sb.WriteString(e.Value.Type().String())
	sb.WriteString(": ")
	sb.WriteString(e.Err.Error())
	return sb.String()
}

type EscapeError struct {
	Message string
}

func (e *EscapeError) Error() string {
	return e.Message
}
