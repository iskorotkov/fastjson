package tokenizer_test

import (
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/iskorotkov/fastjson/tokenizer"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

func TestTokenizerNext(t *testing.T) {
	cases := []struct {
		name   string
		json   string
		tokens []tokenizer.Token
	}{
		{
			name:   "null",
			json:   "null",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeNull}},
		},
		{
			name:   "true",
			json:   "true",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeTrue}},
		},
		{
			name:   "false",
			json:   "false",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeFalse}},
		},
		{
			name:   "int",
			json:   "42",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeLiteral, Literal: []byte("42")}},
		},
		{
			name:   "float",
			json:   "42.13",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeLiteral, Literal: []byte("42.13")}},
		},
		{
			name:   "string",
			json:   `"hello"`,
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"hello"`)}},
		},
		{
			name:   "escaped string",
			json:   `"hello \"escaped\""`,
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"hello \"escaped\""`)}},
		},
		{
			name: "array",
			json: "[2, 3, 5, 8, 13]",
			tokens: []tokenizer.Token{
				{Type: tokenizer.TokenTypeArrayStart},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("2")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("3")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("5")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("8")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("13")},
				{Type: tokenizer.TokenTypeArrayEnd},
			},
		},
		{
			name: "object",
			json: `{"key": "value"}`,
			tokens: []tokenizer.Token{
				{Type: tokenizer.TokenTypeObjectStart},
				{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"key"`)},
				{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"value"`)},
				{Type: tokenizer.TokenTypeObjectEnd},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tok := tokenizer.NewFromString(c.json)
			var tokens []tokenizer.Token
			for {
				token := tok.Next()
				if token.Type == tokenizer.TokenTypeEOF {
					break
				}
				tokens = append(tokens, token)
			}
			if !reflect.DeepEqual(tokens, c.tokens) {
				t.Fatalf("expected %v, got %v", c.tokens, tokens)
			}
		})
	}
}

func TestTokenizerAll(t *testing.T) {
	cases := []struct {
		name   string
		json   string
		tokens []tokenizer.Token
	}{
		{
			name:   "null",
			json:   "null",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeNull}},
		},
		{
			name:   "true",
			json:   "true",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeTrue}},
		},
		{
			name:   "false",
			json:   "false",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeFalse}},
		},
		{
			name:   "int",
			json:   "42",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeLiteral, Literal: []byte("42")}},
		},
		{
			name:   "float",
			json:   "42.13",
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeLiteral, Literal: []byte("42.13")}},
		},
		{
			name:   "string",
			json:   `"hello"`,
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"hello"`)}},
		},
		{
			name:   "escaped string",
			json:   `"hello \"escaped\""`,
			tokens: []tokenizer.Token{{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"hello \"escaped\""`)}},
		},
		{
			name: "array",
			json: "[2, 3, 5, 8, 13]",
			tokens: []tokenizer.Token{
				{Type: tokenizer.TokenTypeArrayStart},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("2")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("3")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("5")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("8")},
				{Type: tokenizer.TokenTypeLiteral, Literal: []byte("13")},
				{Type: tokenizer.TokenTypeArrayEnd},
			},
		},
		{
			name: "object",
			json: `{"key": "value"}`,
			tokens: []tokenizer.Token{
				{Type: tokenizer.TokenTypeObjectStart},
				{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"key"`)},
				{Type: tokenizer.TokenTypeQuotedLiteral, Literal: []byte(`"value"`)},
				{Type: tokenizer.TokenTypeObjectEnd},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tokenizer := tokenizer.NewFromString(c.json)
			tokens := tokenizer.All()
			if !reflect.DeepEqual(tokens, c.tokens) {
				t.Fatalf("expected %v, got %v", c.tokens, tokens)
			}
		})
	}
}

func BenchmarkTokenizer(b *testing.B) {
	value := `{"key": "value", "array": [1, 2, 3]}`

	b.Run("next", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(value)))

		for b.Loop() {
			tokens := tokenizer.NewFromString(value)
			for {
				token := tokens.Next()
				if token.Type == tokenizer.TokenTypeEOF {
					break
				}
			}
		}
	})

	b.Run("all", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(value)))

		for b.Loop() {
			tok := tokenizer.NewFromString(value)
			tokens := tok.All()
			_ = tokens
		}
	})
}
