package decoder

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/iskorotkov/fastjson/stats"
	"github.com/iskorotkov/fastjson/tokenizer"
	"github.com/iskorotkov/fastjson/xreflect"
	"github.com/iskorotkov/fastjson/xstrconv"
)

var decodersByKind [26]func(typ reflect.Type) Decoder

var decodersByType [4]CustomDecoder

func init() {
	decodersByKind = [...]func(typ reflect.Type) Decoder{
		reflect.Bool:    boolDecoder,
		reflect.Int:     intDecoder,
		reflect.Int8:    intDecoder,
		reflect.Int16:   intDecoder,
		reflect.Int32:   intDecoder,
		reflect.Int64:   intDecoder,
		reflect.Uint:    uintDecoder,
		reflect.Uint8:   uintDecoder,
		reflect.Uint16:  uintDecoder,
		reflect.Uint32:  uintDecoder,
		reflect.Uint64:  uintDecoder,
		reflect.Float32: floatDecoder,
		reflect.Float64: floatDecoder,
		reflect.String:  stringDecoder,
		reflect.Array:   arrayDecoder,
		reflect.Slice:   sliceDecoder,
		reflect.Map:     mapDecoder,
		reflect.Struct:  structDecoder,
		reflect.Pointer: pointerDecoder,
	}

	decodersByType = [...]CustomDecoder{
		{
			Type:    reflect.TypeFor[time.Duration](),
			Decoder: decodeTimeDuration,
		},
		{
			Type:    reflect.TypeFor[json.Unmarshaler](),
			Decoder: decodeJsonUnmarshaler,
		},
		{
			Type:    reflect.TypeFor[encoding.TextUnmarshaler](),
			Decoder: decodeTextUnmarshaler,
		},
		{
			Type:    reflect.TypeFor[encoding.BinaryUnmarshaler](),
			Decoder: decodeBinaryUnmarshaler,
		},
	}
}

func New(typ reflect.Type) Decoder {
	if typ == nil {
		return decodeNil
	}

	for _, dec := range decodersByType {
		if typ.AssignableTo(dec.Type) || reflect.PointerTo(typ).AssignableTo(dec.Type) {
			return dec.Decoder
		}
	}

	kind := typ.Kind()
	if kind >= reflect.Kind(len(decodersByKind)) {
		panic(&UnsupportedTypeError{
			Type: typ,
		})
	}

	f := decodersByKind[kind]
	if f == nil {
		panic(&UnsupportedTypeError{
			Type: typ,
		})
	}

	return f(typ)
}

type Decoder func(value reflect.Value, tokens *tokenizer.Tokenizer)

type CustomDecoder struct {
	Type    reflect.Type
	Decoder Decoder
}

func decodeNil(value reflect.Value, tokens *tokenizer.Tokenizer) {}

func boolDecoder(typ reflect.Type) Decoder {
	return decodeBool
}

func decodeBool(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	switch token.Type {
	case tokenizer.TokenTypeTrue:
		value.SetBool(true)
	case tokenizer.TokenTypeFalse:
		value.SetBool(false)
	default:
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{
				tokenizer.TokenTypeFalse,
				tokenizer.TokenTypeTrue,
			},
			Actual: token,
			Value:  value,
		})
	}
}

func intDecoder(typ reflect.Type) Decoder {
	return decodeInt
}

func decodeInt(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	integer, err := strconv.ParseInt(xstrconv.BytesToString(token.Literal), 10, 64)
	if err != nil {
		panic(&LiteralParseError{
			Err:   err,
			Token: token,
			Value: value,
		})
	}

	value.SetInt(integer)
}

func uintDecoder(typ reflect.Type) Decoder {
	return decodeUint
}

func decodeUint(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	integer, err := strconv.ParseUint(xstrconv.BytesToString(token.Literal), 10, 64)
	if err != nil {
		panic(&LiteralParseError{
			Err:   err,
			Token: token,
			Value: value,
		})
	}

	value.SetUint(integer)
}

func floatDecoder(typ reflect.Type) Decoder {
	return decodeFloat
}

func decodeFloat(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	float, err := strconv.ParseFloat(xstrconv.BytesToString(token.Literal), 64)
	if err != nil {
		panic(&LiteralParseError{
			Err:   err,
			Token: token,
			Value: value,
		})
	}

	value.SetFloat(float)
}

func stringDecoder(typ reflect.Type) Decoder {
	return decodeDecoder
}

func decodeDecoder(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeQuotedLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	value.SetString(xstrconv.BytesToString(token.Unquote()))
}

func arrayDecoder(typ reflect.Type) Decoder {
	elemType := typ.Elem()
	itemsDecoder := New(elemType)
	length := typ.Len()
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		if token.Type != tokenizer.TokenTypeArrayStart {
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeArrayStart},
				Actual:   token,
				Value:    value,
			})
		}

		token = tokens.Peek()
		if token.Type == tokenizer.TokenTypeArrayEnd {
			tokens.Next()
			return
		}

		var index int
		for {
			if index >= length {
				panic(&ArrayLengthError{
					Expected: length,
					Value:    value,
				})
			}

			elemValue := value.Index(index)
			itemsDecoder(elemValue, tokens)

			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeArrayEnd {
				tokens.Next()
				return
			}

			index++
		}
	}
}

func sliceDecoder(typ reflect.Type) Decoder {
	elemType := typ.Elem()
	itemsDecoder := New(elemType)
	var stats stats.BestStat
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		switch token.Type {
		case tokenizer.TokenTypeNull:
			value.Set(reflect.MakeSlice(typ, 0, 0))
		case tokenizer.TokenTypeArrayStart:
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeArrayEnd {
				tokens.Next()
				return
			}

			value.Grow(stats.Get())
			for {
				length := value.Len()
				value.Grow(1)
				value.SetLen(length + 1)

				elemValue := value.Index(length)
				itemsDecoder(elemValue, tokens)

				token = tokens.Peek()
				if token.Type == tokenizer.TokenTypeArrayEnd {
					tokens.Next()
					stats.Add(length + 1)
					return
				}
			}
		default:
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeArrayStart},
				Actual:   token,
				Value:    value,
			})
		}
	}
}

func mapDecoder(typ reflect.Type) Decoder {
	itemsDecoder := New(typ.Elem())
	var stats stats.BestStat
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		switch token.Type {
		case tokenizer.TokenTypeNull:
			value.Set(reflect.MakeMap(typ))
		case tokenizer.TokenTypeObjectStart:
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeObjectEnd {
				tokens.Next()
				return
			}

			value.Set(reflect.MakeMapWithSize(typ, stats.Get()))

			mapKeyValue := reflect.New(typ.Key()).Elem()
			mapElemValue := reflect.New(typ.Elem()).Elem()
			var pairs int
			for {
				token := tokens.Next()
				if token.Type != tokenizer.TokenTypeQuotedLiteral {
					panic(&UnexpectedTokenError{
						Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
						Actual:   token,
						Value:    value,
					})
				}

				key := xstrconv.BytesToString(token.Unquote())
				mapKeyValue.SetString(key)

				itemsDecoder(mapElemValue, tokens)

				value.SetMapIndex(mapKeyValue, mapElemValue)
				pairs++

				token = tokens.Peek()
				if token.Type == tokenizer.TokenTypeObjectEnd {
					tokens.Next()
					stats.Add(pairs)
					return
				}
			}
		default:
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeObjectStart},
				Actual:   token,
				Value:    value,
			})
		}
	}
}

func structDecoder(typ reflect.Type) Decoder {
	var properties Properties
	for i := range typ.NumField() {
		field := typ.Field(i)
		name := xreflect.JSONTag(field)
		dec := New(field.Type)
		properties.Add(Property{Index: i, Name: name, Decoder: dec})
	}
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		if token.Type != tokenizer.TokenTypeObjectStart {
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeObjectStart},
				Actual:   token,
				Value:    value,
			})
		}

		token = tokens.Peek()
		if token.Type == tokenizer.TokenTypeObjectEnd {
			tokens.Next()
			return
		}

		for {
			token := tokens.Next()
			if token.Type != tokenizer.TokenTypeQuotedLiteral {
				panic(&UnexpectedTokenError{
					Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
					Actual:   token,
					Value:    value,
				})
			}

			name := xstrconv.BytesToString(token.Unquote())
			property := properties.Find(name)
			if property.Name == "" {
				panic(&UnknownFieldError{
					Name:  name,
					Value: value,
				})
			}

			valueField := value.Field(property.Index)
			property.Decoder(valueField, tokens)

			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeObjectEnd {
				tokens.Next()
				return
			}
		}
	}
}

func pointerDecoder(typ reflect.Type) Decoder {
	dec := New(typ.Elem())
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Peek()
		if token.Type == tokenizer.TokenTypeNull {
			tokens.Next()
			value.Set(reflect.Zero(typ))
			return
		}

		value.Set(reflect.New(typ.Elem()))
		dec(value.Elem(), tokens)
	}
}

func decodeTimeDuration(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	switch token.Type {
	case tokenizer.TokenTypeLiteral:
		dur, err := strconv.ParseInt(xstrconv.BytesToString(token.Literal), 10, 64)
		if err != nil {
			panic(&LiteralParseError{
				Err:   err,
				Token: token,
				Value: value,
			})
		}

		value.SetInt(int64(dur))
	case tokenizer.TokenTypeQuotedLiteral:
		dur, err := time.ParseDuration(xstrconv.BytesToString(token.Unquote()))
		if err != nil {
			panic(&LiteralParseError{
				Err:   err,
				Token: token,
				Value: value,
			})
		}

		value.SetInt(int64(dur))
	default:
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral, tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	}
}

func decodeJsonUnmarshaler(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeLiteral && token.Type != tokenizer.TokenTypeQuotedLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral, tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	marshaler, ok := xreflect.TypeAssert[json.Unmarshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[json.Unmarshaler](value.Addr())
	}

	if err := marshaler.UnmarshalJSON(token.Literal); err != nil {
		panic(&UnmarshalerError{
			Err:   err,
			Value: value,
		})
	}
}

func decodeTextUnmarshaler(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeQuotedLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	marshaler, ok := xreflect.TypeAssert[encoding.TextUnmarshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[encoding.TextUnmarshaler](value.Addr())
	}

	if err := marshaler.UnmarshalText(token.Unquote()); err != nil {
		panic(&UnmarshalerError{
			Err:   err,
			Value: value,
		})
	}
}

func decodeBinaryUnmarshaler(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	if token.Type != tokenizer.TokenTypeQuotedLiteral {
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	}

	marshaler, ok := xreflect.TypeAssert[encoding.BinaryUnmarshaler](value)
	if !ok {
		marshaler, _ = xreflect.TypeAssert[encoding.BinaryUnmarshaler](value.Addr())
	}

	if err := marshaler.UnmarshalBinary(token.Unquote()); err != nil {
		panic(&UnmarshalerError{
			Err:   err,
			Value: value,
		})
	}
}
