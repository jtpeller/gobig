# gobig

[![Go Reference](https://pkg.go.dev/badge/github.com/jtpeller/gobig.svg)](https://pkg.go.dev/github.com/jtpeller/gobig)

## What

gobig is a golang package for useful wrapper functions to make using golang's math/big package easier to use.

## Why

This package exists to allow easier usage of golang's math/big package. For instance, using `New(int64)` instead of `big.NewInt(int64)` means fewer keystrokes and easier-to-read code, especially with complicated mathematical formulae.

## How

Installation can be achieved by:

1. Use this command to install the package:

```
go get -u github.com/jtpeller/gobig
```

2. Import it:

```go
import "github.com/jtpeller/gobig"
```
