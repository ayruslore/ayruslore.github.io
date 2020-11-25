---
layout: page
title: GoLang
permalink: /golang/
---
<img src="/images/python.png" width="720" height="200">

__What is GoLang?__ Go is a compiled, statically typed language with a C-like syntax and garbage collection. What does that mean?

__*Compilation*__ is the process of translating the source code that you write into a lower level language – either assembly (as is the case with Go), or some other intermediary language (as with Java and C#). Compiled languages can be unpleasant to work with because compilation can be slow. It’s hard to iterate quickly if you have to spend minutes or hours waiting for code to compile. Compilation speed is one of the major design goals of Go. This is good news for people working on large projects as well as those of us used to a quick feedback cycle offered by interpreted languages. Compiled languages tend to run faster and the executable can be run without additional dependencies (at least, that’s true for languages like C, C++ and Go which compile directly to assembly).

Being __*statically typed*__ means that variables must be of a specific type (int, string, bool, []byte, etc.). This is either achieved by specifying the type when the variable is declared or, in many cases, letting the compiler infer the type.

__*Garbage Collected*__ : Some variables, when created, have an easy-to-define life. A variable local to a function, for example, disappears when the function exits. In other cases, it isn’t so obvious – at least to a compiler. For example, the lifetime of a variable returned by a function or referenced by other variables and objects can be tricky to determine. Without garbage collection, it’s up to developers to free the memory associated with such variables at a point where the developer knows the variable isn’t needed. How? In C, you’d literally free(str); the variable. Languages with garbage collectors (e.g., Ruby, Python, Java, JavaScript, C#, Go) are able to keep track of these and free them when they’re no longer used. Garbage collection adds overhead, but it also eliminates a number of devastating bugs.

__Getting Started :__ If you’re looking to play a little with Go, you should check out the [Go Playground](https://play.golang.org/) which lets you run code online without having to install anything. This is also the most common way to share Go code when seeking help in [Go’s discussion forum](https://groups.google.com/forum/#!forum/golang-nuts) and places like StackOverflow. Installing Go is straightforward. You can install it from source, When you [go to the download page](https://golang.org/dl/), you’ll see installers for various platforms. Follow the instructions to install go in Linux:
* Visit the [Go Download Page](https://golang.org/dl/) and click on Download Go for Linux.
* Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go.

    `sudo tar -C /usr/local -xzf go1.15.5.linux-amd64.tar.gz`
* Add /usr/local/go/bin to the PATH environment variable.

    `export PATH=$PATH:/usr/local/go/bin`
* Verify that you've installed Go by opening a command prompt and typing the following command

    `go version`
* Confirm that the command prints the installed version of Go.

__Version :__ `go version go1.15.5 linux/amd64`

__Different Links to learn GoLang from?__
* GoLang [Tutorials Point](https://www.tutorialspoint.com/go/go_overview.htm)
* [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/) Book
* [Many Other Sources](https://hackr.io/tutorials/learn-golang/page/2)

__GoLang Programming Paradigms__
* Compiled
* Statically Typed

## Contents
* [Basics Tutorial](/golang/basics)

<br>

<b>Note :</b> The following documentation is just for pocket reference, placement preperations, etc.

Links I Use:
* https://gobyexample.com/
* http://www.golang-book.com/books/intro
