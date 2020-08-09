# shifu
![build-test](https://github.com/Neeraj-Natu/shifu/workflows/build-test/badge.svg)

<p align="center">
<img width="360" height="360" src="imgs/shifu.jpg">
</p>

<div align="center">
This is a personal project intended to write my own interpreter in go lang.
</div>

<br/>

---

### *Getting Started*:

```
go run main.go
```

### *Language Featuers*:

- Dynamically typed
- Supports higher order functions
- Supports closures
- Supports integer arithmetic
- Supports strings, integers, arrays and maps
- Supports builtin functions
- Functions are first class citizens
- Completely written in golang

---

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
- The main idea behind pratt parser is to associate parsing functions with token types, whenever a token type is encountered the appropriate parsing function is called which returns an AST node that represents the expression.
- Each token type can have upto two parsing functions associated with it, depending on whether the otken is found in a prefix position or infix position.
- The parser here won't be fastest or have a formal proof of its correctness and its error recovey process and detection of errorneous syntax won't be always right as it's just the begining for me.
- Supports prefix and infix operators. work for supporting postfix operators in progress.
- Supports let statements, return statements and expressions.

#### *Evaluator*:

- Evaluator/Interpreter is the part of the language that takes the parsed code from the parser as an input and then executes it.
- This gives the meaning to the language and makes it come to life.
- This is the final stage to understand/interpret the input.
- There are several different ways to build an Interpreter, some of them listed in terms of increasing complexity and performance:
  - Tree Walking interpreter that interprets the AST on the fly. (The one implemented here.)
  - Compiling the AST to an Intermediate byteCode that is compact and then use a virtual machine (something like JVM) to interpret this byteCode.
  - Convert the byteCode compiled in the above step to highly optimized machine code right before interpreting that machine code. This is also called JIT or Just In Time Interpreter. Although here would need to support different machine architectures.
  - Others skip the conversion to byteCode and directly convert the AST to machine code and then interpret it.

#### *Ast*:

- AST also known as Abstract Syntax Tree is a datastructure that is used to store the langugage tokens to make sense.

#### *Repl*:

- Repl stands for "Read Evaluate Print Loop".
- This loop takes any input and performs the said steps.
- Reads the input, evaluates the input, prints the output and loops again.
- This is a standard in many languages such as Python and Javascript that comes with inbuilt REPL.
- Usually REPL starts with a prompt ">>" thus in shifu the repl starts with the same.

---
Credits:

For this project I have been following through a great book written by [Thorsten Ball](https://thorstenball.com/). The Book name is : [Writing an Interpreter in GO](https://interpreterbook.com/), Also have been adding along a few changes as per my understanding.