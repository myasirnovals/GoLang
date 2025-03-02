package main

import "fmt"

func main() {
    var g = 0
    
    for g < 5 {
        fmt.Println(g)
        
        if g == 0 {
            break
        }
        
        g++
    }
}
