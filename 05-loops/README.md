# Chapter 2: Loops

In Go, loops are incredibly simple because there is only one looping keyword: `for`

However, this single for loop is versatile enough to be used in several different ways, covering all the looping patterns you might be familiar with from other languages.

## 1. The Classic `for` Loop

This is the most common loop structure, similar to `for` loops in C, Java, or JavaScript. It has three components separated by semicolons:

* **Init:** A statement executed before the first iteration (e.g., declaring a counter).
* **Condition:** An expression evaluated before every iteration. The loop stops if it's false.
* **Post:** A statement executed at the end of every iteration (e.g., incrementing a counter).

```go

for i := 0; i < 5; i++ {
    fmt.Println(i) // Prints 0, 1, 2, 3, 4
}

```

## 2. The "While" Loop Style

You can create the equivalent of a `while` loop by dropping the init and post statements, leaving only the condition.

```go

n := 0

for n < 5 {
    fmt.Println(n)
    n++
}

```

## 3. The Infinite Loop

If you drop the condition as well, you create an `infinite` loop. You'll need a way to exit the loop, typically with a break statement.

```go

for {
    fmt.Println("This will run forever...")
    // You would have some condition to exit the loop
    // if someCondition {
    //     break
    // }
}

```

## 4. The for...range Loop (Most Idiomatic)

This is the most common and powerful looping construct in Go. It's used to iterate over elements in a collection like a slice, array, map, string, or channel.

* Iterating over a Slice or Array

When used with a slice or array, `for...range` returns two values on each iteration: the `index` and a` copy of the value` at that index.

```go

nums := []string{"one", "two", "three"}
for index, value := range nums {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}

```

If you don't need the index, you can ignore it with the blank identifier (_).

```go

for _, value := range nums {
    fmt.Println(value)
}

```

* Iterating over a Map

When used with a map, it returns the **key** and the **value**.

```go

capitals := map[string]string{"USA": "Washington D.C.", "Japan": "Tokyo"}
for key, value := range capitals {
    fmt.Printf("The capital of %s is %s\n", key, value)
}

```