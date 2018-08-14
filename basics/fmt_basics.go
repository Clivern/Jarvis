package main

// https://golang.org/pkg/fmt/
import "fmt"

func main() {
    x := "World"
    y := 70
    var s string
    var i int
    fmt.Println("Hello World\n")
    fmt.Print("Hello World\n")
    fmt.Printf("Hello World\n")
    fmt.Printf("Hello %s\n", "World")
    fmt.Printf("Hello %s\n", x)
    fmt.Printf("My Age %d\n", y)
    fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
    fmt.Printf("var s=%s\n", s)
    fmt.Printf("var i=%d\n", i)
    fmt.Printf("%6.2f", 12.0285)
}