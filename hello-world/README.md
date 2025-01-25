## Go CLI

* go build: Compiles a bunch of go source code files
* go run: Compliles and executes one or two files
* go fmt: Formats all the code in each file in the current directoy
* go install: Compiles and "installs" a package
* go get: Downloads the raw source code of someone else's package
* go test: Runs any tests associated with the current project

## Packages

* Executable: Generates a file that we can run
    * package main: Defines a package that can be compiled and then `executed`. *Must have a func called 'main'*
* Reusable: Code used as 'helpers', good place to put reusable logic.
    * package calculator: Defines a package that can be used as dependency (helper code)
    * package: uploader: Defines a package that can be used as dependency (helper code)

