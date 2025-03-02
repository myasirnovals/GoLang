package main

import "fmt"

func main() {
    var jam = 22
    
    if (jam < 10) {
        fmt.Println("Selamat pagi")
    } else if (jam < 14) {
        fmt.Println("Selamat siang")
    } else if (jam < 18) {
        fmt.Println("Selamat sore")
    } else {
        fmt.Println("Selamat malam")
    }
}
