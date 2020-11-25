package main

import (
    "fmt"
)

type Person struct {
    Name string
    Age int
}

func main() {
    p1 := Person{"John", 20}
    increaseAge(p1)
    fmt.Println(p1)

    p2 := Person{"John", 20}
    increaseAge1(&p2)
    fmt.Println(p2)
}

func increaseAge(p Person) {
    p.Age += 1
}

func increaseAge1(p *Person) {
    p.Age += 1
}