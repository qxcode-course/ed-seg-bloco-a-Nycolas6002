package main
import "fmt"
func main() {
    var qtdElementos int
    fmt.Scan(&qtdElementos)

    animais := make(map[int]int)
    var pares int

    for i := 0; i < qtdElementos; i++ {
        var animal int
        fmt.Scan(&animal)

        if( animais[-animal] > 0){
            animais[-animal]--
            pares++
        }else{
            animais[animal]++
        }
    }

    fmt.Println(pares)
}
