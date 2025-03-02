package main

import "fmt"

func main() {
    a := 7
    b := 17
    c := 27
    
    var hasil bool
    
    hasil = (a > b) && (a > c)
    
    fmt.Println(hasil)
}
