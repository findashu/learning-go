# Learning Go

Prerequisites
 * [Install Go](https://go.dev/dl/)

## Basic File Structure

* Go works as module
* Define start file of the project as main package
* Go also needs starting point that needs to be defined as `main function`

```go
package main
import "fmt"

func main () {
    fmt.println("Welcome to Go")
}

```

## Go Packages

* Go programs are organized into packages
* Go's standard library, provides diffrent core packages for us to use
* "fmt" is one of these, which you can use by importing it.
* [Go Packages](https://pkg.go.dev)

go mod init booking-app