# gocalc

[![CI](https://github.com/ottenwbe/gocalc/actions/workflows/ci.yml/badge.svg)](https://github.com/ottenwbe/gocalc/actions/workflows/ci.yml)

Small calculator written in Go.

## Purpose

This repository exists for learning and teaching Go. It is not intended for production use; the goal is to demonstrate idiomatic Go and keep the code approachable for newcomers (like I was when I first wrote this code).

## Build

```go
go build -o gocalc
```

## Usage 

This initial version of the calculator relies on a postfix evaluation. 

```
./gocalc 5 5 +
```

## How it works

`gocalc` evaluates terms in **postfix** notation: the operator comes *after* its operands, so no parentheses or precedence rules are needed. Reading left to right, numbers are pushed onto a stack; when an operator is encountered, the top two stack values are popped, the operator is applied, and the result is pushed back.

Example: `5 3 +` → push 5, push 3, apply `+` → result `8`. A longer term `2 3 + 4 *` means `(2 + 3) * 4` = `20`.

## Tests

```
go test ./...
```

## Notes

This repository was also used to test the capabilities of Vibe Code, Mistral AI's async software-engineering agent, which performed the cleanup and modernization of this project.
