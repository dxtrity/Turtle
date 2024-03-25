# 🐢 Turtle
Turtle is a tiny little language. It has a lexer, (semi-working) parser and an 'evaluator' for parsing mathematical expressions.

## Basics
Turtle has little to no functionality. It's a barebones interpreter for maths. That's it.
It does support: addition, subtraction, multiplication and division.
It does also support variables. Here's how it looks like:

```
1 + 1           // addition
10 - 5          // subtraction
10 / 2          // division
2 * 6           // multiplication

a = 5           // declare variables
b = 10          // can be any word letter etc.

5 + a           // use variable

5 + a - b + 6   // complex expression
```

## Language Implementation
- [x] Addition
- [x] Subtraction
- [x] Multiplication
- [x] Division
- [x] Basic Variables
- [ ] Parentheses

## Known Issues
There is a lot of issue with the interpreter at the moment. Such as:

Variable Mutability
 : This isn't implemented

Variable Expressions
 : Can't use mathematical expressions on just variables `x + b` returns an error

Any type of Parentheses Operations
 : They just don't work

## Installation
To install and build this you will need to install the latest version of **Go**.
The repo has a build script for **Windows PowerShell** as it is my primary shell.
I will add a build script for Linux and Mac in the future if I do anything with this.

```ps1
git clone reponame .
```

```./build.ps1```

Create a `test.tl` file
Write some expressions

```./test.ps1```

And you should be fine. I hope.