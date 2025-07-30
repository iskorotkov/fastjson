use std::io::{self, Read};

#[derive(Debug, Clone, Copy, PartialEq)]
pub enum TokenType {
    Null = 0,
    True,
    False,
    Literal,
    ObjectStart,
    ObjectEnd,
    ArrayStart,
    ArrayEnd,
    EOF,
}

#[derive(Debug, Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub literal: Vec<u8>,
}

impl Token {
    pub fn new(token_type: TokenType) -> Self {
        Self {
            token_type,
            literal: Vec::new(),
        }
    }

    pub fn new_with_literal(token_type: TokenType, literal: Vec<u8>) -> Self {
        Self {
            token_type,
            literal,
        }
    }

    pub fn unquote(&self) -> &[u8] {
        if self.literal.len() < 2
            || self.literal[0] != b'"'
            || self.literal[self.literal.len() - 1] != b'"'
        {
            &self.literal
        } else {
            &self.literal[1..self.literal.len() - 1]
        }
    }
}

impl std::fmt::Display for Token {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self.token_type {
            TokenType::Null => write!(f, "null"),
            TokenType::True => write!(f, "true"),
            TokenType::False => write!(f, "false"),
            TokenType::Literal => write!(f, "literal({})", String::from_utf8_lossy(&self.literal)),
            TokenType::ObjectStart => write!(f, "{{"),
            TokenType::ObjectEnd => write!(f, "}}"),
            TokenType::ArrayStart => write!(f, "["),
            TokenType::ArrayEnd => write!(f, "]"),
            TokenType::EOF => write!(f, "eof"),
        }
    }
}

static NUMBER_LITERAL_BYTES: [u8; 256] = {
    let mut arr = [0u8; 256];
    arr[b'0' as usize] = 1;
    arr[b'1' as usize] = 1;
    arr[b'2' as usize] = 1;
    arr[b'3' as usize] = 1;
    arr[b'4' as usize] = 1;
    arr[b'5' as usize] = 1;
    arr[b'6' as usize] = 1;
    arr[b'7' as usize] = 1;
    arr[b'8' as usize] = 1;
    arr[b'9' as usize] = 1;
    arr[b'-' as usize] = 1;
    arr[b'+' as usize] = 1;
    arr[b'.' as usize] = 1;
    arr[b'e' as usize] = 1;
    arr[b'E' as usize] = 1;
    arr
};

static SKIP_BYTES: [u8; 256] = {
    let mut arr = [0u8; 256];
    arr[b' ' as usize] = 1;
    arr[b'\n' as usize] = 1;
    arr[b'\r' as usize] = 1;
    arr[b'\t' as usize] = 1;
    arr[b',' as usize] = 1;
    arr[b':' as usize] = 1;
    arr
};

pub struct Tokenizer {
    buf: Vec<u8>,
    has_peeked: bool,
    peeked_token: Token,
}

impl Tokenizer {
    pub fn new_from_bytes(buf: Vec<u8>) -> Self {
        Self {
            buf,
            has_peeked: false,
            peeked_token: Token::new(TokenType::EOF),
        }
    }

    pub fn new_from_string(s: String) -> Self {
        Self::new_from_bytes(s.into_bytes())
    }

    pub fn new_from_reader<R: Read>(mut reader: R) -> io::Result<Self> {
        let mut buf = Vec::new();
        reader.read_to_end(&mut buf)?;
        Ok(Self::new_from_bytes(buf))
    }

    pub fn all(&mut self) -> Vec<Token> {
        let mut tokens = Vec::new();
        loop {
            let token = self.next();
            if token.token_type == TokenType::EOF {
                return tokens;
            }
            tokens.push(token);
        }
    }

    pub fn peek(&mut self) -> Token {
        if self.has_peeked {
            return self.peeked_token.clone();
        }

        self.skip_whitespace();

        if self.buf.is_empty() {
            return Token::new(TokenType::EOF);
        }

        let (skip, token) = self.parse_token().unwrap_or_else(|| {
            panic!(
                "unexpected '{}'",
                String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
            )
        });

        self.buf.drain(..skip);
        self.has_peeked = true;
        self.peeked_token = token.clone();

        token
    }

    pub fn next(&mut self) -> Token {
        if self.has_peeked {
            self.has_peeked = false;
            return self.peeked_token.clone();
        }

        self.skip_whitespace();

        if self.buf.is_empty() {
            return Token::new(TokenType::EOF);
        }

        let (skip, token) = self.parse_token().unwrap_or_else(|| {
            panic!(
                "unexpected '{}'",
                String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
            )
        });

        self.buf.drain(..skip);
        token
    }

    fn skip_whitespace(&mut self) {
        while !self.buf.is_empty() && SKIP_BYTES[self.buf[0] as usize] == 1 {
            self.buf.remove(0);
        }
    }

    fn parse_token(&self) -> Option<(usize, Token)> {
        if self.buf.is_empty() {
            return None;
        }

        match self.buf[0] {
            b'{' => Some((1, Token::new(TokenType::ObjectStart))),
            b'}' => Some((1, Token::new(TokenType::ObjectEnd))),
            b'[' => Some((1, Token::new(TokenType::ArrayStart))),
            b']' => Some((1, Token::new(TokenType::ArrayEnd))),
            b'n' => {
                if !is_null(&self.buf) {
                    panic!(
                        "expected 'null', got: '{}'",
                        String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
                    );
                }
                Some((4, Token::new(TokenType::Null)))
            }
            b't' => {
                if !is_true(&self.buf) {
                    panic!(
                        "expected 'true', got: '{}'",
                        String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
                    );
                }
                Some((4, Token::new(TokenType::True)))
            }
            b'f' => {
                if !is_false(&self.buf) {
                    panic!(
                        "expected 'false', got: '{}'",
                        String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
                    );
                }
                Some((5, Token::new(TokenType::False)))
            }
            b'"' => {
                let (literal, ok) = string_literal(&self.buf);
                if !ok {
                    panic!(
                        "unclosed string literal: '{}'",
                        String::from_utf8_lossy(&self.buf[..std::cmp::min(20, self.buf.len())])
                    );
                }
                Some((
                    literal.len(),
                    Token::new_with_literal(TokenType::Literal, literal),
                ))
            }
            b'0'..=b'9' | b'-' | b'+' => {
                let literal = number_literal(&self.buf);
                Some((
                    literal.len(),
                    Token::new_with_literal(TokenType::Literal, literal),
                ))
            }
            _ => None,
        }
    }
}

fn number_literal(src: &[u8]) -> Vec<u8> {
    for (i, &byte) in src.iter().enumerate() {
        if NUMBER_LITERAL_BYTES[byte as usize] == 0 {
            return src[..i].to_vec();
        }
    }
    src.to_vec()
}

fn string_literal(src: &[u8]) -> (Vec<u8>, bool) {
    if src.is_empty() {
        return (Vec::new(), false);
    }

    let mut escaped = true;
    for (i, &byte) in src.iter().enumerate() {
        match byte {
            b'"' => {
                if !escaped {
                    return (src[..=i].to_vec(), true);
                }
                escaped = false;
            }
            b'\\' => {
                escaped = !escaped;
            }
            _ => {
                escaped = false;
            }
        }
    }

    (Vec::new(), false)
}

fn is_null(b: &[u8]) -> bool {
    b.len() >= 4 && b[0] == b'n' && b[1] == b'u' && b[2] == b'l' && b[3] == b'l'
}

fn is_true(b: &[u8]) -> bool {
    b.len() >= 4 && b[0] == b't' && b[1] == b'r' && b[2] == b'u' && b[3] == b'e'
}

fn is_false(b: &[u8]) -> bool {
    b.len() >= 5 && b[0] == b'f' && b[1] == b'a' && b[2] == b'l' && b[3] == b's' && b[4] == b'e'
}
