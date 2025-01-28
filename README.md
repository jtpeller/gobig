# gobig

[![Go Reference](https://pkg.go.dev/badge/github.com/jtpeller/gobig.svg)](https://pkg.go.dev/github.com/jtpeller/gobig)

## What

gobig *(or go home)*

gobig is a Go package for useful wrapper functions to make using Go's math/big package easier to use.

## Why

This package exists to allow easier usage of Go's math/big package. For instance, using `New(int64)` instead of `big.NewInt(int64)` means fewer keystrokes and easier-to-read code, especially with complicated mathematical formulae.

## How

Installation can be achieved by:

1. Use this command to install the package:

    ```sh
    go get -u github.com/jtpeller/gobig
    ```

2. Import it:

    ```go
    import "github.com/jtpeller/gobig"
    ```

3. Then use it! For my use-case, I write a separate file in the same package that needs the functions. This file wraps the functions I need from gobig as private functions. Then, I can access them without any dot operators for even easier readability, less keystrokes, etc.

Example use on [my OEIS repo](https://github.com/jtpeller/OEIS/)
