# gocalc

[![CI](https://github.com/ottenwbe/gocalc/actions/workflows/ci.yml/badge.svg)](https://github.com/ottenwbe/gocalc/actions/workflows/ci.yml)
 [![codecov](https://codecov.io/gh/ottenwbe/gocalc/branch/master/graph/badge.svg)](https://codecov.io/gh/ottenwbe/gocalc)
 
Small calculator written in Go.

## Purpose

This repository exists for learning and teaching Go. It is not intended for production use; the goal is to demonstrate idiomatic Go and keep the code approachable for newcomers.

## Build

```go
go build -o gocalc
```

## Usage 

This initial version of the calculator relies on a postfix evaluation. 

```
./gocalc 5 5 +
``` 
