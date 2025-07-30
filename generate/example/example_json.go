package example

// var linkSchema = objectSchema{
// 	fields: map[string]fieldSchema{
// 		"title": stringSchema{},
// 		"url":   stringSchema{},
// 	},
// }

// func unmarshalLink(tokens tokenizer.Tokenizer) (Link, error) {
// 	tokens, token, err := tokens.Next()
// 	if err != nil {
// 		return Link{}, fmt.Errorf("next token: %w", err)
// 	}

// 	if token.Type != tokenizer.TokenTypeObjectStart {
// 		return Link{}, fmt.Errorf("expected object start, got %s", token.Type)
// 	}

// 	var link Link
// 	for {
// 		tokens, token, err = tokens.Next()
// 		if err != nil {
// 			return Link{}, fmt.Errorf("next token: %w", err)
// 		}
// 	}

// 	tokens, token, err = tokens.Next()
// 	if err != nil {
// 		return Link{}, fmt.Errorf("next token: %w", err)
// 	}

// 	if token.Type != tokenizer.TokenTypeObjectEnd {
// 		return Link{}, fmt.Errorf("expected object end, got %s", token.Type)
// 	}

// 	return link, nil
// }
