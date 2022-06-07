# Personal Go Notes ✍️
These are personal notes taken from [Get Programming with Go](https://www.manning.com/books/get-programming-with-go), a well thought out book by Nathan Youngman and Roger Peppé. I highly recommend you buy the book for a real learning experience.

Herein lie notes I took for myself for reference and revisions.

# Lesson One

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

# Lesson Two

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