# MonkeyLang - A programming language written in Go

Learning how to create a programming language by following the book [Writing an interpreter in Go](https://interpreterbook.com/) and building a simple programming language called Monkey.

> 📘 Info
>
> This project is a work in progress and is mostly a) just for fun, b) to learn more about more low-level programming concepts and language design, and c) to improve my Go skills.

Monkeylang is a simple programming language that:

- Has a C-like syntax
- Supports integers, booleans, strings, arrays and hashmaps
- First-class and higher-order functions
- Closures
- Built-in functions
- Arithmetic expressions
- Variable bindings
- Control structures (if/else)
- A REPL
- And more!

Example code (notice the recursive calls and implicit returns):

```monkey
let fibonacci = fn(x) {
    if (x == 0) {
        o
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2)
        }
    }
}
```

Monkey also supports higher order functions:

```monkey
let twice = fn(fnc, int) {
    return fnc(fnc(int))
}

let addTwo = fn(n) {
    return n + 2
}

twice(addTwo, 2) // returns 6
```

Read more about Monkeylang [here](https://monkeylang.org/)
