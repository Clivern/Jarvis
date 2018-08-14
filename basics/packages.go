/*
Useful packages
===============
The standard libary of Go includes a huge number of packages. It is very enlightening to browse the $GOROOT/src/pkg directory and look at the packages. We cannot comment on each package, but the following are worth a mention: 12

fmt
----
Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format verbs are derived from C's but are simpler. Some verbs (%-sequences) that can be used:

io
---
This package provides basic interfaces to I/O primitives. Its primary job is to wrap existing implementations of such primitives, such as those in package os, into shared public interfaces that abstract the functionality, plus some other related primitives.

bufio
-----
This package implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

sort
----
The sort package provides primitives for sorting arrays and user-defined collections.

strconv
-------
The strconv package implements conversions to and from string representations of basic data types.

os
--
The os package provides a platform-independent interface to operating system functionality. The design is Unix-like.

sync
----
The package sync provides basic synchronization primitives such as mutual exclusion locks.

flag
-----
The flag package implements command-line flag parsing.

encoding/json
-------------
The encoding/json package implements encoding and decoding of JSON objects as defined in RFC 4627 [0].

html/template
-------------
Data-driven templates for generating textual output such as HTML.

Templates are executed by applying them to a data structure. Annotations in the template refer to elements of the data structure (typically a field of a struct or a key in a map) to control execution and derive values to be displayed. The template walks the structure as it executes and the "cursor" @ represents the value at the current location in the structure.

net/http
--------
The net/http package implements parsing of HTTP requests, replies, and URLs and provides an extensible HTTP server and a basic HTTP client.

unsafe
------
The unsafe package contains operations that step around the type safety of Go programs. Normally you don't need this package, but it is worth mentioning that unsafe Go programs are possible.

reflect
-------
The reflect package implements run-time reflection, allowing a program to manipulate objects with arbitrary types. The typical use is to take a value with static type interface{} and extract its dynamic type information by calling TypeOf, which returns an object with interface type Type. See Chapter 6. Interfaces, Section Introspection and reflection.

os/exec
-------
The os/exec package runs external commands.
*/
package main

import "even"
import "fmt"

func main() {

    if even.Even(2) == true {
        fmt.Println("2 is Even")
    }else{
        fmt.Println("2 is Odd")
    }

    if even.Odd(3) == true {
        fmt.Println("3 is Odd")
    }else{
        fmt.Println("3 is Even")
    }

}