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
what exactly constitutes a "token" varies between different lexer implementations.
A production-ready lexer might also attach the line number, column number and filename to
a token. A reason may be so that we can output useful error messages in the parsing stage:
```
"error: expected semicolon token. line 42, column 23, program.khaled"
```

### 2. Define tokens
```bash
let five = 5; #1. numbers like 5
let add = fn(x, y) { #2. keywords like fn, let
x + y; #3. vars like x, y, add, result
};
let result = add(five, ten); #4. special chars like (, ), {, }, ;
```
