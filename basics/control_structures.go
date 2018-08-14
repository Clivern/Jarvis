package main

import "fmt"

func test1(x int) int {
    return x
}

func test2(y string) string {
    return y
}

func label() {
    i := 0
    again:
        fmt.Println(i)
        i++
        if i <= 5 {
            goto again
        }
}

func loop() {
    for i := 0; i <= 5; i++ {
        fmt.Println(i)
    }
}

func  loop_with_break() {
    for i := 0; i <= 1000; i++ {
        fmt.Println(i)
        if i == 5 {
            break
        }
    }
}

func loop_with_continue() {
    for i := 0; i <= 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }
}

func ranges() {
    items := []string{"a", "b", "c", "d"}
    for k,v := range items {
        fmt.Printf("%d -> %s\n",k,v)
    }

    for k, v := range "Hello!" {
        fmt.Printf("%d -> %c\n",k,v)
    }

    x := "Hello"
    fmt.Printf("%c\n", x[1])
}

func switch_call() {
    y := 3
    switch y {
        case 1:
            fmt.Println("1")
        case 2:
            fmt.Println("2")
        case 3:
            fmt.Println("3")
        default:
            fmt.Println("noop")
    }
}

func main() {
    x := 3
    if x == test1(3) {
        fmt.Println("x == test1()")
    }

    if y := test2("hi"); y == "hi" {
        fmt.Println("y == hi")
    }
    label()
    loop()
    loop_with_break()
    loop_with_continue()
    ranges()
    switch_call()
}