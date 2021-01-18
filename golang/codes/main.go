package main

import (
    "fmt"
    "strconv"
)

type Parser interface {
    Parse() string
}

type IntegerParser struct {
    Message int
}

func (ip IntegerParser) Parse() string {
    return strconv.Itoa(ip.Message)
}

type BooleanParser struct {
    Message bool
}

func (ip BooleanParser) Parse() string {
    return strconv.FormatBool(ip.Message)
}

func main() {
    var example Parser = IntegerParser{Message: 200}
    fmt.Println(example.Parse())

    var example1 Parser = BooleanParser{Message: false}
    fmt.Println(example1.Parse())
}