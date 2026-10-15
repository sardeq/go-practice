
For a new project:
```bash
mkdir go-practice
cd go-practice
go mod init example.com/go-practice
```
create main.go and run:
```bash
go run .
```

_go get_ adds, updates, or removes dependencies in your Go module.
```bash
go get github.com/google/uuid
```
> This adds the github.com/google/uuid package to your project's dependencies, usually updating go.mod and go.sum.

_go mod tidy_ synchronizes your project's dependencies with the packages your source code actually uses.

```bash
go mod tidy
```
- Adds missing dependencies required by your source code.
- Removes unused dependencies that are no longer needed.
- Updates go.mod and go.sum to keep dependency information consistent.



| Command | Purpose |
| :--- | :--- |
| **go get package/path** | Add or update a specific dependency |
| **go get package/path@version** | Request a specific dependency version |
| **go mod tidy** | Add missing and remove unused dependencies |
| **go mod download** | Download dependencies recorded in the module files |
| **go mod init name** | Initialize a new Go module |



# Variables
```Go
// Declare with an explicit type
var name string = "Sadeq"
var age int = 22

// Let Go infer the type
var score = 95.5

// Short declaration (most common inside functions)
country := "Jordan"

// Declare without an initial value
var isStudent bool // defaults to false

// Multiple variables
x, y := 10, 20

// Constants cannot be reassigned
const pi = 3.14159

fmt.Println(name, age, score, country)
fmt.Println(isStudent, x, y, pi)
```

# Data Types
```Go
var name string = "Alex"
var age int = 25
var temperature float64 = 23.5
var isActive bool = true

var small int8 = 100
var number16 int16 = 1000
var number32 int32 = 100000
var number64 int64 = 10000000000

var unsigned uint = 100
var byteValue byte = 'A'
var runeValue rune = '$'
```

Go does not automatically convert between different numeric types.
```Go
var x int = 10
var y int64 = int64(x)
```

# Arrays
```Go
var numbers [3]int = [3]int{10, 20, 30}

// Let Go count the elements
scores := [...]int{90, 85, 100}

// Access elements (index starts at 0)
fmt.Println(numbers[0]) // 10
fmt.Println(scores[2])  // 100

// Modify an element
numbers[1] = 50

// Get the array length
fmt.Println(len(numbers)) // 3

fmt.Println(numbers)
```

# Slices
```Go
// Create a slice
numbers := []int{10, 20, 30}

// Add elements
numbers = append(numbers, 40)
numbers = append(numbers, 50, 60)

// Access and modify
fmt.Println(numbers[0]) // 10
numbers[1] = 25

// Get length and capacity
fmt.Println(len(numbers))
fmt.Println(cap(numbers))

// Slice an existing slice
part := numbers[1:4]
fmt.Println(part) // [25 30 40]

// Create an empty slice with a starting capacity
values := make([]int, 0, 10)
values = append(values, 100)
```

# Maps
```Go
// Create a map
ages := map[string]int{
    "Alex": 22,
    "Sara": 24,
}

// Read a value
fmt.Println(ages["Alex"]) // 22

// Add or update
ages["John"] = 30
ages["Alex"] = 23

// Check whether a key exists
age, exists := ages["Sara"]
if exists {
    fmt.Println(age)
}

// Delete a key
delete(ages, "John")

fmt.Println(ages)
```
To initialize an empty Map
```Go
scores := make(map[string]int)
scores["Math"] = 95
```

# Conditions
```Go
 age := 20

if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Child")
}

// Declare a variable inside the if statement
if score := 85; score >= 50 {
    fmt.Println("Passed")
} else {
    fmt.Println("Failed")
}
```

# Switch
```Go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Almost the weekend")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Regular day")
}
```

# Loops
### For Loop
```Go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
### While Loop
```Go
i := 0

for i < 5 {
    fmt.Println(i)
    i++
}
```
### Infinite Loop
```Go
for {
    fmt.Println("Running")
    break
}
```
### Loop through a slice
```Go
names := []string{"Alex", "Sara", "John"}

for index, name := range names {
    fmt.Println(index, name)
}
```
>range gives the index and value.

_If you don't need the index, use _ to ignore it:_
```Go
for _, name := range names {
    fmt.Println(name)
}
```

# Functions
### Function Declaring Pattern
```Go
func functionName(parameter type) returnType {
    return value
}
```
>A function with no return value does not need a return type. A function can also return multiple values, often including an error.

```Go
package main

import "fmt"

// No parameters, no return value
func greet() {
    fmt.Println("Hello!")
}

// Parameters and a return value
func add(a int, b int) int {
    return a + b
}

// Same-type parameters can share a type
func multiply(a, b int) int {
    return a * b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }

    return a / b, nil
}

func main() {
    greet()

    result := add(5, 3)
    fmt.Println(result)

    answer, err := divide(10, 2)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(answer)
}
```

# Structs
```Go
package main

import "fmt"

type User struct {
    Name  string
    Age   int
    Email string
}

func main() {
    user := User{
        Name:  "Sadeq",
        Age:   22,
        Email: "sadeq@example.com",
    }

    fmt.Println(user.Name)
    fmt.Println(user.Age)

    user.Age = 23
    fmt.Println(user)
}
```

# Pointers
```Go
package main

import "fmt"

func main() {
    x := 10

    p := &x // Address of x

    fmt.Println(*p) // Read the value: 10

    *p = 20 // Change x through the pointer

    fmt.Println(x) // 20
}
```

# Methods
```Go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}

    fmt.Println(rect.Area()) // 15
}
```

# Interfaces
```Go
package main

import "fmt"

type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func makeSound(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    makeSound(Dog{})
    makeSound(Cat{})
}
```

# Errors
```Go
package main

import (
    "errors"
    "fmt"
)

func withdraw(balance, amount int) (int, error) {
    if amount <= 0 {
        return balance, errors.New("amount must be positive")
    }

    if amount > balance {
        return balance, errors.New("insufficient funds")
    }

    return balance - amount, nil
}

func main() {
    balance, err := withdraw(100, 30)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Remaining:", balance)
}
```

# Defer
```Go
package main

import "fmt"

func main() {
    defer fmt.Println("Third")
    defer fmt.Println("Second")

    fmt.Println("First")
}
```
> Deferred calls run in reverse order. This is commonly used to close files and release resources.

# Packages and Imports
```Go
package main

import (
    "fmt"
    "strings"
)

func main() {
    name := "sadeq"

    fmt.Println(strings.ToUpper(name))
}
```

# Closures
> A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

```Go
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}
```

# Currying
_Function currying turns a function with multiple arguments into successive functions that each accept one argument._

For example, we can supply a shirt's unit price now and the quantity later:

```Go
func main() {
	shirtCost := curriedCost(20)
	fmt.Println(shirtCost(3))
	// prints 60

	fmt.Println(shirtCost(5))
	// prints 100
}

func cost(unitPrice, quantity int) int {
	return unitPrice * quantity
}

func curriedCost(unitPrice int) func(int) int {
	return func(quantity int) int {
		return cost(unitPrice, quantity)
	}
}
```
>curriedCost(20) returns a closure that remembers the unit price and waits for a quantity. Calling shirtCost(3) supplies that quantity and calls cost(20, 3).

You can also call the returned function immediately:
```Go

cost(20, 3)        // returns 60
curriedCost(20)(3) // returns 60
```