package main

import "fmt"

func main() {

    /*
    var (
        z int
        k bool
    )
    */

    // Constants
    const (
        c = 56
        l = "hko"
    )
    //var x, y int
    //var d int32

    _, b := 34, 35

    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(l)
    fmt.Println(l[0])


    var first_name1 string
    var last_name1 string
    var age1 int
    var is_male1 bool

    first_name1 = "Joe"
    last_name1 = "Doe"
    age1 = 30
    is_male1 = true

    first_name2 := "Jessie"
    last_name2 := "Mole"
    age2 := 40
    is_male2 := false

    if is_male1 {
        fmt.Printf(first_name1 + " " + last_name1 + " is a Male with %d Years old.\n", age1)
    }else{
        fmt.Printf(first_name1 + " " + last_name1 + " is a Female with %d Years old.\n", age1)
    }
    if is_male2 {
        fmt.Printf(first_name2 + " " + last_name2 + " is a Male with %d Years old.\n", age2)
    }else{
        fmt.Printf(first_name2 + " " + last_name2 + " is a Female with %d Years old.\n", age2)
    }
}