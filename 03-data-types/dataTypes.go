package main

import "fmt"

func main() {
	// Integer types
	var a int = 42
	var b int8 = -128
	var c uint16 = 65535 // Unsigned integer only positive values

	// Floating-point types
	var d float32 = 3.14
	var e float64 = 2.718281828459045
	//var k float = 100.6506 // with through have to specify float32 or float64

	// Boolean type
	var f bool = true

	// String type
	var g string = "Hello, Go!"

	// Complex types
	var h complex64 = 1 + 2i
	var i complex128 = 3 + 4i

	// Print the variables and their types
	fmt.Printf("a: %d (type: %T)\n", a, a)
	fmt.Printf("b: %d (type: %T)\n", b, b)
	fmt.Printf("c: %d (type: %T)\n", c, c)
	fmt.Printf("d: %f (type: %T)\n", d, d)
	fmt.Printf("e: %f (type: %T)\n", e, e)
	fmt.Printf("f: %t (type: %T)\n", f, f)
	fmt.Printf("g: %s (type: %T)\n", g, g)
	fmt.Printf("h: %v (type: %T)\n", h, h)
	fmt.Printf("i: %v (type: %T)\n", i, i)
}
