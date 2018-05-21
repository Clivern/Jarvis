/*
 Note that field names that start with a capital letter are exported, i.e. can be set or read from other packages.
 Field names that start with a lowercase are private to the current package. The same goes for functions defined in packages
*/
package main

import "fmt"


type Student struct{
    name string
    age int
}

type stu_alias Student

// Create a function that takes the type as an argument.
func pretty(stud *Student) string{
    return fmt.Sprintf("|%s -> %d|", stud.name, stud.age)
}

// Create a function that works on the type
func (stud *Student) self_pretty() string {
    return fmt.Sprintf("|%s -> %d|", stud.name, stud.age)
}

// S is a valid implementation for interface I, because it defines the two methods which I requires.
// Note that this is true even though there is no explicit declaration that S implements I.

type IFace interface {
    SetSomeField(newValue string)
    GetSomeField() string
}

type IFace2 interface {
    GetItem() string
}

type Implementation struct {
    someField string
}

func (i *Implementation) GetSomeField() string {
    return i.someField
}

func (i *Implementation) SetSomeField(newValue string) {
    i.someField = newValue
}

func fI(ifa IFace) {
    fmt.Println(ifa.GetSomeField())
}

func getType(face IFace) {
    switch face.(type){
        case *Implementation:
            fmt.Println("Implementation")
        default:
            fmt.Println("Ooops!")
    }
}

func main() {
    arr := []string{"a", "b", "c"}
    dict := map[string]string{
        "name": "Joe",
        "age": "50",
    }

    joe := new(Student)
    joe.name = "Joe"
    joe.age = 50

    mark := Student{"Mark", 30}

    john := stu_alias{"John", 50}

    helen := Student{}
    helen.name = "Helen"


    var item IFace
    item = &Implementation{"Hello World"}


    item2 := &Implementation{"Hi Buddy"}

    fmt.Printf("%v\n", arr)
    fmt.Printf("%v\n", dict)
    fmt.Printf("%v\n", joe)
    fmt.Printf("%s\n", joe.name)
    fmt.Printf("%s\n", pretty(joe))
    fmt.Printf("%s\n", joe.self_pretty())
    fmt.Printf("%s\n", mark.self_pretty())
    fmt.Printf("%s\n", john.name)
    fmt.Printf("%d\n", john.age)
    //fmt.Printf("%s\n", pretty(john)) // Error cannot use john (type stu_alias) as type *Student in argument to pretty
    //fmt.Printf("%s\n", john.self_pretty()) Error cannot use john (type stu_alias) as type *Student in argument to pretty
    fmt.Printf("%s\n", helen.self_pretty()) //|Helen -> 0|
    fmt.Printf("%s\n", item.GetSomeField())

    fI(item)
    fI(item2)
    getType(item)

    if _, ok := item.(IFace); ok {
        fmt.Println(item.(IFace))
    }else{
        fmt.Println("Noop")
    }
    if _, ok := item.(IFace2); ok {
        fmt.Println(item.(IFace2))
    }else{
        fmt.Println("Noop")
    }

}