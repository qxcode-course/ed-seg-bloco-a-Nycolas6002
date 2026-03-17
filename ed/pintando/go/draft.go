package main

import (
	"fmt"
    "math"
)
func main() {
    var a float64
    var b float64
    var c float64

    fmt.Scan(&a,&b,&c)

    perimetro := (a + b + c) / 2
    
    // fmt.Println(a,b,c)

    area := math.Sqrt(perimetro*(perimetro - a )*(perimetro - b )*(perimetro - c ))

    fmt.Printf("%.2f\n", area)
}
