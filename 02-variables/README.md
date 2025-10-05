# Chapter 2: Variables

This chapter covers the basics of declaring variables and the fundamental data types available in Go.

## 🎯 Key Concepts

* **Variable Declaration**: Using `var`, the `:=` short declaration operator, and declaring variables in blocks.
* **Basic Types**: Understanding `int`, `float64`, `bool`, and `string`.
* **Zero Values**: What a variable holds by default when it's declared but not initialized.

## Varibale Declarations

Go is statically typed, but it offers flexibility in how you declare variables. Go's compiler enforces a simple but powerful rule: `if you declare a variable, you must use it` 

* The Explicit Way `(var name string = "Go")`:
You state the variable's name, type, and value. It's perfectly clear but verbose. Think of it as being formal.
* The Inferred Way `(var name = "Go")`:
You let Go figure out the type from the value you provide. This is Go's type inference in action and is a common way to use `var`. It's important to know that this method works only if you initialise the variable with the value.
* The Idiomatic Way `(name := "Go")`:
The short declaration operator `(:=)` is the most popular method for declaring and initializing variables inside functions. It's concise and clean.
* Constants are declared with the `const` keyword and are **immutable**. Their value must be known at compile time.

## 💻 Code Example

The code for this chapter is in the `main.go` file. It demonstrates:

1.  How to declare and initialize variables using different methods.

## ▶️ How to Run

Navigate to this directory in your terminal and run the Go file:

```bash
go run main.go
```