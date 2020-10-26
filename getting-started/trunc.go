package main

import "fmt"

// Truncation assignment
func main() {
    var trunc int
    var floatnumber float32
    fmt.Printf("Enter a floating point number: ")
    num, err := fmt.Scan(&floatnumber)
    if err != nil || num != 1 {
     // handle invalid input
       fmt.Println(num, err)
       return
    }

    trunc = int(floatnumber)
    fmt.Println(trunc)
}

