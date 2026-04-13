package main
import "fmt"

func imprimir(palavra string, tamanho int) {

    // caso base
    if(tamanho == 0){
        return
    }

    //passo recursivo

    imprimir(palavra[1:], tamanho-1)
    fmt.Println(palavra[0:])

}

func main() {
    // fmt.Println("Hello, World!")
    var palavra string
    fmt.Scan(&palavra)

    // fmt.Println(palavra[1:])
    imprimir(palavra, len(palavra))

}
