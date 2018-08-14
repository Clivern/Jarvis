package main

import "fmt"

func main() {
    x := [] int {1, 2, 3, 4, 5, 6}
    y := [] string {"a", "B", "c"}
    z := [][] string {{"a", "b"}, {"c", "d"}}
    x[2] = 300
    y[2] = "Hello"
    // y[3] = "New" invalid array index 3 (out of bounds for 3-element array)
    slice := y[0:1]
    new_slice := append(slice, "h", "k", "o")

    dict := map[string]string{
        "key1": "value1",
        "key2": "value2",
        "key3": "value3",
        "key4": "value4",
    }

    fmt.Printf("%d\n", x[2])
    fmt.Printf("%s\n", y[2])
    fmt.Printf("%s\n", z[0][0])
    fmt.Printf("%s\n", slice[0])
    fmt.Printf("%s\n", new_slice[2])
    fmt.Printf("%s\n", new_slice[3])
    fmt.Printf("%s\n", dict["key1"])


    for _, value := range dict {
        fmt.Printf("?=%s\n", value)
    }

    for key, _ := range dict {
        fmt.Printf("%s=?\n", key)
    }

    for key, value := range dict {
        fmt.Printf("%s=%s\n", key, value)
    }
}