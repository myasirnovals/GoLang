package main

import "fmt"

func main() {
    jumlah(20, 88)
    jumlah(27, 46)
}

func jumlah(a int, b int) {
    var hasil = a + b
    
    fmt.Println(a, "+", b, "=", hasil)
}
