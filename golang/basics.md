---
layout: page
permalink: /golang/basics
---

<br>

<h1><b>GoLang Basics</b></h1>
Welcome to the basics part of GoLang Programming

<h1><b>Table Of Contents</b></h1>

* TOC
{:toc}

### Running Go Code
Let’s start our journey by creating a simple program and learning how to compile and execute it. Open your favorite text editor and write the following code, save the file as `main.go`
```go
package main

import "fmt"

// In Go, the entry point to a program has to be a
// function called `main` within a package `main`
func main() {
    println("Hello, World")
}
```
Finally, run the program by entering the following command 
```shell
$ go run main.go
Hello, World
```
What about compilation though? `go run` is a handy command that compiles and runs your code. It uses a temporary directory to build the program, executes it and then cleans itself up. You can see the location of the temporary file by running
```shell
$ go run --work main.go
```
To explicitly compile code, use go build
```shell
$ go build main.go
```
This will generate an executable main which you can run. On Linux, don’t forget that you need to prefix the executable with dot-slash, so you need to type `./main`. While developing, you can use either `go run` or `go build`. When you deploy your code however, you’ll want to deploy a binary via `go build` and execute that.

*We’ll talk more about packages in a later chapter. For now, while we focus on understanding the basics of Go, we’ll always write our code within the main package.*

__*Imports :*__ In Go, the import keyword is used to declare the packages that are used by the code. We're using one of Go's standard packages - `fmt`. You've probably noticed we prefix the function name with the package, e.g., `fmt.Println`. Go is strict about importing packages. It will not compile if you import a package but don't use it. Try to run the following
```go
package main

import "fmt"

func main() {
}
```
Output
```shell
$ go run main.go
# command-line-arguments
main.go:4:2: imported and not used: "fmt"
```
__go get :__ Go is strict about this because unused imports can slow compilation; admittedly a problem most of us don’t have to this degree. Another thing to note is that Go’s standard library is well documented. You can head over to [https://golang.org/pkg/fmt/#Println](https://golang.org/pkg/fmt/#Println) to learn more about the `Println` function that we used. To install `godoc` use the `go get` command
```shell
$ go get -v golang.org/x/tools/cmd/godoc
$ go doc fmt.Println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```
If you’re ever stuck without internet access, you can get the documentation running locally via `godoc`
```shell
$ sudo apt install golang-golang-x-tools
$ godoc
```
And Navigate to `http://localhost:6060` on browser

### Variables - Declarations and Types
```go
package main

import "fmt"

func main() {
    var variableA int
    variableA = 20
    variableB := 30
    fmt.Println(variableA, " ", variableB)
}
```
Here, we declare two variables `variableA` and `variableB` using to different declarations. The first one we declare a variable `variableA` of type `int`. By default, Go assigns a zero value to variables. Integers are assigned `0`, booleans `false`, strings `""` and so on. Next, we assign `20` to our `variableA`. We can merge the first two lines
```go
var variableA int = 20
```
Still, that's a lot of typing. Go has a handy short variable declaration operator, `:=`, which can infer the type
```go
variableB := 30
```
__Types :__ Go has various types including strings, integers, floats, booleans, etc. The `TypeOf()` method of the reflect package is used to determine the datatype of the variables. Here's an example
```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    // strings can be added together with +
    fmt.Println("go" + "lang", reflect.TypeOf("go"))
    // integers
    fmt.Println("1 + 2 =", 1+2, reflect.TypeOf(1))
    // floats
    fmt.Println(7.0/2, reflect.TypeOf(7.0/2))
    // booleans
    fmt.Println(true && false, true || false, !true, reflect.TypeOf(true))
}
```
Output
```shell
$ go run golang/codes/main.go 
golang string
1 + 2 = 3 int
3.5 float64
false true false bool
```

__Constants :__ Go supports *constants* of character, string, boolean, and numberic values. `const` declares a constant value. A numberic constant has no type until it's given one, such as by an explicit conversion
```go
const s string = "string constant"
const i1 int = 30
const i2 = 40
```

__Note #1 :__ It’s important that you remember that `:=` is used to declare the variable as well as assign a value to it. Why? Becausea variable can’t be declared twice (not in the same scope anyway). If you try to run the following, you’ll get an error.
```go
package main

import "fmt"

func main() {
    varA := 20
    fmt.Println(varA);
    varA := 30 // will throw error
    fmt.Printlb(varB);
}
```
Output
```shell
$ go run main.go
# command-line-arguments
main.go:10:10: no new variables on left side of :=
```
The above error means that when we first declare a variable,we use `:=` but on subsequent assignment, we use the assignment operator `=`. If you read the error message closely, you’ll notice that *variables* is plural. That’s because Go lets you assign multiple variables (using either = or :=).
```go
// as long as one of the variables is new := can be used
varA, varB := 20, 30
fmt.Println(varA, varB)
varB, varC := 25, 35
fmt.Println(varB, varC)
```

__Note #2 :__ Go won't let you have unused variables. Like unused imports it'll throw an error when unused variables are present. Let's look at an example
```go
package main

import (
    "fmt"
)

func main() {
    varA, varB := 20, 30
    fmt.Println(varA);
}
```
Output
```shell
$ go run main.go
# command-line-arguments
main.go:8:11: varB declared and not used
```

### Functions
Like many other languages, functions in Golang allow us to divide code into useful blocks, make it more readable, reuse it. Like in Python, functions in Golang can return multiple values. We declare functions in Golang using the `func` keyword. Let's look at some example
```go
// takes in one argument of int type and no return value
func function1(arg1 int) {
}

// takes in two arguments, one of int and the other is a string
// and returns an int value
func function2(arg1 int, arg2 string) int {
    return 1
}

// takes in two arguments, both of type int and returns two
// values one is int and the other is a string
func function3(arg1, arg2 int) (int, string) {
    return 1, "1"
}
// see the shorthand for argument types, if they are of the
// same type then we can do the above instead of explicitly specifying the type
```
In the last case we can use
```go
a, b := function3(1, 3)
// if we want to ignore the first value then we can do so by using _
_, b := function3(1, 3)
```
This is more than a convention. `_, the blank identifier`, is special in that the return value isn’t actually assigned This lets you use `_` over and over again regardless of the returned type.

## Structures
GoLang isn't an object-oriented (OO) language like C++ and Java. It doesn't have objects nor inheritance and thus, doesn't have the many concepts associated with OO such as polymorphism and overloading. But GoLang does have structures , which can be associated with methods. It supports a simple but effective form of composition. Although Go doesn’t do OO like you may be used to, you’ll notice a lot of similarities between the definition of a structureand that of a class. The following is the syntax for declaring a `struct` in Go
```go
// type and struct are keywords
// Person is the name of the structure
// Name and Age are attributes of the Person structure
type Person struct {
    Name string
    Age int
}
```
Let's take a look at Declaring and Initializing the struct
```go
package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age int
}

func main() {
    p1 := Person{
        Name: "John",
        Age: 22,
    } // the trailing , in the above structure is required
    fmt.Println(p1)

    p2 := Person{}
    // Just like unassigned variables have a zero value, so do fields.
    fmt.Println(p2)
    // or
    p2 = Person{Name: "Jane"}
    fmt.Println(p2)
    p2.Age = 24
    fmt.Println(p2)
}
```
Output
```shell
$ go run golang/codes/main.go 
{John 22}
{ 0}
{Jane 0}
{Jane 24}
```
__Pointers :__ Many times though, we don't want a variable that is directly associated with our value but rather  variable that has a *pointer* to our value. A *pointer* is a memory address; it's the location of where to find the actual value. Why do we need pointers, lets look at an example
```go
func main() {
    p1 := Person("John", 20)
    increaseAge(p1)
    fmt.Println(p1)

    p2 := Person("John", 20)
    increaseAge1(&p2)
    fmt.Println(p2)
}

func increaseAge(p Person) {
    p1.Age += 1
}

func increaseAge1(p *Person) {
    p1.Age += 1
}
```
Output
```shell
$ go run main.go 
{John 20}
{John 21}
```
The above happens because *increaseAge* makes changes to the copy of the *p1* and thus, changes made in *increaseAge* weren't reflected in the caller. To make this work, we need to pass a pointer to the function as is the case for *increaseAge1* and *p2*.

__Note #1 :__ The `&` operator is used to get the address of a value, whereas `*x` means *pointer to the value of type X*.

__Creating Structures :__ Despite the lack of constructors, Go does have a built-in `new` function which is used to allocate the memory required by the type.
```go
p1 := new(Person)
// is the same as
p1 := &Person{}
```
Which you use is upto you, but you'll find that most people prefer the latter. A structures don't have constructors, many prefer to create a function that returns an instance of the desired type
```go
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age: age,
    }
}
```

__Methods on Structures :__ We can associate a method with a structure, like the way we associate methods in classes.
```go
type Person struct {
    Name string
    Age int
}

func (p *Person) increaseAge() {
    p.Age += 1
}
```
In the above code, we say that the type **Person* is the *receiver* of the *increaseAge* method. We call *increaseAge* as follows
```go
p1 = Person{"John", 20}
p1.increaseAge()
fmt.Println(p1.Age)
```

__Note #2 :__ The fields of a `struct` can be of any type - including arrays, maps, interfaces and functions.