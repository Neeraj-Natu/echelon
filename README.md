# shifu
<p align="center">
<img width="360" height="360" src="imgs/shifu.jpg">


This is a personal project intended to write my own interpreter in go lang.
</p>

### *Language Featuers*:

- Dynamically typed
- Supports most of the control statements
- Supports strings and integers
- Functions are first class citizens
- Completely written in golang

---

<br/>

### *Language Parts*:

#### *Lexer*:

- Lexer is the lexeical analyser for the language it converts the raw input into Tokens.
- The list of valid tokens is within the the token.go file.
- If the input has anything apart from valid tokens then lexer assigns it an Illegal token.
- This is the first stage of understanding/interpreting the input.

#### *Parser*:

- Parser just like anything else parses the input into a meaningful datastructure.
- Our Parser takes the input from Lexer and converts the tokens into an AST (Abstract Syntax Tree).
- This is the second stage to understanding/interpreting the input.
- The type of Parser we use here is Recursive Descent Parser that works from top down.
- This is also called top down operator precedence parser or a Pratt Parser.
- The parser here won't be fastest or have a formal proof of its correctness and its error recovey process and detection of errorneous syntax won't be always right as it's just the begining for me.

#### *Ast*:

- AST also known as Abstract Syntax Tree is a datastructure that is used to store the langugage tokens to make sense.

#### *Repl*:

- Repl stands for "Read Evaluate Print Loop".
- This loop takes any input and performs the said steps.
- Reads the input, evaluates the input, prints the output and loops again.
- This is a standard in many languages such as Python and Javascript that comes with inbuilt REPL.
- Usually REPL starts with a prompt ">>" thus in shifu the repl starts with the same.
