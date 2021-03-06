# Go Programming Language Study Notes ✍️
These are personal notes taken from [Get Programming with Go](https://www.manning.com/books/get-programming-with-go), a well thought out book by Nathan Youngman and Roger Peppé. I highly recommend you buy the book for a real learning experience.

Herein lie notes I took for myself for reference and revisions.

# 01 Hello, World!

Go unlike Ruby, is a compiled programming language. You can have a feel of Go on the [official playground](https://go.dev/play/) that runs in the browser, keep in mind however, that the output of browser compiler might be cached so your variables could be holding hold values even if the variable values on the screen change.

This is Go's Hello World.

```golang
package main

import "fmt"

func main {
  fmt.Println("Hello, world!")
}
```

`fmt` is pronounced "FOOMT" by Gophers. We have a package, akin to "modules" in Ruby to package code. `fmt` is an imported package of functions or methods you can use. To call a function from a package, you call the package name, then the name of the function like `fmt.Println`. Every Go progam has to have an initiation point, that's denoted with `func main`.

In Go, the opening brace is on the same line as the `func` keyword, whereas the closing brace `}` is on its own line. This is the "one true brace style"([1TBS](https://en.wikipedia.org/wiki/Indentation_style#Variant:_1TBS_(OTBS))), your code won't work any other way. This is because code statements use to end with `;` until Gophers decided to make the compiler silently append these semicolons so don't don't have to, but we need to follow the 1TBS in exchange for this to work. 

# 02 Variables and Printing

Just like other language, Go has constructs to print information to the screen. Among them are:

```go
fmt.Print   // Prints to screen without a line break
fmt.Println // Prints and breaks line
fmt.Printf  // Provides ways to print with values inserted in the string
            // An example would be `fmt.Printf("Jordan is %v cm tall.", 200)`
```

The `%v` seen in `fmt.Printf` is called a *format verb*, which is substituted for the value of the expression provided by the second argument to `fmt.Printf`. There's a lot more to [format verbs](https://pkg.go.dev/fmt).

Constants and variables are represented by `const` and `var` respectively.

You can declare constants in a group like: 

```go
var (
  name = "Jordan"
  height = 300
)
```

Or on a single line like: 

```go
height, width = 50, 50
```

# 03 Loops And Branches

In some languages, some values might be considered *truthy* or *falsy*. Values like `nil` in Ruby are *falsy* and values like `0` and `""`are *truthy*, just not `true` itself. This isn't so in Go. Go has only true and false values and they are `true` and `false`, surprise! 

There are a few ways to get boolean values. Using comparison operators like `==` and `>=` return booleans. Some fuctions from the standard library also return booleans, like:

```go
var transport = "taxi"
var word = string.Contains(transport, "taxi") // => true
```

Go has a single equality operate... `==`. Unlike JS and PHP that have both `==` and `===`.

Branching is straight-forward:

```go
package main

import "fmt"

func main() {
  var age = 20

  if age == 18 {
    fmt.Println("Almost not a minor, come in.")
  } else if age > 20 {
    fmt.Println("Not a minor.")
  } else {
    fmt.Println("Come back in a few years' time.")
  }
}
```

There's also the `switch` for branching:

```go
var name = "Jordan"

switch name {
case "Mike":
  fmt.Print("Mike")
case "John", "Lisa":
  fmt.Print("Nice names")
default:
  fmt.Print("No name") 
}
```

There's a more concise form of the `switch` statement in code [here](https://github.com/siaw23/learning_go_notes/blob/main/lesson_three.md) that also show the use of the `fallthrough` keyword.

Go has logical operators, `||` and `&&` for *or* and *and* respectively. And a "not" operator `!`.

Looping can be done, in several forms with the `for` keyword:

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  var count = 3

  for count > 0 {
    fmt.Println(count)
    time.Sleep(time.Second)
    count--
  }

  fmt.Println("Go!")
}
```

# 04 Variable Scope

Much like many languages, variables in Go have a scope, that is, a region within a portion of the code where these variables cann be accessed. `{` and `}` demacate this regions. Variables inside the curly braces cannot be accessed out of it. The `case` and `default` keywords also define a scope.

Variables in Go are declared like so `var num = 10`, but there's a shorter form where the `var` keyword can be dropped: `num := 10`, a short declaration. This lets you do stuff like:

```go 
for num := 0; num == 10; num++ {
  fmt.Println(num)
}

if num := rand.Int(3); num == 0 {
  // do something
} else if num == 1 {
  // do another thing
}
// You can use short declartion in a `switch` statement too.
```

# 05 Capstone 


# 06 Real Numbers

`float32` and `float64` are 2 of Go's 15 numeric types. The following are equivalent ways to declare a `float64`:

```go 
distance := 256.248
var distance = 256.248
var distance float64 = 256.248
```

Go infers the type of a variable by evaluating the value on the right side of the declaration, making `var distance float64 = 256.248` a bit unncecessary. Wen you declare a variable with a number and a decimal point, the type will always be a `float64`. [Real numbers](https://en.wikipedia.org/wiki/Real_number) are inferred as `float64`. You need to specify the type if you want any of the other numeric types.

`float64` is a 64-bit floating point type (double precision) while `float32` is 32-bit, requiring 8 and 4 bytes in memory.

Each type has a *zero value*, that is a default value that's used a placeholder when you don't declare a value with a value. The zero value of `float64` is `0.0`.
So in this:

```go
var price float64
```

The value of `price` is `0.0`

You can use the `%f` very to format printed values with `Printf`. 

Examples:

```go
fmt.Printf("%.2f", value)
fmt.Printf("%3.2f", value)

// 3 is width, 2 is precision
```

If the width is unspecified, `Printf` will use the number of characters necessary to
display the value.

# 07 Whole Numbers

Go has 10 types to represent integers. Each integer type has a limited range, they can be signed (positive or negative) or unsigned (positive). Integer types can only be whole numbers and don't suffer from the inaccuracies of Floats. `int` is signed, `uint` is unsigned, these types plus the ones listed below make the 10 integer types in Go. Depending on the architecture on a device, `int` or `uint` might be 32-bit or 64-bit. 

Go always picks `int` when inferring the type of a whole number.

The following table shows 8 architectiure-independent integer types:

![Architecture-independent integer types](/assets/digital_ocean_gopher_guides.png "Architecture-independent integer types")

`%T` used with `Printf` will tell the type of a value. Example:

```golang
val bool = false
fmt.Printf("Type: %T Value: %v\n", val, val)
// Type: bool Value: false
```

Integer ranges wrap over when they overflow. Example:

```golang
zero = 0 
zero--
fmt.Println(zero) // 255
// 0 is uint, -1 of that wraps to 255 in binary
```

Integer types should carefully be chosen to avoid wrapping.

# 08 Big Numbers

Go has a `big` package to handle gigantic numbers and will infer a number with exponents like `41.3e12` as `float64`.

`big` provides three types:

```golang
big.Int // for big ints
big.Float 
big.Rat // for fractions like 1/3
```

There are two ways to create a `big.Int`:

```golang
secondsPerDay := big.NewInt(86400)
// OR
secondsPerDay := new(bing.Int)
secondsPerDay.SetString("864000", 10) // 10 indicates base (decimal in this case)
```

Constants in Go can be untyped, unlike variables whose type is inferred. While `const distance uint64 = 24000000000000000000` will overflow with an error constants like `const distance = 24000000000000000000` won't.

Untyped constants must be converted to typed variables when passed to functions.

# 09 Multilingual Text

Backticks indicate a raw string literal.

```golang
fmt.Println("Hello,\nWorld!")
// Hello,
// World!
fmt.Println(`Hello,\n World!`)
// Hello,\n World!
```
`rune` is an alias for `int32` that Go uses to represent a single Unicode code point. A `byte` is an alias for `uint8` intended for binary data. A `byte` can also be use for English characters as defined by ASCII, a subset of Unicode.

```golang
var pi rune = 800 
// The above is equivelent to the snippet below
var pi int32 = 800
```

Strings in Go are immutable.

 # 10 Converting Between Types

You can't mismatch types when performing operations. For example floating point numbers add to other floating point numbers. You need to convert one type to a common type if you need to perform operations on them. 

## Numeric Type Conversion
If you need to convert an integer to a floating type, you need to wrap that variable with the new type:

```golang
height := 188
exactHeight := float64(188)

// Similary you can do:

height := int(400.0)
```
Go provides a `math` package that makes it possible to check whether a type conversion will result in an invalid value. Here's an example:

```golang
if num < math.MinInt16 || num > math.MinInt16 {
  // handle out of range value
}
// This check the limits of `int16` integers.
```
## String Type Conversion
To convert a `rune` or `byte` into a `string`, you can use the same syntax for numeric type conversions as above. 

```golang
var pi rune = 550
fmt.Print(string(pi))
```

You can convert an integer to to string, if for example you're looking to interpolate the value of the integer to be printed. Here's how:

```golang
age := 54
str := "I'm " + strconv.Itoa(age) + " years old."
fmt.Print(str)
```

The `Itoa` stands for "integer to ASCII. On ther other side we have `strconv.Atoi` which is "ASCII to integer.

## Boolean Type Conversion
Just like the string conversion above, you can convert booleans to strings for the purpose of printing with format verbs for example. 

```golang
launch := true
text := fmt.Sprintf("%v", launch)
fmt.Println("Ready for launch:", text)
```

# 11 Capstone 

# 12 Functions

# 13 Methods

# 14 First-Class Functions
