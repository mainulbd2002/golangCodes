package main

import "fmt"

func main () {
    a := 100
    
    fmt.Println(a);
    fmt.Println(&a)
    
    var b *int = &a
    
    fmt.Println(*b)
}