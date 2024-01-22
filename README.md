# Nutlang 🥜

The nuttiest of langs.

Base code from [Writing An Interpreter In Go](https://interpreterbook.com/).
You might recognize Monkey Language or similar names,
but I wanted to give it my own name.

The goal is to improve the language far enough to solve a couple of leetcode
or advent of code problems, comfortably.

## Features

After following the book I ended up with these features:

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
- puts (print to STDOUT)

### Things I want to add

This is a list of things I feel are missing for it to be considered a limited,
but still somewhat complete language

#### General features

- [x] Value Assignment (meaning re-assignment, after let statement)
  - [x] Arrays
  - [x] Hashes
  - [x] Integers
- [ ] File IO
  - [ ] Open files
  - [ ] Close files
  - [ ] Read file content
  - [ ] Write to file
- [x] Floats
- [ ] structs/types (some way to group items together)
- [x] <= and >=
- [x] for(& while) loop
- [ ] Get input from STDIN (scanf like function?)

#### Arrays

- [x] Pop
- [ ] Insert
- [x] Remove
- [x] Shift
- [x] Unshift
- [x] includes

#### Integers/Floats

- [x] Random
- [x] Min
- [x] Max
- [ ] Modulo

#### Strings

- [x] split
- [x] trim
- [ ] startsWith
- [ ] endsWith
- [x] includes
