package main

import "fmt"

func main() {
    var hari = 12
    
    switch hari {
        case 1:
            fmt.Println("Nilai hari = 1")
            fmt.Println("Maka, hari = Senin")
        case 2:
            fmt.Println("Nilai hari = 2")
            fmt.Println("Maka, hari = Selasa")
        case 3:
            fmt.Println("Nilai hari = 3")
            fmt.Println("Maka, hari = Rabu")
        case 4:
            fmt.Println("Nilai hari = 4")
            fmt.Println("Maka, hari = Kamis")
        case 5:
            fmt.Println("Nilai hari = 5")
            fmt.Println("Maka, hari = Jumat")
        case 6:
            fmt.Println("Nilai hari = 6")
            fmt.Println("Maka, hari = Sabtu")
        case 7:
            fmt.Println("Nilai hari = 7")
            fmt.Println("Maka, hari = Minggu")
        default:
            fmt.Println("Nilai hari = ", hari)
            fmt.Println("Maka, hari = tidak valid")
    }
}
