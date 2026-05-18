package main

import "fmt"

func main() {
    // Method 1: explicit type
    var name string = "Noman"
    var age  int    = 24
    var gpa  float64 = 3.60
    var isStudent bool = true

    // Method 2: short declaration (inside functions only)
    city :="Dhaka"   // Go infers the type automatically

    fmt.Println(name, age, gpa, isStudent, city)
}