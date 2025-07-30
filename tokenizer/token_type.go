package tokenizer

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=TokenType -linecomment -output token_type_string.go

const (
	TokenTypeNull          TokenType = iota // null
	TokenTypeTrue                           // true
	TokenTypeFalse                          // false
	TokenTypeLiteral                        // literal
	TokenTypeQuotedLiteral                  // quoted_literal
	TokenTypeObjectStart                    // {
	TokenTypeObjectEnd                      // }
	TokenTypeArrayStart                     // [
	TokenTypeArrayEnd                       // ]
	TokenTypeEOF                            // eof
)

type TokenType uint8
