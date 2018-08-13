## khaled          


khaled is an interpreted programming language.           
It's name is derived from hip hop artiste; DJ Khaled.                      

Implemented as I was reading; [Writing An Interpreter In Go - by Thorsten Ball.](https://interpreterbook.com/)   
That book is worth every penny.             

khaled ships with an inbuilt REPL, which you can start by typing;             

`> khaled`
```bash
Bless up komuw! 
	Major Key alert. 
	This is the khaled programming language!
You can type in commands
>> 
```


**Contents:**          
[1. Intro](1.Intro.md)  
[2. Lexer](2.Lexing.md)  


##### TODO:
- [ ] Implement the ideas in this talk: [Lexical Scanning in Go by Rob Pike](https://www.youtube.com/watch?v=HxaD_trXwRE) especially;
  - [ ] using int as TokenType
  - [ ] lexing and parsing concurrently(run lexer in one goroutine and parser in another communicating over a channel)
