package main

import "fmt"

func main() {
    var mobil = [3]string{"Toyota", "Honda", "Tesla"}
    
    fmt.Println(mobil[1])
    
    mobil[1] = "Mazda"
    
    fmt.Println(mobil[1])
}
