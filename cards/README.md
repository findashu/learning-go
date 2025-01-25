## Variable

Go is static type language.

`var varName datatype = value`
```go
// one way of declaration
var card string = "Ace of Spades"

// let go decide the datatype (both declaration are 100% equivalent)

// syntax := will be only used when creating new variable or initialise

card := "Ace of Spades"

// re-assignment

card = "Hi there!"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)
}

```

## Basic Go DataTypes

* bool: `true` `false`
* string: `Hi` `Ace of Spades`
* int: `0` `-1000` `999999`
* float64: `10.00001` `0.000009` `-10.003`

## Function declaration

` func funcName() returnDataType {}`

```go

func newCard() string {
    return "Five of Diamonds"
}

```

## List

Go has 2 types data structure to handle list `Array` and `slice`

* Array: Fixed length list of things
* Slice: An array that grow and shrink, every element in a slice must be of same type 

Slices and Arrays both needs to be define with Datatypes

```go

// declaring slice and assigning value

cards := []string{"Diamond", "Ace"}

// adding new value using append, will add new value at the end, append returns new slice which will be assigned to the variable

cards = append(cards, "Five of Diamonds")

scores := []int {1,2,3,4,5,6,7,8}

fmt.Println(socores[0])
fmt.Println(socores[2])

// get a slice scores[startIndexIncluding :upToNotIncluding]

fmt.Println(socores[0:3]) // [1,2,3]
fmt.Println(socores[:3]) // [1,2,3]
fmt.Println(socores[3:]) // [4,5,6,7,8]



```

## Iteration

```go

for i, card := range cards {
    fmt.Println(i, card)
}

```
