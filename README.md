# Nutlang 🥜

The nuttiest of langs.

I am following along a book for [Writing An Interpreter In Go](https://interpreterbook.com/).

The goal is to tweak it a bit afterwards to my liking :)
(also, I would like to get it far enough to solve a couple of leetcode,
or advent of code, problems comfortably)

## Features

Currently the language supports

- integers
- booleans
- strings
- arrays
- hashes
- prefix-, infix- and index operators
- conditionals
- global and local bindings
- first-class functions
- return statements
- closures

The number of built-in functions are relatively sparse.

- len (strings and arrays)
- push
- first
- last
- rest (copy array, without the first element)

### Things I want to add

This is a list of things I feel are missing for it to be considered a limited,
but still somewhat complete language

#### General features

- [ ] File IO
- [ ] Floats
- [ ] structs/types (some way to group items together)
- [ ] <= and >=
- [ ] while/for loop

#### Arrays

- [ ] Insert
- [ ] Remove
- [ ] Shift
- [ ] Unshift

#### Integers/Floats

- [ ] Random
- [ ] Min
- [ ] Max
- [ ] Modulo

#### Strings

- [ ] split
- [ ] trim
- [ ] startsWith
- [ ] endsWith
- [ ] includes
