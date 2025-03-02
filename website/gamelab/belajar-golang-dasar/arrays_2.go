package main

import "fmt"

func main() {
    var mobil = [3]string{"Toyota", "Honda", "Tesla"}
    
    for i := 0; i < len(mobil); i++ {
        fmt.Println(mobil[i])
    }
}

