# Chapter 2: Data Types in Go

This chapter covers the the fundamental data types available in Go.

## 🎯 Key Concepts

* **Basic Types**: Understanding `int`, `float64`, `bool`, and `string`.
* **Zero Values**: What a variable holds by default when it's declared but not initialized.

## Strings

* For textual data, defined with Double quotes, e.g: "This is a String"

```go

    var name string = "Hello World!"

```

## Intgers

* Representing whole numbers, positive and negative, e.g: 5, 340, -50
* Go provides a wide variety of integer types that differ in size and whether they are signed (can be positive, negative, or zero) or unsigned (can only be zero or positive).

### 🔢 Signed and Unsigned Integers

Go's integer types are explicitly sized, which helps prevent overflow errors and makes memory usage predictable.

* **Signed Integers (`int`):** These can hold both positive and negative numbers. The name indicates the number of bits used to store the value.

| Type    | Size    | Range                                                 |
| :------ | :------ | :---------------------------------------------------- |
| `int8`  | 8 bits  | -128 to 127 (`-2^7` to `2^7-1`)                        |
| `int16` | 16 bits | -32,768 to 32,767 (`-2^15` to `2^15-1`)            |
| `int32` | 32 bits | -2,147,483,648 to 2,147,483,647 (`-2^31` to `2^31-1`) |
| `int64` | 64 bits | `-2^63` to `2^63-1`                                 |

* **Unsigned Integers (`uint`)**: These can only hold non-negative numbers.

| Type     | Size    | Range                     |
| :------- | :------ | :------------------------ |
| `uint8`  | 8 bits  | 0 to 255 (`2^8-1`)        |
| `uint16` | 16 bits | 0 to 65,535 (`2^16-1`)    |
| `uint32` | 32 bits | 0 to 4,294,967,295 (`2^32-1`) |
| `uint64` | 64 bits | 0 to `2^64-1`           |

### ⚙️ Architecture-Dependent Types

Go includes three special integer types whose size depends on the system's architecture.

* **`int`**: The most common integer type. It's either **32 bits** on a 32-bit system or **64 bits** on a 64-bit system.
* **`uint`**: The unsigned equivalent of `int`.
* **`uintptr`**: An unsigned integer large enough to store the bit pattern of any pointer. This is for advanced use cases.

The **zero value** for all integer types is **0**. If you declare an integer variable without initializing it, its value will be `0`.

### How to choose right type:

* For general-purpose use (counting, indexing), use `int`. It's the standard and most efficient choice.
* Use explicitly sized types like `int64` or `uint8` when you need to control the exact size for memory optimization, database mapping, or when working with binary data.
* Use **unsigned integers** (`uint`) only when you are certain the value will never be negative.



```go

    var myNum int // myNum is automatically assigned to zero
    var numberOfDoors int = 20

```

## Floating-Point Types

* `float32` is a single-precision float. It's less precise but uses half the memory of a `float64`.
* `float64` is a double-precision float. It offers much greater precision and is the default and most commonly used floating-point type in Go
* When you declare a floating-point number without specifying a type, Go automatically defaults to `float64`
* The zero value for both `float32` and `float64` is **0.0**.

```go

// For variable declarations of float type, you MUST be explicit.
var price float32 = 19.99
var pi float64 = 3.1415926535

// But when using the := operator, Go defaults to float64 for you.
tax := 0.07 // Go infers this as float64

```

## Array and Slices

In Go, arrays and slices are both used to manage sequences of data, but they have one fundamental difference: an array's size is fixed, while a slice's size is dynamic

### 📦 Arrays: The Fixed-Size Container

An array is a numbered sequence of elements of a single, specific type. Its size is defined at declaration and cannot be changed.

* **Fixed Length:** The length is part of its type. An `[3]int` is a different type from a `[4]int`.
* **Value Type:** When you pass an array to a function, you are passing a copy of the entire array. Changes made inside the function do not affect the original array.

```go

// Declares an array of 3 integers. All elements are initialized to 0.
var a [3]int

// Declares and initializes an array in one line.
primes := [5]int{2, 3, 5, 7, 11}

```

### ✂️ Slices: The Flexible View

A slice is a more flexible and powerful view into the elements of an array. Slices are far more common in Go than arrays.

* *Dynamic Length:* You can add elements to a slice, and it will grow automatically.
* *Reference Type:* A slice doesn't store any data itself; it just describes a section of an underlying array. When you pass a slice to a function, you are passing a reference. Changes made inside the function will affect the original slice's underlying data.

```go

// Create a slice from an existing array.
primes := [5]int{2, 3, 5, 7, 11}
var s []int = primes[1:4] // s is a slice containing {3, 5, 7}

// Create a slice with the `make` function (more common).
// Creates a slice for 5 integers, with capacity for 5.
mySlice := make([]string, 5)

// You can add to a slice using the `append` function.
mySlice = append(mySlice, "new_element")


```
### 🤔 When to Use Which

The simple rule in Go is: use slices unless you have a specific reason to use an array.

* Use a slice for almost everything, especially when the number of elements might change or is unknown. This is the idiomatic way in Go.

* Use an array only when you need a collection of a specific, fixed size that will never change, and you want to enforce that rule at compile time.