package main

import "fmt"


func name(first string, last string) string {
    return first + " " + last
}

func show_name(first string, last string) {
    fmt.Println(first + " " + last)
    return
}

func callback(y int, f func(x int) int) int{
    return f(y)
}


func sum(x int, y int) int {
    z := x + y

    if z == 5 {
        fmt.Println(z)
        return z
    }

    if z == 6 {
        fmt.Println(z)
        return z
    }

    fmt.Println(z)
    return z
}

func sum_with_defer(x int, y int) int {
    z := x + y

    defer fmt.Println(z)

    if z == 5 {
        return z
    }

    if z == 6 {
        return z
    }

    return z
}

// Call function on return
func f() (ret int) {
    defer func() {
        ret++
    }()
    return 0
}

func var_arg(arg ...string){
    for _, value := range arg{
        fmt.Println(value)
    }
}

func main() {
    fmt.Println(name("Joe", "Doe"))
    show_name("Joe", "Doe")

    a := func() string{
        return "Hello World!"
    }
    fmt.Println(a())

    dict := map[string]func() string{
        "a": func() string { return "A Value" },
    }
    fmt.Println(dict["a"]())


    call := func (x int) int {
        return x * 20
    }

    fmt.Println(callback(3, func (x int) int {
        return x + 20
    }))
    fmt.Println(callback(3, call))
    //Callbacks

    sum(1, 2)
    sum(3, 2)
    sum(3, 3)

    sum_with_defer(1, 2)
    sum_with_defer(3, 2)
    sum_with_defer(3, 3)

    fmt.Println(f())
    var_arg("a", "b", "c")
}