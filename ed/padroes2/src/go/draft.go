package main
import "fmt"

func qtdPecas (x int) int{

    //caso base
    if(x == 1){
        return 3
    }

    //passo recursivo
    // resultado := qtdPecas(x-1) *  + (2 * x)
    resultado := qtdPecas(x-1) + 3 + ((x-1) *2)

    //retorno
    return resultado
}


func main() {
    var ordem int
    fmt.Scan(&ordem)
    // resultado := (ordem*ordem) + (ordem * 2)
    fmt.Println(qtdPecas(ordem))
    // fmt.Println(resultado)
    // ordem 3 
    // 15 peças

}
