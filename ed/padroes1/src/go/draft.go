package main
import (
    "fmt"
)

func padrao(n int) int{

    //caso base
    if(n == 1){
        return 20
    }

    //passo recursivo

    resultado := padrao(n-1) + 8

    //retorno
    return resultado


    // resultado
}

func main() {
    var n int
    fmt.Scan(&n)

    // resultado := (4+n) * (4+n) -n*n -4
    
    fmt.Println(padrao(n))
        

}
