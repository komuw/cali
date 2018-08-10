# LEXING

### 1. Lexical analysis
```bash
                   lexing/tokenization/scanning                  parsing
                   (lexer/tokenizer/scannner)                    (parser)
Source-code       ------------------------------>   Tokens   ---------------> AST 
"let x = 5 + 5;"                                    [
                                                    LET,
                                                    IDENTIFIER("x"),
                                                    EQUAL_SIGN,
                                                    INTEGER(5),
                                                    PLUS_SIGN,
                                                    INTEGER(5),
                                                    SEMICOLON
                                                    ]
```


