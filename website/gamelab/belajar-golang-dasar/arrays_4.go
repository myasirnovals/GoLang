package main

import "fmt"

func main() {
    var mobil = [3]string{"Toyota", "Honda", "Tesla"}
    
    m := contains(mobil, "Tesla")
    
    fmt.Println(m)
}

func contains(array [3]string, str string) bool {
    for _, a := range array {
        if a == str {
            return true
        }
    }
    
    return false
}
