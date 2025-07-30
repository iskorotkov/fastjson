package decoder

import (
	"encoding"
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/iskorotkov/fastjson/cast"
	"github.com/iskorotkov/fastjson/fields"
	"github.com/iskorotkov/fastjson/stats"
	"github.com/iskorotkov/fastjson/tokenizer"
	"github.com/iskorotkov/fastjson/view"
)

var (
	decodersByKind [26]func(typ reflect.Type) Decoder

	decodersByType [4]CustomDecoder

	decoderCache sync.Map // map[reflect.Type]*cachedDecoder

	decodeInt8  = makeIntDecoder(8)
	decodeInt16 = makeIntDecoder(16)
	decodeInt32 = makeIntDecoder(32)
	decodeInt64 = makeIntDecoder(64)

	decodeUint8  = makeUintDecoder(8)
	decodeUint16 = makeUintDecoder(16)
	decodeUint32 = makeUintDecoder(32)
	decodeUint64 = makeUintDecoder(64)

	decodeFloat32 = makeFloatDecoder(32)
	decodeFloat64 = makeFloatDecoder(64)
)

type cachedDecoder struct {
	dec Decoder
}

func init() {
	decodersByKind = [...]func(typ reflect.Type) Decoder{
		reflect.Bool:    boolDecoder,
		reflect.Int:     intDecoder,
		reflect.Int8:    int8Decoder,
		reflect.Int16:   int16Decoder,
		reflect.Int32:   int32Decoder,
		reflect.Int64:   int64Decoder,
		reflect.Uint:    uintDecoder,
		reflect.Uint8:   uint8Decoder,
		reflect.Uint16:  uint16Decoder,
		reflect.Uint32:  uint32Decoder,
		reflect.Uint64:  uint64Decoder,
		reflect.Uintptr: uintptrDecoder,
		reflect.Float32: float32Decoder,
		reflect.Float64: float64Decoder,
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

	if cached, ok := decoderCache.Load(typ); ok {
		return cached.(*cachedDecoder).dec
	}

	entry := &cachedDecoder{}
	actual, loaded := decoderCache.LoadOrStore(typ, entry)
	if loaded {
		return actual.(*cachedDecoder).dec
	}

	dec := buildDecoder(typ)
	entry.dec = dec
	return dec
}

func buildDecoder(typ reflect.Type) Decoder {
	for _, dec := range decodersByType {
		if typ.AssignableTo(dec.Type) || reflect.PointerTo(typ).AssignableTo(dec.Type) {
			return dec.Decoder
		}
	}

	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Uint8 {
		return decodeByteSlice
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
	if strconv.IntSize == 32 {
		return decodeInt32
	}
	return decodeInt64
}

func int8Decoder(typ reflect.Type) Decoder  { return decodeInt8 }
func int16Decoder(typ reflect.Type) Decoder { return decodeInt16 }
func int32Decoder(typ reflect.Type) Decoder { return decodeInt32 }
func int64Decoder(typ reflect.Type) Decoder { return decodeInt64 }

func makeIntDecoder(bitSize int) Decoder {
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		if token.Type != tokenizer.TokenTypeLiteral {
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
				Actual:   token,
				Value:    value,
			})
		}
		integer, err := strconv.ParseInt(view.BytesAsStr(token.Literal), 10, bitSize)
		if err != nil {
			panic(&LiteralParseError{Err: err, Token: token, Value: value})
		}
		value.SetInt(integer)
	}
}

func uintDecoder(typ reflect.Type) Decoder {
	if strconv.IntSize == 32 {
		return decodeUint32
	}
	return decodeUint64
}

func uint8Decoder(typ reflect.Type) Decoder  { return decodeUint8 }
func uint16Decoder(typ reflect.Type) Decoder { return decodeUint16 }
func uint32Decoder(typ reflect.Type) Decoder { return decodeUint32 }
func uint64Decoder(typ reflect.Type) Decoder { return decodeUint64 }

func uintptrDecoder(typ reflect.Type) Decoder {
	if strconv.IntSize == 32 {
		return decodeUint32
	}
	return decodeUint64
}

func makeUintDecoder(bitSize int) Decoder {
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		if token.Type != tokenizer.TokenTypeLiteral {
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
				Actual:   token,
				Value:    value,
			})
		}
		integer, err := strconv.ParseUint(view.BytesAsStr(token.Literal), 10, bitSize)
		if err != nil {
			panic(&LiteralParseError{Err: err, Token: token, Value: value})
		}
		value.SetUint(integer)
	}
}

func float32Decoder(typ reflect.Type) Decoder { return decodeFloat32 }
func float64Decoder(typ reflect.Type) Decoder { return decodeFloat64 }

func makeFloatDecoder(bitSize int) Decoder {
	return func(value reflect.Value, tokens *tokenizer.Tokenizer) {
		token := tokens.Next()
		if token.Type != tokenizer.TokenTypeLiteral {
			panic(&UnexpectedTokenError{
				Expected: []tokenizer.TokenType{tokenizer.TokenTypeLiteral},
				Actual:   token,
				Value:    value,
			})
		}
		f, err := strconv.ParseFloat(view.BytesAsStr(token.Literal), bitSize)
		if err != nil {
			panic(&LiteralParseError{Err: err, Token: token, Value: value})
		}
		value.SetFloat(f)
	}
}

func stringDecoder(typ reflect.Type) Decoder {
	return decodeString
}

func parseQuotedString(token tokenizer.Token, value reflect.Value) string {
	switch token.Type {
	case tokenizer.TokenTypeQuotedLiteral:
		return string(token.Unquote())
	case tokenizer.TokenTypeQuotedEscapedLiteral:
		str, err := unescapeJSON(token.Unquote())
		if err != nil {
			panic(&StringEscapeError{Err: err, Token: token, Value: value})
		}
		return str
	default:
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral, tokenizer.TokenTypeQuotedEscapedLiteral},
			Actual:   token,
			Value:    value,
		})
	}
}

func decodeString(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	value.SetString(parseQuotedString(token, value))
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
			if length > 0 {
				panic(&ArrayLengthError{
					Expected: length,
					Value:    value,
				})
			}
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
				if index+1 < length {
					panic(&ArrayLengthError{
						Expected: length,
						Value:    value,
					})
				}
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
			value.SetZero()
		case tokenizer.TokenTypeArrayStart:
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeArrayEnd {
				tokens.Next()
				value.Set(reflect.MakeSlice(typ, 0, 0))
				return
			}

			value.Grow(stats.Get())
			capacity := value.Cap()

			var length int
			for {
				if length >= capacity {
					value.Grow(max(capacity, 1))
					capacity = value.Cap()
				}
				value.SetLen(length + 1)

				elemValue := value.Index(length)
				itemsDecoder(elemValue, tokens)
				length++

				token = tokens.Peek()
				if token.Type == tokenizer.TokenTypeArrayEnd {
					tokens.Next()
					stats.Add(length)
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
			value.SetZero()
		case tokenizer.TokenTypeObjectStart:
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeObjectEnd {
				tokens.Next()
				value.Set(reflect.MakeMap(typ))
				return
			}

			value.Set(reflect.MakeMapWithSize(typ, stats.Get()))

			mapKeyValue := reflect.New(typ.Key()).Elem()
			mapElemValue := reflect.New(typ.Elem()).Elem()
			var pairs int
			for {
				token := tokens.Next()
				mapKeyValue.SetString(parseQuotedString(token, value))

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
	for f := range fields.Extract(typ) {
		properties.Add(Property{Index: f.Index, Name: f.Tag.Name, Decoder: New(f.Type)})
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
			name := parseQuotedString(token, value)

			property := properties.Find(name)
			if property.Name == "" {
				skipValue(tokens)
				token = tokens.Peek()
				if token.Type == tokenizer.TokenTypeObjectEnd {
					tokens.Next()
					return
				}
				continue
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

func skipValue(tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	switch token.Type {
	case tokenizer.TokenTypeObjectStart:
		for {
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeObjectEnd {
				tokens.Next()
				return
			}
			tokens.Next()
			skipValue(tokens)
		}
	case tokenizer.TokenTypeArrayStart:
		for {
			token = tokens.Peek()
			if token.Type == tokenizer.TokenTypeArrayEnd {
				tokens.Next()
				return
			}
			skipValue(tokens)
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
		dur, err := strconv.ParseInt(view.BytesAsStr(token.Literal), 10, 64)
		if err != nil {
			panic(&LiteralParseError{
				Err:   err,
				Token: token,
				Value: value,
			})
		}

		value.SetInt(int64(dur))
	case tokenizer.TokenTypeQuotedLiteral:
		dur, err := time.ParseDuration(view.BytesAsStr(token.Unquote()))
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

	marshaler, _ := cast.To[json.Unmarshaler](value)

	if err := marshaler.UnmarshalJSON(token.Literal); err != nil {
		panic(&UnmarshalerError{
			Err:   err,
			Value: value,
		})
	}
}

func decodeTextUnmarshaler(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	switch token.Type {
	case tokenizer.TokenTypeNull:
		if value.Kind() == reflect.Slice {
			value.SetZero()
			return
		}
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
		})
	case tokenizer.TokenTypeQuotedLiteral:
		marshaler, _ := cast.To[encoding.TextUnmarshaler](value)

		if err := marshaler.UnmarshalText(token.Unquote()); err != nil {
			panic(&UnmarshalerError{
				Err:   err,
				Value: value,
			})
		}
	default:
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral},
			Actual:   token,
			Value:    value,
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

	marshaler, _ := cast.To[encoding.BinaryUnmarshaler](value)

	if err := marshaler.UnmarshalBinary(token.Unquote()); err != nil {
		panic(&UnmarshalerError{
			Err:   err,
			Value: value,
		})
	}
}

func decodeByteSlice(value reflect.Value, tokens *tokenizer.Tokenizer) {
	token := tokens.Next()
	switch token.Type {
	case tokenizer.TokenTypeNull:
		value.SetZero()
	case tokenizer.TokenTypeQuotedLiteral:
		if len(token.Literal) == 2 {
			value.SetBytes([]byte{})
			return
		}
		decoded, err := base64.StdEncoding.AppendDecode(nil, token.Unquote())
		if err != nil {
			panic(&Base64DecodeError{Err: err, Value: value})
		}
		value.SetBytes(decoded)
	default:
		panic(&UnexpectedTokenError{
			Expected: []tokenizer.TokenType{tokenizer.TokenTypeQuotedLiteral, tokenizer.TokenTypeNull},
			Actual:   token,
			Value:    value,
		})
	}
}
